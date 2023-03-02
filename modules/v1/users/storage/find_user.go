package storage

import (
	"context"
	"github.com/server/modules/v1/users/model"
)

func (s *sqlStore) GetUser(ctx context.Context, cond map[string]interface{}) (*model.UserResponse, error) {
	var data *model.UserResponse
	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
