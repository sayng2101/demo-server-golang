package business

import (
	"context"

	"github.com/server/modules/v1/users/model"
)

type UpdateByIdUserStore interface {
	GetUser(ctx context.Context, cond map[string]interface{}) (*model.UserResponse, error)
	UpdateUser(ctx context.Context, cond map[string]interface{}, dataUpdate *model.UpdateUser) error
}

type UpdateUser struct {
	store UpdateByIdUserStore
}

func UpdateByIdUserBusiness(store UpdateByIdUserStore) *UpdateUser {
	return &UpdateUser{store: store}
}

func (biz *UpdateUser) UpdateUserById(ctx context.Context, id int, dataUpdate *model.UpdateUser) error {
	data, err := biz.store.GetUser(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if data.Id != data.Id {
		return err
	}
	if err := biz.store.UpdateUser(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return err
	}
	return nil
}
