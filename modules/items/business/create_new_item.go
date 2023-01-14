package business

import (
	"context"
	"errors"
	"github.com/server/modules/items/model"
	"strings"
)

type CreateItemStore interface {
	CreateItem(ctx context.Context, data *model.ToDoItemCreate) error
}

type createStore struct {
	store CreateItemStore
}

func NewCreateBusiness(store CreateItemStore) *createStore {
	return &createStore{store: store}
}

func (biz *createStore) CreateNewItem(ctx context.Context, data *model.ToDoItemCreate) error {

	title := strings.TrimSpace(data.Title)
	if title == "" {
		return errors.New("khong duoc de trong")
	}
	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}
	return nil
}
