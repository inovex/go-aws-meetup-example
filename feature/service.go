package feature

import (
	"example.com/service/models"
)

type itemRepository interface {
	getItemByName(name string) (models.Item, error)
	putItem(item models.Item) error
}

type itemService struct {
	repo itemRepository
}

func newItemService(repo itemRepository) *itemService {
	return &itemService{
		repo: repo,
	}
}

func (service itemService) addItem(item models.Item) error {
	// check if item exists
	_, err := service.repo.getItemByName(item.Name)

	if err == nil {
		// if getItem returns valid item with the same id, fail
		return models.ErrItemAlreadyExists
	}

	// otherwise, try to save the item
	err = service.repo.putItem(item)
	if err != nil {
		return models.ErrWriteFailedWithCause(err)
	}

	// and return nil on success
	return nil
}
