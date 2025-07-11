package checkin

import (
	"Luckify/app/checkin/cmd/rpc/checkin"
	"Luckify/common/utility"
	"context"

	"Luckify/app/checkin/cmd/api/internal/svc"
	"Luckify/app/checkin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCheckinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取签到状态以及积分
func NewGetCheckinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCheckinLogic {
	return &GetCheckinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCheckinLogic) GetCheckin(req *types.GetCheckinReq) (resp *types.GetCheckinResp, err error) {
	userId := utility.GetUserIdFromCtx(l.ctx)
	pbResp, err := l.svcCtx.CheckinRpc.GetCheckinRecordByUserId(l.ctx, &checkin.GetCheckinRecordByUserIdReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetCheckinResp{
		ContinuousCheckinDays: pbResp.ContinuousCheckinDays,
		State:                 pbResp.State,
		Integral:              pbResp.Integral,
		SubStatus:             pbResp.SubStatus,
	}, nil
}
