package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ TaskRecordModel = (*customTaskRecordModel)(nil)
var taskRecordOmitColumns = []string{"create_time", "update_time"}

type (
	// TaskRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTaskRecordModel.
	TaskRecordModel interface {
		taskRecordModel
		customTaskRecordLogicModel

		FindByUserIdAndTaskId(ctx context.Context, userId, taskId int64) (*TaskRecord, error)
		FindByUserId(ctx context.Context, userId int64) ([]*TaskRecord, error)
	}

	customTaskRecordModel struct {
		*defaultTaskRecordModel
	}

	customTaskRecordLogicModel interface {
	}
)

func (c *customTaskRecordModel) FindByUserIdAndTaskId(ctx context.Context, userId, taskId int64) (*TaskRecord, error) {
	taskRecord := &TaskRecord{}
	err := c.QueryNoCacheCtx(ctx, func(db *gorm.DB) error {
		return db.Table(c.table).Where("user_id = ? AND task_id = ?", userId, taskId).First(&taskRecord).Error
	})
	if err != nil {
		return nil, err
	}
	return taskRecord, nil
}

func (c *customTaskRecordModel) FindByUserId(ctx context.Context, userId int64) ([]*TaskRecord, error) {
	taskRecords := make([]*TaskRecord, 0)
	err := c.QueryNoCacheCtx(ctx, func(db *gorm.DB) error {
		return db.Table(c.table).Where("user_id = ?", userId).Find(&taskRecords).Error
	})
	if err != nil {
		return nil, err
	}
	return taskRecords, nil
}

// NewTaskRecordModel returns a model for the database table.
func NewTaskRecordModel(conn *gorm.DB, c cache.CacheConf) TaskRecordModel {
	return &customTaskRecordModel{
		defaultTaskRecordModel: newTaskRecordModel(conn, c),
	}
}

func (m *defaultTaskRecordModel) customCacheKeys(data *TaskRecord) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
