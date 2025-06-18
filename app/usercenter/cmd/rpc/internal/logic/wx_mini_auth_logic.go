package logic

import (
	"context"

	"Luckify/app/usercenter/cmd/rpc/internal/svc"
	"Luckify/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type WxMiniAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWxMiniAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxMiniAuthLogic {
	return &WxMiniAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WxMiniAuthLogic) WxMiniAuth(in *pb.WxMiniAuthReq) (*pb.WxMiniAuthResp, error) {
	// todo: add your logic here and delete this line

	return &pb.WxMiniAuthResp{}, nil
}
