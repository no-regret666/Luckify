package logic

import (
	"Luckify/common/xerr"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"Luckify/app/usercenter/cmd/rpc/internal/svc"
	"Luckify/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SponsorDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSponsorDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SponsorDetailLogic {
	return &SponsorDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SponsorDetailLogic) SponsorDetail(in *pb.SponsorDetailReq) (*pb.SponsorDetailResp, error) {
	userSponsor, err := l.svcCtx.UserSponsorModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find user sponsor failed: %v", err)
	}

	resp := new(pb.SponsorDetailResp)
	_ = copier.Copy(resp, userSponsor)

	return &pb.SponsorDetailResp{}, nil
}
