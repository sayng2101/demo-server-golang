package business

import (
	"context"
	"errors"
	"github.com/server/modules/v1/users/model"
	"strings"
)

type CreateUserStore interface {
	CreateUser(ctx context.Context, data *model.CreateUser) error
}

type createUserStore struct {
	store CreateUserStore
}

func NewCreateUserBusiness(store CreateUserStore) *createUserStore {
	return &createUserStore{store: store}
}

func (biz *createUserStore) CreateNewUSer(ctx context.Context, data *model.CreateUser) error {
	account := strings.TrimSpace(data.Account)
	password := strings.TrimSpace(data.Password)
	hoten := strings.TrimSpace(data.Hoten)
	//avatar := strings.TrimSpace(data.Avatar)
	if account == "" {
		return errors.New("khong duoc de trong")
	}
	if password == "" {
		return errors.New("khong duoc de trong")
	}
	if hoten == "" {
		return errors.New("khong duoc de trong")
	}
	//if avatar == "" {
	//	return errors.New("khong duoc de trong")
	//}
	if err := biz.store.CreateUser(ctx, data); err != nil {
		return err
	}
	return nil
}
