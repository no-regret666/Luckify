package event

import (
	"context"

	"Luckify/app/notice/cmd/api/internal/svc"
	"Luckify/app/notice/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReceiveEventLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 接受小程序回调消息
func NewReceiveEventLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReceiveEventLogic {
	return &ReceiveEventLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReceiveEventLogic) ReceiveEvent(_ *types.ReceiveEventReq, r *types.ReceiveEventReq) (resp *types.ReceiveEventResp, err error) {
	logx.Info("ReceiveEventLogic received an event", r)

	return
}
