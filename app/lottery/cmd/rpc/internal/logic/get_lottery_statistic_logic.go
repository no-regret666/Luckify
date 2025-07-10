package logic

import (
	"Luckify/common/xerr"
	"context"
	"github.com/pkg/errors"

	"Luckify/app/lottery/cmd/rpc/internal/svc"
	"Luckify/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLotteryStatisticLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLotteryStatisticLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryStatisticLogic {
	return &GetLotteryStatisticLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLotteryStatisticLogic) GetLotteryStatistic(in *pb.GetLotteryStatisticReq) (*pb.GetLotteryStatisticResp, error) {
	createdCount, err := l.svcCtx.LotteryModel.GetCreatedCountByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_CREATED_COUNT_BY_USER_ID), "GetCreatedCountByUserId error: %v", err)
	}
	wonCount, err := l.svcCtx.LotteryParticipationModel.GetWonCountByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_WON_COUNT_BY_USER_ID), "GetWonCountByUserId error: %v", err)
	}
	participationCount, err := l.svcCtx.LotteryParticipationModel.GetParticipationCountByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_PARTICIPATION_COUNT_BY_USER_ID), "GetParticipationCountByUserId error: %v", err)
	}

	return &pb.GetLotteryStatisticResp{
		CreatedCount:       int64(createdCount),
		WonCount:           int64(wonCount),
		ParticipationCount: int64(participationCount),
	}, nil
}
