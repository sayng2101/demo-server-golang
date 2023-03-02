package business

import (
	"context"

	"github.com/server/modules/v1/users/model"
)

type GetByIdUserStore interface {
	GetUser(ctx context.Context, cond map[string]interface{}) (*model.UserResponse, error)
}

type GetUser struct {
	store GetByIdUserStore
}

func GetByIdUserBusiness(store GetByIdUserStore) *GetUser {
	return &GetUser{store: store}
}

func (biz *GetUser) GetUser(ctx context.Context, id int) (*model.UserResponse, error) {
	user, err := biz.store.GetUser(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return user, nil
}
