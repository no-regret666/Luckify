package logic

import (
	"Luckify/common/xerr"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"

	"Luckify/app/lottery/cmd/rpc/internal/svc"
	"Luckify/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnnounceLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnnounceLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnnounceLotteryLogic {
	return &AnnounceLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type Winner struct {
	LotteryId int64
	UserId    int64
	PrizeId   int64
}

// 策略模式
type LotteryStrategy interface {
	Run() error
}

type PeopleLotteryStrategy struct {
	l           *AnnounceLotteryLogic
	CurrentTime time.Time
}

type TimeLotteryStrategy struct {
	l           *AnnounceLotteryLogic
	CurrentTime time.Time
}

func (p *TimeLotteryStrategy) Run() error {
	l := p.l
	pendingList, err := l.svcCtx.LotteryModel.GetPendingLotteryListWhichTypeIsTimeStrategyAndAnnouncedTimeBeforeNow(l.ctx, time.Now())
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_PENDING_LOTTERY_LIST_FAIL), "GetPendingLotteryListWhichTypeIsTimeStrategyAndAnnouncedTimeBeforeNow err:%v", err)
	}
	for _, lottery := range pendingList {
		// 开奖
		err = l.svcCtx.TransactCtx(l.ctx, func(db *gorm.DB) error {
			// 获取所有的奖品
			prizes, err := l.svcCtx.PrizeModel.FindByLotteryId(l.ctx, lottery.Id)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_PRIZE_ERROR), "FindByLotteryId err:%v", err)
			}

			// 获取所有参与者

		})
	}
}

func (l *AnnounceLotteryLogic) AnnounceLottery(in *pb.AnnounceLotteryReq) (*pb.AnnounceLotteryResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AnnounceLotteryResp{}, nil
}
