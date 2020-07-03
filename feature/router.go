package feature

import "github.com/go-chi/chi"

type Configuration interface {
	ItemTableName() string
}

type Feature struct {
	service *itemService
	repo    *dynamoRepo
}

func Configure(cfg Configuration) Feature {
	repo := newDynamoRepo(cfg.ItemTableName())
	return Feature{
		service: newItemService(repo),
		repo:    repo,
	}
}

func (feat Feature) RouteExample(r chi.Router) {
	r.Get("/items", CreateGetItemsHandler(feat.repo))
	r.Post("/items", CreateNewItemHandler(feat.service))
}
