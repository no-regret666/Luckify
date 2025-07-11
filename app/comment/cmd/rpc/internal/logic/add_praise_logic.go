package logic

import (
	"Luckify/app/comment/model"
	"Luckify/common/xerr"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"Luckify/app/comment/cmd/rpc/internal/svc"
	"Luckify/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPraiseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPraiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPraiseLogic {
	return &AddPraiseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddPraiseLogic) AddPraise(in *pb.AddPraiseReq) (*pb.AddPraiseResp, error) {
	dbPraise := &model.Praise{}
	_ = copier.Copy(dbPraise, in)
	err := l.svcCtx.PraiseModel.Insert(l.ctx, nil, dbPraise)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_INSERT_PRAISE_ERROR), "AddPraise err: %v", err)
	}

	return &pb.AddPraiseResp{}, nil
}
