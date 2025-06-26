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

type AddLotteryParticipationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 参与抽奖
func NewAddLotteryParticipationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLotteryParticipationLogic {
	return &AddLotteryParticipationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLotteryParticipationLogic) AddLotteryParticipation(req *types.AddLotteryParticipationReq) (resp *types.AddLotteryParticipationResp, err error) {
	_, err = l.svcCtx.LotteryRpc.AddLotteryParticipation(l.ctx, &pb.AddLotteryParticipationReq{
		LotteryId: req.LotteryId,
		UserId:    utility.GetUserIdFromCtx(l.ctx),
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrParticipateLottery, "AddLotteryParticipation rpc error: %v", err)
	}

	return &types.AddLotteryParticipationResp{}, nil
}
