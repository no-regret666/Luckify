package checkin

import (
	"Luckify/app/checkin/cmd/rpc/pb"
	"Luckify/app/checkin/model"
	"Luckify/common/utility"
	"context"
	"github.com/pkg/errors"

	"Luckify/app/checkin/cmd/api/internal/svc"
	"Luckify/app/checkin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 签到操作
func NewCheckinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckinLogic {
	return &CheckinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckinLogic) Checkin(req *types.CheckinReq) (resp *types.CheckinResp, err error) {
	userId := utility.GetUserIdFromCtx(l.ctx)
	pbRecord, err := l.svcCtx.CheckinRpc.UpdateCheckinRecord(l.ctx, &pb.UpdateCheckinRecordReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrCheckin, "Checkin failed,req: %+v", req)
	}

	return &types.CheckinResp{
		ContinuousCheckinDays: pbRecord.ContinuousCheckinDays,
		State:                 pbRecord.State,
		Integral:              pbRecord.Integral,
	}, nil
}
