package user_sponsor

import (
	"Luckify/app/usercenter/cmd/rpc/pb"
	"Luckify/app/usercenter/model"
	"context"
	"github.com/pkg/errors"

	"Luckify/app/usercenter/cmd/api/internal/svc"
	"Luckify/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SponsorDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除赞助商
func NewSponsorDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SponsorDelLogic {
	return &SponsorDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SponsorDelLogic) SponsorDel(req *types.SponsorDelReq) (resp *types.SponsorDelResp, err error) {
	_, err = l.svcCtx.UsercenterRpc.DelUserSponsor(l.ctx, &pb.DelUserSponsorReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrDeleteUserSponsor, "rpc error: %v", err)
	}

	return &types.SponsorDelResp{}, nil
}
