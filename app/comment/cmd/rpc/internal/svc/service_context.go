package svc

import (
	"Luckify/app/comment/cmd/rpc/internal/config"
	"Luckify/app/comment/model"
	"context"
	"database/sql"
	"github.com/SpectatorNan/gorm-zero/gormc/config/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config       config.Config
	CommentModel model.CommentModel
	PraiseModel  model.PraiseModel
	TransactCtx  func(context.Context, func(db *gorm.DB) error, ...*sql.TxOptions) error
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := mysql.Connect(c.MySQL)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:       c,
		CommentModel: model.NewCommentModel(db, c.Cache),
		PraiseModel:  model.NewPraiseModel(db, c.Cache),
		TransactCtx: func(ctx context.Context, fn func(db *gorm.DB) error, opts ...*sql.TxOptions) error {
			return db.WithContext(ctx).Transaction(fn, opts...)
		},
	}
}
