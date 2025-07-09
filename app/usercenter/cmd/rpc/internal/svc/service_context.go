package svc

import (
	"Luckify/app/usercenter/cmd/rpc/internal/config"
	"Luckify/app/usercenter/model"
	"context"
	"database/sql"
	"github.com/SpectatorNan/gorm-zero/gormc/config/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config           config.Config
	UserModel        model.UserModel
	UserAuthModel    model.UserAuthModel
	UserSponsorModel model.UserSponsorModel

	//事务包装器函数，当需要多个数据库写操作必须同时成功或失败时，就应该使用
	TransacCtx func(context.Context, func(db *gorm.DB) error, ...*sql.TxOptions) error
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := mysql.Connect(c.MySQL)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:           c,
		UserModel:        model.NewUserModel(db, c.Cache),
		UserAuthModel:    model.NewUserAuthModel(db, c.Cache),
		UserSponsorModel: model.NewUserSponsorModel(db, c.Cache),
		TransacCtx: func(ctx context.Context, f func(db *gorm.DB) error, options ...*sql.TxOptions) error {
			return db.WithContext(ctx).Transaction(f, options...)
		},
	}
}
