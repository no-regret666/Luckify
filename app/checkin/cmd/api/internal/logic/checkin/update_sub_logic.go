package checkin

import (
	"Luckify/app/checkin/cmd/rpc/checkin"
	"Luckify/app/checkin/model"
	"Luckify/common/utility"
	"context"
	"github.com/pkg/errors"

	"Luckify/app/checkin/cmd/api/internal/svc"
	"Luckify/app/checkin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSubLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 改变任务签到状态
func NewUpdateSubLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubLogic {
	return &UpdateSubLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSubLogic) UpdateSub(req *types.UpdateSubReq) (resp *types.UpdateSubResp, err error) {
	userId := utility.GetUserIdFromCtx(l.ctx)
	_, err = l.svcCtx.CheckinRpc.UpdateSub(l.ctx, &checkin.UpdateSubReq{
		UserId: userId,
		State:  req.State,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrUpdateSub, "UpdateSub err: %s", err.Error())
	}
	return
}
