package user_sponsor

import (
	"Luckify/app/usercenter/cmd/rpc/usercenter"
	"Luckify/app/usercenter/model"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"Luckify/app/usercenter/cmd/api/internal/svc"
	"Luckify/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SponsorDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 赞助商详情
func NewSponsorDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SponsorDetailLogic {
	return &SponsorDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SponsorDetailLogic) SponsorDetail(req *types.SponsorDetailReq) (resp *types.SponsorDetailResp, err error) {
	pbResp, err := l.svcCtx.UsercenterRpc.SponsorDetail(l.ctx, &usercenter.SponsorDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrGetUserSponsorDetail, "rpc error %v", err)
	}

	resp = new(types.SponsorDetailResp)
	_ = copier.Copy(resp, pbResp)

	return resp, nil
}
