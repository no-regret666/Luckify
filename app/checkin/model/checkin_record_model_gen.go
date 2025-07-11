// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/SpectatorNan/gorm-zero/gormc/batchx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var (
	cacheGoLotteryCheckinCheckinRecordIdPrefix = "cache:goLotteryCheckin:checkinRecord:id:"
)

type (
	checkinRecordModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *CheckinRecord) error
		BatchInsert(ctx context.Context, tx *gorm.DB, news []CheckinRecord) error
		FindOne(ctx context.Context, id int64) (*CheckinRecord, error)
		Update(ctx context.Context, tx *gorm.DB, data *CheckinRecord) error
		BatchUpdate(ctx context.Context, tx *gorm.DB, olds, news []CheckinRecord) error
		BatchDelete(ctx context.Context, tx *gorm.DB, datas []CheckinRecord) error

		Delete(ctx context.Context, tx *gorm.DB, id int64) error
		// deprecated. recommend add a transaction in service context instead of using this
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultCheckinRecordModel struct {
		gormc.CachedConn
		table string
	}

	CheckinRecord struct {
		Id                    int64        `gorm:"column:id;primary_key"`
		UserId                int64        `gorm:"column:user_id"`
		ContinuousCheckinDays int64        `gorm:"column:continuous_checkin_days"` // Number of consecutive check-in days
		State                 int64        `gorm:"column:state"`                   // Whether to sign in on the day, 1 means signed, 0 means not signed.
		LastCheckinDate       sql.NullTime `gorm:"column:last_checkin_date"`       // Date of last check-in
		CreateTime            time.Time    `gorm:"column:create_time"`             // 创建时间
		UpdateTime            time.Time    `gorm:"column:update_time"`             // 更新时间
	}
)

func (CheckinRecord) TableName() string {
	return "`checkin_record`"
}

func newCheckinRecordModel(conn *gorm.DB, c cache.CacheConf) *defaultCheckinRecordModel {
	return &defaultCheckinRecordModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`checkin_record`",
	}
}

func (m *defaultCheckinRecordModel) GetCacheKeys(data *CheckinRecord) []string {
	if data == nil {
		return []string{}
	}
	goLotteryCheckinCheckinRecordIdKey := fmt.Sprintf("%s%v", cacheGoLotteryCheckinCheckinRecordIdPrefix, data.Id)
	cacheKeys := []string{
		goLotteryCheckinCheckinRecordIdKey,
	}
	cacheKeys = append(cacheKeys, m.customCacheKeys(data)...)
	return cacheKeys
}

func (m *defaultCheckinRecordModel) Insert(ctx context.Context, tx *gorm.DB, data *CheckinRecord) error {

	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Omit(checkinRecordOmitColumns...).Save(&data).Error
	}, m.GetCacheKeys(data)...)
	return err
}
func (m *defaultCheckinRecordModel) BatchInsert(ctx context.Context, tx *gorm.DB, news []CheckinRecord) error {

	err := batchx.BatchExecCtx(ctx, m, news, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Create(&news).Error
	})

	return err
}

func (m *defaultCheckinRecordModel) FindOne(ctx context.Context, id int64) (*CheckinRecord, error) {
	goLotteryCheckinCheckinRecordIdKey := fmt.Sprintf("%s%v", cacheGoLotteryCheckinCheckinRecordIdPrefix, id)
	var resp CheckinRecord
	err := m.QueryCtx(ctx, &resp, goLotteryCheckinCheckinRecordIdKey, func(conn *gorm.DB) error {
		return conn.Model(&CheckinRecord{}).Where("`id` = ?", id).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCheckinRecordModel) Update(ctx context.Context, tx *gorm.DB, data *CheckinRecord) error {
	old, err := m.FindOne(ctx, data.Id)
	if err != nil && errors.Is(err, ErrNotFound) {
		return err
	}
	clearKeys := append(m.GetCacheKeys(old), m.GetCacheKeys(data)...)
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Omit(checkinRecordOmitColumns...).Save(data).Error
	}, clearKeys...)
	return err
}
func (m *defaultCheckinRecordModel) BatchUpdate(ctx context.Context, tx *gorm.DB, olds, news []CheckinRecord) error {
	clearData := make([]CheckinRecord, 0, len(olds)+len(news))
	clearData = append(clearData, olds...)
	clearData = append(clearData, news...)
	err := batchx.BatchExecCtx(ctx, m, clearData, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(&news).Error
	})

	return err
}

func (m *defaultCheckinRecordModel) Delete(ctx context.Context, tx *gorm.DB, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil
		}
		return err
	}
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Delete(&CheckinRecord{}, id).Error
	}, m.GetCacheKeys(data)...)
	return err
}

func (m *defaultCheckinRecordModel) BatchDelete(ctx context.Context, tx *gorm.DB, datas []CheckinRecord) error {
	err := batchx.BatchExecCtx(ctx, m, datas, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Delete(&datas).Error
	})

	return err
}

// deprecated. recommend add a transaction in service context instead of using this
func (m *defaultCheckinRecordModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.TransactCtx(ctx, fn)
}

func (m *defaultCheckinRecordModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheGoLotteryCheckinCheckinRecordIdPrefix, primary)
}

func (m *defaultCheckinRecordModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&CheckinRecord{}).Where("`id` = ?", primary).Take(v).Error
}
