package business

import (
	"context"
	"github.com/server/modules/v1/users/model"
)

type GetQuizStore interface {
	GetQuiz(
		ctx context.Context,
		filter *model.Filter,
	) ([]model.Quiz, error)
}

type GetQuiz struct {
	store GetQuizStore
}

func GetQuizBusiness(store GetQuizStore) *GetQuiz {
	return &GetQuiz{store: store}
}

func (biz *GetQuiz) GetQuiz(
	ctx context.Context,
	filter *model.Filter,

) ([]model.Quiz, error) {
	user, err := biz.store.GetQuiz(ctx, filter)
	if err != nil {
		return nil, err
	}

	return user, nil
}
