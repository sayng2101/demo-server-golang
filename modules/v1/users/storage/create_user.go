package storage

import (
	"context"
	"github.com/server/modules/v1/users/model"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *model.CreateUser) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
