package storage

import (
	"context"
	"github.com/server/modules/v1/users/model"
)

func (s *sqlStore) LoginUser(ctx context.Context, cond map[string]interface{}) (*model.UserLoginResponse, error) {
	var data *model.UserLoginResponse
	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
