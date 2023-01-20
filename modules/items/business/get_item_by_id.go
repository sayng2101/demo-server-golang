package business

import (
	"context"
	"github.com/server/modules/items/model"
)

type GetItemStore interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.ToDoItem, error)
}

type getStore struct {
	store GetItemStore
}

func GetItemBusiness(store GetItemStore) *getStore {
	return &getStore{store: store}
}

func (biz *getStore) GetItemById(ctx context.Context, id int) (*model.ToDoItem, error) {

	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}
