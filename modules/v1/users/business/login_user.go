package business

import (
	"context"
	"github.com/server/modules/v1/users/model"
)

type GetUserStore interface {
	LoginUser(ctx context.Context, cond map[string]interface{}) (*model.UserLoginResponse, error)
}

type loginUser struct {
	store GetUserStore
}

func LoginUserBusiness(store GetUserStore) *loginUser {
	return &loginUser{store: store}
}

func (biz *loginUser) LoginUser(ctx context.Context, account, password string) (*model.UserLoginResponse, error) {
	user, err := biz.store.LoginUser(ctx, map[string]interface{}{"account": account, "password": password})
	if err != nil {
		return nil, err
	}

	return user, nil
}
