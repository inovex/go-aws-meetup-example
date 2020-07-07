package feature

import (
	"context"
	"encoding/json"
	"errors"
	"example.com/service/logger"
	"example.com/service/models"
	"fmt"
	"net/http"
)

type itemGetter interface {
	getItems(context.Context) ([]models.Item, error)
}

func CreateGetItemsHandler(g itemGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := logger.FromRequest(r)
		log.Info("retrieving items")
		items, err := g.getItems(r.Context())
		if err != nil {
			const msg = "unexpected server error"
			log.Errorw(msg, "error", err)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
		_ = json.NewEncoder(w).Encode(items)
	}
}

type itemAdder interface {
	addItem(ctx context.Context, item models.Item) error
}

func CreateNewItemHandler(a itemAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var item models.Item
		err := json.NewDecoder(r.Body).Decode(&item)
		log := logger.FromRequest(r).With("item", item)
		ctx := logger.AddLoggerToContext(r.Context(), log)

		if err != nil || !item.Valid() {
			log.Warn("invalid input")
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
		}

		log.Info("trying to add item")
		err = a.addItem(ctx, item)
		if errors.Is(err, models.ErrItemAlreadyExists) {
			const msg = "item with the same name already exists"
			log.Warn(msg, "name", item.Name)
			http.Error(w, fmt.Sprintf("%s: %s", msg, item.Name), http.StatusConflict)
			return
		} else if errors.Is(err, models.ErrWriteFailed) {
			const msg = "write of valid new item failed for unknown reason"
			log.Errorw(msg, "error", err)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		} else if err != nil {
			const msg = "unexpected server error"
			log.Errorw(msg, "error", err)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
