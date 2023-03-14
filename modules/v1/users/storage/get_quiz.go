package storage

import (
	"context"
	"github.com/server/modules/v1/users/model"
)

func (s *sqlStore) GetQuiz(
	ctx context.Context,
	filter *model.Filter,
) ([]model.Quiz, error) {
	var data []model.Quiz
	db := s.db.Where("mon <=> ?", filter.Genre)

	if err := s.db.Table(model.Quiz{}.TableName()).Error; err != nil {
		return nil, err
	}
	if err := db.Order("rand()").
		Limit(filter.Limit).
		Order("id desc").
		Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
