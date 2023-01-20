package storage

import (
	"context"
	"github.com/server/common"
	"github.com/server/modules/items/model"
)

func (s *sqlStore) ListItem(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]model.ToDoItem, error) {
	var result []model.ToDoItem
	db := s.db.Where("status <> ?", "Deleted")

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}
	if err := s.db.Table(model.ToDoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := s.db.Order("id desc").Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
