package logic

import (
	"Luckify/app/comment/model"
	"Luckify/common/xerr"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"Luckify/app/comment/cmd/rpc/internal/svc"
	"Luckify/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PraiseCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPraiseCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PraiseCommentLogic {
	return &PraiseCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PraiseCommentLogic) PraiseComment(in *pb.PraiseCommentReq) (*pb.PraiseCommentResp, error) {
	dbPraise, err := l.svcCtx.PraiseModel.IsPraise(l.ctx, in.CommentId, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_PRAISE_COMMENT_ERROR), "PraiseComment err: %v", err)
	}
	//若点赞，取消点赞
	if dbPraise.Id != 0 {
		err := l.svcCtx.TransactCtx(l.ctx, func(db *gorm.DB) error {
			err := l.svcCtx.PraiseModel.Delete(l.ctx, db, dbPraise.Id)
			if err != nil {
				return err
			}
			_, err = l.svcCtx.CommentModel.AddPraiseCount(l.ctx, in.CommentId, -1)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_PRAISE_COMMENT_ERROR), "PraiseComment err: %v", err)
		}
	} else {
		err := l.svcCtx.TransactCtx(l.ctx, func(db *gorm.DB) error {
			err := l.svcCtx.PraiseModel.Insert(l.ctx, db, &model.Praise{
				CommentId: in.CommentId,
				UserId:    in.UserId,
			})
			if err != nil {
				return err
			}
			_, err = l.svcCtx.CommentModel.AddPraiseCount(l.ctx, in.CommentId, 1)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_PRAISE_COMMENT_ERROR), "PraiseComment err: %v", err)
		}
	}

	return &pb.PraiseCommentResp{}, nil
}
