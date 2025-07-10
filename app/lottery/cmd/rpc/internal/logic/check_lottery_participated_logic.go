package logic

import (
	"Luckify/common/xerr"
	"context"
	"github.com/pkg/errors"

	"Luckify/app/lottery/cmd/rpc/internal/svc"
	"Luckify/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLotteryParticipatedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLotteryParticipatedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLotteryParticipatedLogic {
	return &CheckLotteryParticipatedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLotteryParticipatedLogic) CheckLotteryParticipated(in *pb.CheckLotteryParticipatedReq) (*pb.CheckLotteryParticipatedResp, error) {
	count, err := l.svcCtx.LotteryParticipationModel.GetLotteryParticipatedCount(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_LOTTERY_PARTICIPATED_COUNT), "Failed to get lottery participated count: %v", err)
	}

	count = min(count, 1)

	return &pb.CheckLotteryParticipatedResp{
		Participated: count,
	}, nil
}
