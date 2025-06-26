package lottery

import (
	"Luckify/app/lottery/cmd/rpc/pb"
	"Luckify/app/lottery/model"
	"Luckify/common/utility"
	"context"
	"github.com/pkg/errors"

	"Luckify/app/lottery/cmd/api/internal/svc"
	"Luckify/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddInstantLotteryParticipationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 参与即抽即中抽奖
func NewAddInstantLotteryParticipationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddInstantLotteryParticipationLogic {
	return &AddInstantLotteryParticipationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddInstantLotteryParticipationLogic) AddInstantLotteryParticipation(req *types.AddInstantLotteryParticipationReq) (resp *types.AddInstantLotteryParticipationResp, err error) {
	pbResp, err := l.svcCtx.LotteryRpc.AddInstantLotteryParticipation(l.ctx, &pb.AddInstantLotteryParticipationReq{
		LotteryId: req.LotteryId,
		UserId:    utility.GetUserIdFromCtx(l.ctx),
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrParticipateLottery, "AddInstantLotteryParticipation rpc err: %v", err)
	}

	return &types.AddInstantLotteryParticipationResp{
		PrizeId: pbResp.Id,
	}, nil
}
