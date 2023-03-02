package storage

import (
	"context"

	"github.com/server/modules/v1/users/model"
)

func (s *sqlStore) UpdateUser(ctx context.Context, cond map[string]interface{}, dataUpdate *model.UpdateUser) error {
	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}
