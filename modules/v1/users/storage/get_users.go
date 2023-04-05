package storage

import (
	"context"
	"github.com/server/modules/v1/users/model"
)

func (s *sqlStore) GetUsers(
	ctx context.Context,
) ([]model.UserResponse, error) {
	var data []model.UserResponse
	if err := s.db.Order("score desc").Limit(10).Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
