package logic

import (
	"Luckify/app/usercenter/model"
	"Luckify/common/xerr"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"Luckify/app/usercenter/cmd/rpc/internal/svc"
	"Luckify/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserSponsorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserSponsorLogic {
	return &AddUserSponsorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserSponsorLogic) AddUserSponsor(in *pb.AddUserSponsorReq) (*pb.AddUserSponsorResp, error) {
	userSponsor := new(model.UserSponsor)
	_ = copier.Copy(userSponsor, in)
	err := l.svcCtx.UserSponsorModel.Insert(l.ctx, nil, userSponsor)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "AddUserSponsor err: %v", err)
	}

	return &pb.AddUserSponsorResp{
		Id: userSponsor.Id,
	}, nil
}
