package storage

import (
	"context"

	"github.com/server/modules/v1/users/model"
)

func (s *sqlStore) UpdateScoreUserById(ctx context.Context, cond map[string]interface{}, dataUpdate *model.UpdateScoreUser) error {
	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}
