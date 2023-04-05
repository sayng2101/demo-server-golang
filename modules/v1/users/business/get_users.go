package business

import (
	"context"
	"github.com/server/modules/v1/users/model"
)

type GetUsersStore interface {
	GetUsers(
		ctx context.Context,
	) ([]model.UserResponse, error)
}

type GetUsers struct {
	store GetUsersStore
}

func GetUsersBusiness(store GetUsersStore) *GetUsers {
	return &GetUsers{store: store}
}

func (biz *GetUsers) GetUsers(
	ctx context.Context,

) ([]model.UserResponse, error) {
	user, err := biz.store.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}
