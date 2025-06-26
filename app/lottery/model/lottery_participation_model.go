package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ LotteryParticipationModel = (*customLotteryParticipationModel)(nil)
var lotteryParticipationOmitColumns = []string{"create_time", "update_time"}

type (
	// LotteryParticipationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLotteryParticipationModel.
	LotteryParticipationModel interface {
		lotteryParticipationModel
		customLotteryParticipationLogicModel
	}

	customLotteryParticipationModel struct {
		*defaultLotteryParticipationModel
	}

	customLotteryParticipationLogicModel interface {
	}
)

// NewLotteryParticipationModel returns a model for the database table.
func NewLotteryParticipationModel(conn *gorm.DB, c cache.CacheConf) LotteryParticipationModel {
	return &customLotteryParticipationModel{
		defaultLotteryParticipationModel: newLotteryParticipationModel(conn, c),
	}
}

func (m *defaultLotteryParticipationModel) customCacheKeys(data *LotteryParticipation) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
