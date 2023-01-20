package storage

import (
	"context"
	"github.com/server/modules/items/model"
)

func (s *sqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*model.ToDoItem, error) {
	var data model.ToDoItem
	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
