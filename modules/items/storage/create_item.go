package storage

import (
	"context"
	"github.com/server/modules/items/model"
)

func (s *sqlStore) CreateItem(ctx context.Context, data *model.ToDoItemCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
