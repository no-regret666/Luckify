package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ CommentModel = (*customCommentModel)(nil)
var commentOmitColumns = []string{"create_time", "update_time"}

type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
	CommentModel interface {
		commentModel
		customCommentLogicModel
	}

	customCommentModel struct {
		*defaultCommentModel
	}

	customCommentLogicModel interface {
	}
)

// NewCommentModel returns a model for the database table.
func NewCommentModel(conn *gorm.DB, c cache.CacheConf) CommentModel {
	return &customCommentModel{
		defaultCommentModel: newCommentModel(conn, c),
	}
}

func (m *defaultCommentModel) customCacheKeys(data *Comment) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
