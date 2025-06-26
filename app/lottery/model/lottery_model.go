package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ LotteryModel = (*customLotteryModel)(nil)
var lotteryOmitColumns = []string{"create_time", "update_time"}

type (
	// LotteryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLotteryModel.
	LotteryModel interface {
		lotteryModel
		customLotteryLogicModel
	}

	customLotteryModel struct {
		*defaultLotteryModel
	}

	customLotteryLogicModel interface {
	}
)

// NewLotteryModel returns a model for the database table.
func NewLotteryModel(conn *gorm.DB, c cache.CacheConf) LotteryModel {
	return &customLotteryModel{
		defaultLotteryModel: newLotteryModel(conn, c),
	}
}

func (m *defaultLotteryModel) customCacheKeys(data *Lottery) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
