package business

import (
	"context"
	"github.com/server/modules/v1/users/model"
)

type UpdateUserStore interface {
	GetUser(ctx context.Context, cond map[string]interface{}) (*model.User, error)
	UpdateUser(ctx context.Context, cond map[string]interface{}, dataUpdate *model.UpdateUser) error
}

type updateUser struct {
	store UpdateUserStore
}

func UpdateUserBusiness(store UpdateUserStore) *updateUser {
	return &updateUser{store: store}
}

func (biz *updateUser) UpdateUser(ctx context.Context, id int, dataUpdate *model.UpdateUser) error {
	data, err := biz.store.GetUser(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if data != nil {
		return err
	}
	if err := biz.store.UpdateUser(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return nil
	}
	return nil
}
