package svc

import (
	"Luckify/app/lottery/cmd/rpc/internal/config"
	"Luckify/app/lottery/model"
	"Luckify/app/usercenter/cmd/rpc/usercenter"
	"context"
	"database/sql"
	"github.com/SpectatorNan/gorm-zero/gormc/config/mysql"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	goredislib "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config                    config.Config
	LotteryModel              model.LotteryModel
	LotteryParticipationModel model.LotteryParticipationModel
	PrizeModel                model.PrizeModel
	UsercenterRpc             usercenter.Usercenter
	RedsyncClient             *redsync.Redsync
	TransactCtx               func(context.Context, func(db *gorm.DB) error, ...*sql.TxOptions) error
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := mysql.Connect(c.MySQL)
	if err != nil {
		panic(err)
	}
	client := goredislib.NewClient(&goredislib.Options{
		Addr:     c.Cache[0].Host,
		Password: c.Cache[0].Pass,
	})
	pool := goredis.NewPool(client)

	return &ServiceContext{
		Config:                    c,
		LotteryModel:              model.NewLotteryModel(db, c.Cache),
		LotteryParticipationModel: model.NewLotteryParticipationModel(db, c.Cache),
		PrizeModel:                model.NewPrizeModel(db, c.Cache),
		UsercenterRpc:             usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterConf)),
		RedsyncClient:             redsync.New(pool),
		TransactCtx: func(ctx context.Context, fn func(*gorm.DB) error, opts ...*sql.TxOptions) error {
			return db.WithContext(ctx).Transaction(fn, opts...)
		},
	}
}
