package business

import (
	"context"
	"github.com/server/common"
	"github.com/server/modules/items/model"
)

type ListItemStore interface {
	ListItem(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]model.ToDoItem, error)
}

type listStore struct {
	store ListItemStore
}

func ListItemBusiness(store ListItemStore) *listStore {
	return &listStore{store: store}
}

func (biz *listStore) ListItem(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
) ([]model.ToDoItem, error) {
	data, err := biz.store.ListItem(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return data, nil
}
