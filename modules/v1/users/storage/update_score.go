package storage

import (
	"context"
	"gorm.io/gorm"

	"github.com/server/modules/v1/users/model"
)

func (s *sqlStore) UpdateScoreUserById(ctx context.Context, cond map[string]interface{}, dataUpdate *model.UpdateScoreUser) error {
	if err := s.db.Model(&model.User{}).Where(cond).Updates(map[string]interface{}{"score": gorm.Expr("score + ?", dataUpdate.Score)}).Error; err != nil {
		return err
	}
	return nil
}
