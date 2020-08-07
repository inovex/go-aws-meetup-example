// SPDX-FileCopyrightText: 2020 inovex GmbH <https://www.inovex.de>
// 
// SPDX-License-Identifier: MIT
package feature

import (
	"context"
	"example.com/service/models"
)

type itemRepository interface {
	getItemByName(ctx context.Context, name string) (models.Item, error)
	putItem(ctx context.Context, item models.Item) error
}

type itemService struct {
	repo itemRepository
}

func newItemService(repo itemRepository) *itemService {
	return &itemService{
		repo: repo,
	}
}

func (service itemService) addItem(ctx context.Context, item models.Item) error {
	// check if item exists
	_, err := service.repo.getItemByName(ctx, item.Name)

	if err == nil {
		// if getItem returns valid item with the same id, fail
		return models.ErrItemAlreadyExists
	}

	// otherwise, try to save the item
	err = service.repo.putItem(ctx, item)
	if err != nil {
		return models.ErrWriteFailedWithCause(err)
	}

	// and return nil on success
	return nil
}
