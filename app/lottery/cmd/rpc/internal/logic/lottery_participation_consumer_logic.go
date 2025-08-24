package logic

import (
	"Luckify/app/lottery/cmd/rpc/internal/svc"
	"Luckify/app/lottery/cmd/rpc/pb"
	"Luckify/app/lottery/model"
	"Luckify/common/constants"
	"Luckify/common/utility"
	"Luckify/common/xerr"
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type LotteryConsumerGroupHandler struct {
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	batchSize  int
	buffer     []*sarama.ConsumerMessage
	lastCommit time.Time
}

func NewLotteryConsumerHandler(ctx context.Context, svcCtx *svc.ServiceContext) *LotteryConsumerGroupHandler {
	return &LotteryConsumerGroupHandler{
		ctx:        ctx,
		svcCtx:     svcCtx,
		batchSize:  5,
		buffer:     make([]*sarama.ConsumerMessage, 0, 5),
		lastCommit: time.Now(),
	}
}

func (l *LotteryConsumerGroupHandler) Setup(sess sarama.ConsumerGroupSession) error {
	return nil
}

func (l *LotteryConsumerGroupHandler) Cleanup(sess sarama.ConsumerGroupSession) error {
	return nil
}

func (l *LotteryConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var req pb.MessageValue
		if err := proto.Unmarshal(msg.Value, &req); err != nil {
			logx.Errorf("X proto 解析失败: %v,message: %s", err, string(msg.Value))
			continue
		}

		logx.Infof("消费消息: ", req)

		requestIdKey := fmt.Sprintf("lottery:processed_request_id:%d", req.LotteryId)
		isMember, err := l.svcCtx.RedisClient.SIsMember(l.ctx, requestIdKey, req.RequestId).Result()
		if err != nil {
			logx.Errorf("❌ Redis 查询失败: %v,requestId = %s", err, req.RequestId)
			continue // 避免整个consumer挂掉
		}

		if isMember {
			logx.Infof("⚠️ 重复 requestId: %s，跳过", req.RequestId)
			sess.MarkMessage(msg, "")
			continue
		}

		randomCode := utility.Random(constants.ProbabilityMax)
		prizePoolKey := constants.InstantLotteryPrizePoolRedisKey + strconv.Itoa(int(req.LotteryId))
		result, err := l.svcCtx.RedisClient.Eval(l.ctx, constants.LotteryLuaScript, []string{prizePoolKey}, randomCode).Result()

		if err != nil {
			logx.Errorf(" ❌ Lua 脚本执行失败: %v,requestId = %s", err, req.RequestId)
			continue // 可考虑重试
		}

		logx.Infof(" Lua 脚本执行结果: %v", result)

		var prizeId int64
		switch v := result.(type) {
		case string:
			prizeId, _ = strconv.ParseInt(v, 10, 64)
		case float64:
			prizeId = int64(v)
		case int64:
			prizeId = v
		default:
			logx.Errorf("⚠️ prizeId 类型异常: %T", result)
		}

		logx.Infof("抽奖结果: %d", prizeId)

		isWon := constants.PrizeNotWon
		if prizeId != 0 {
			isWon = constants.PrizeHasWon
		}

		// 中奖才需要更新数据库
		err = l.svcCtx.TransactCtx(l.ctx, func(db *gorm.DB) error {
			if isWon == constants.PrizeHasWon {
				if err := l.svcCtx.PrizeModel.DecrStock(l.ctx, prizeId); err != nil && !errors.Is(err, model.ErrRowsAffectedZero) {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_DECR_PRIZE_STOCK_ERR), "库存扣减失败,prize=%d", prizeId)
				}

				// 扣减库存失败也算未中奖
				if errors.Is(err, model.ErrRowsAffectedZero) {
					isWon = constants.PrizeNotWon
					prizeId = 0
				}
			}
			return l.svcCtx.LotteryParticipationModel.Insert(l.ctx, db, &model.LotteryParticipation{
				LotteryId: req.LotteryId,
				UserId:    req.UserId,
				IsWon:     int64(isWon),
				PrizeId:   prizeId,
			})
		})
		if err != nil {
			logx.Errorf("❌ 数据库操作失败: %v, requestId=%s", err, req.RequestId)
			continue
		}

		// 把 requestId 加入 Redis 的去重集合（异步/持久化都可以考虑）
		if _, err := l.svcCtx.RedisClient.SAdd(l.ctx, requestIdKey, req.RequestId).Result(); err != nil {
			logx.Errorf(" ⚠️ 添加去重失败，但不影响主流程: %v", err)
		}

		// 正常消费成功
		sess.MarkMessage(msg, "")
		l.buffer = append(l.buffer, msg)

		if len(l.buffer) >= l.batchSize || time.Since(l.lastCommit) > 3*time.Second {
			logx.Infof("批量提交 offset (%d 条)", len(l.buffer))
			sess.Commit()
			l.buffer = []*sarama.ConsumerMessage{}
			l.lastCommit = time.Now()
		}
	}
	return nil
}

func HandleInstantLotteryParticipationMessage(svcCtx *svc.ServiceContext) {
	conf := sarama.NewConfig()
	conf.Consumer.Return.Errors = true
	conf.Consumer.Offsets.AutoCommit.Enable = true

	groupId := svcCtx.Config.Kafka.Group
	topic := svcCtx.Config.Kafka.Topic

	consumerGroup, err := sarama.NewConsumerGroup(svcCtx.Config.Kafka.Host, groupId, conf)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	handler := NewLotteryConsumerHandler(ctx, svcCtx)

	go func() {
		for {
			if err := consumerGroup.Consume(ctx, []string{topic}, handler); err != nil {
				logx.Infof(" ❌ 消费过程中发生错误: %v", err)
			}
		}
	}()

	logx.Info(" 抽奖消费者组已启动，等待消息...")

	//阻塞主线程
	select {}
}
