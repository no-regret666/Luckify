package model

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ PrizeModel = (*customPrizeModel)(nil)
var prizeOmitColumns = []string{"create_time", "update_time"}

type (
	// PrizeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPrizeModel.
	PrizeModel interface {
		prizeModel
		customPrizeLogicModel
		FindByLotteryId(ctx context.Context, lotteryId int64) ([]*Prize, error)
		FindFirstLevelByLotteryId(ctx context.Context, lotteryId int64) (*Prize, error)
		DecrStock(ctx context.Context, prizeId int64) error
	}

	customPrizeModel struct {
		*defaultPrizeModel
	}

	customPrizeLogicModel interface {
	}
)

func (c *customPrizeModel) DecrStock(ctx context.Context, prizeId int64) error {
	err := c.ExecCtx(ctx, func(db *gorm.DB) error {
		res := db.Model(&Prize{}).Where("id = ? and stock > 0", prizeId).Update("stock", gorm.Expr("stock - 1"))
		if res.RowsAffected == 0 {
			return errors.New("库存不足或奖品不存在")
		}
		return res.Error
	})
	return err
}

func (c *customPrizeModel) FindFirstLevelByLotteryId(ctx context.Context, lotteryId int64) (*Prize, error) {
	var prize Prize
	err := c.QueryNoCacheCtx(ctx, func(db *gorm.DB) error {
		return db.Where("lottery_id = ? and level = 1", lotteryId).Find(&prize).Error
	})
	if err != nil {
		return nil, err
	}
	return &prize, nil
}

func (c *customPrizeModel) FindByLotteryId(ctx context.Context, lotteryId int64) ([]*Prize, error) {
	var prizes []*Prize
	err := c.QueryNoCacheCtx(ctx, func(db *gorm.DB) error {
		return db.Where("lottery_id = ?", lotteryId).Find(&prizes).Error
	})
	if err != nil {
		return nil, err
	}
	return prizes, nil
}

// NewPrizeModel returns a model for the database table.
func NewPrizeModel(conn *gorm.DB, c cache.CacheConf) PrizeModel {
	return &customPrizeModel{
		defaultPrizeModel: newPrizeModel(conn, c),
	}
}

func (m *defaultPrizeModel) customCacheKeys(data *Prize) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
