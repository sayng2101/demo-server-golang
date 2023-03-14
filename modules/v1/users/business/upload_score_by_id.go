package business

import (
	"context"
	"github.com/server/modules/v1/users/model"
)

type UpdateScoreByIdUserStore interface {
	GetUser(ctx context.Context, cond map[string]interface{}) (*model.UserResponse, error)
	UpdateScoreUserById(ctx context.Context, cond map[string]interface{}, dataUpdate *model.UpdateScoreUser) error
}

type UpdateScoreUser struct {
	store UpdateScoreByIdUserStore
}

func UpdateScoreByIdUserBusiness(store UpdateScoreByIdUserStore) *UpdateScoreUser {
	return &UpdateScoreUser{store: store}
}

func (biz *UpdateScoreUser) UpdateScoreUser(ctx context.Context, id int, dataUpdate *model.UpdateScoreUser) error {
	data, err := biz.store.GetUser(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if data.Id != data.Id {
		return err
	}
	if err := biz.store.UpdateScoreUserById(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return err
	}
	return nil
}
