package model

import (
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
	}

	customPrizeModel struct {
		*defaultPrizeModel
	}

	customPrizeLogicModel interface {
	}
)

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
