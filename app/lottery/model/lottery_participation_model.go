package model

import (
	"context"
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
		GetParticipationLotteryIdsByUserId(ctx context.Context, userId int64) ([]int64, error)
		CheckIsParticipatedByUserIdAndLotteryId(ctx context.Context, userId, lotteryId int64) (int64, error)
		GetLotteryWinListByLotteryId(ctx context.Context, lotteryId int64) ([]*LotteryParticipation, error)
		GetUserLotteryWinList(ctx context.Context, userId, lastId, limit int64) ([]*LotteryParticipation, error)
		GetLotteryParticipationListByLotteryId(ctx context.Context, lotteryId, page, limit int64) ([]*LotteryParticipation, error)
		GetLotteryParticipatorsByLotteryId(ctx context.Context, lotteryId int64) ([]int64, error)
		GetParticipatorsCountByLotteryId(ctx context.Context, lotteryId int64) (int64, error)
		GetParticipationCountByUserId(ctx context.Context, userId int64) (int64, error)
		GetWonCountByUserId(ctx context.Context, userId int64) (int64, error)
		GetAllLotteryListByUserId(ctx context.Context, userId, lastId, limit int64) ([]*LotteryParticipation, error)
		UpdateWinner(ctx context.Context, lotteryId, userId, prizeId int64) error
		GetLotteryParticipatedCount(ctx context.Context, userId int64) (int64, error)
	}

	customLotteryParticipationModel struct {
		*defaultLotteryParticipationModel
	}

	customLotteryParticipationLogicModel interface {
	}
)

func (c *customLotteryParticipationModel) GetLotteryParticipatedCount(ctx context.Context, userId int64) (int64, error) {
	var count int64
	err := c.QueryNoCacheCtx(ctx, func(db *gorm.DB) error {
		return db.Table(c.table).Where("user_id = ?", userId).Count(&count).Error
	})
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (c *customLotteryParticipationModel) UpdateWinner(ctx context.Context, lotteryId, userId, prizeId int64) error {
	data, err := c.FindOneByLotteryIdUserId(ctx, lotteryId, userId)
	if err != nil {
		return err
	}
	if data == nil {
		return gorm.ErrRecordNotFound
	}
	data.PrizeId = prizeId
	data.IsWon = 1
	return c.Update(ctx, nil, data)
}

func (c *customLotteryParticipationModel) GetLotteryParticipatorsByLotteryId(ctx context.Context, lotteryId int64) ([]int64, error) {
	userIds := make([]int64, 0)
	err := c.QueryNoCacheCtx(ctx, func(db *gorm.DB) error {
		return db.Table(c.table).Where("lottery_id = ?", lotteryId).Pluck("user_id", &userIds).Error
	})
	if err != nil {
		return nil, err
	}
	return userIds, nil
}

func (c *customLotteryParticipationModel) GetAllLotteryListByUserId(ctx context.Context, userId, lastId, limit int64) ([]*LotteryParticipation, error) {
	list := make([]*LotteryParticipation, 0)
	err := c.QueryNoCacheCtx(ctx, func(db *gorm.DB) error {
		if lastId > 0 {
			return db.Table(c.table).
				Where("user_id = ? AND id < ?", userId, lastId).Order("id desc").Limit(int(limit)).Find(&list).Error
		}
		return db.Table(c.table).Where("user_id = ?", userId).Order("id desc").Limit(int(limit)).Find(&list).Error
	})
	if err != nil {
		return []*LotteryParticipation{}, err
	}
	return list, nil
}

func (c *customLotteryParticipationModel) GetWonCountByUserId(ctx context.Context, userId int64) (int64, error) {
	var count int64
	err := c.QueryNoCacheCtx(ctx, func(db *gorm.DB) error {
		return db.Table(c.table).Where("user_id = ? AND is_won = 1", userId).Count(&count).Error
	})
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (c *customLotteryParticipationModel) GetParticipationCountByUserId(ctx context.Context, userId int64) (int64, error) {
	var count int64
	err := c.QueryNoCacheCtx(ctx, func(db *gorm.DB) error {
		return db.Table(c.table).Where("user_id = ?", userId).Count(&count).Error
	})
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (c *customLotteryParticipationModel) GetParticipatorsCountByLotteryId(ctx context.Context, lotteryId int64) (int64, error) {
	var count int64
	err := c.QueryNoCacheCtx(ctx, func(db *gorm.DB) error {
		return db.Table(c.table).Where("lottery_id = ?", lotteryId).Count(&count).Error
	})
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (c *customLotteryParticipationModel) GetLotteryParticipationListByLotteryId(ctx context.Context, lotteryId, page, limit int64) ([]*LotteryParticipation, error) {
	var list []*LotteryParticipation
	err := c.QueryNoCacheCtx(ctx, func(db *gorm.DB) error {
		return db.Table(c.table).Where("lottery_id = ?", lotteryId).
			Offset(int((page - 1) * limit)).Limit(int(limit)).Find(&list).Error
	})
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (c *customLotteryParticipationModel) GetUserLotteryWinList(ctx context.Context, userId, lastId, limit int64) ([]*LotteryParticipation, error) {
	var list []*LotteryParticipation
	err := c.QueryNoCacheCtx(ctx, func(db *gorm.DB) error {
		if lastId > 0 {
			return db.Table(c.table).
				Where("user_id = ? AND is_won = 1 AND id < ?", userId, lastId).
				Order("id DESC").Limit(int(limit)).Find(&list).Error
		}
		return db.Table(c.table).
			Where("user_id = ? AND is_won = 1", userId).
			Order("id DESC").Limit(int(limit)).Find(&list).Error
	})
	if err != nil {
		return []*LotteryParticipation{}, err
	}
	return list, nil
}

func (c *customLotteryParticipationModel) GetLotteryWinListByLotteryId(ctx context.Context, lotteryId int64) ([]*LotteryParticipation, error) {
	var list []*LotteryParticipation
	err := c.QueryNoCacheCtx(ctx, func(db *gorm.DB) error {
		return db.Table(c.table).
			Where("lottery_id = ? AND is_won = 1", lotteryId).
			Find(&list).Error
	})
	if err != nil {
		return []*LotteryParticipation{}, err
	}
	return list, nil
}

func (c *customLotteryParticipationModel) CheckIsParticipatedByUserIdAndLotteryId(ctx context.Context, lotteryId, userId int64) (int64, error) {
	var count int64
	err := c.QueryNoCacheCtx(ctx, func(db *gorm.DB) error {
		return db.Table(c.table).
			Where("lottery_id = ? AND user_id = ?", lotteryId, userId).Count(&count).Error
	})
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (c *customLotteryParticipationModel) GetParticipationLotteryIdsByUserId(ctx context.Context, userId int64) ([]int64, error) {
	var lotteryIds []int64
	err := c.QueryNoCacheCtx(ctx, func(db *gorm.DB) error {
		return db.Table(c.table).Where("user_id = ?", userId).Pluck("lottery_id", &lotteryIds).Error
	})
	if err != nil {
		return []int64{}, err
	}

	return lotteryIds, nil
}

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
