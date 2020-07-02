package feature

import (
	"encoding/json"
	"errors"
	"example.com/service/models"
	"fmt"
	"net/http"
)

type itemGetter interface {
	getItems() ([]models.Item, error)
}

func CreateGetItemsHandler(g itemGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := g.getItems()
		if err != nil {
			// TODO logging
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_ = json.NewEncoder(w).Encode(items)
	}
}

type itemAdder interface {
	addItem(item models.Item) error
}

func CreateNewItemHandler(a itemAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var item models.Item
		err := json.NewDecoder(r.Body).Decode(&item)
		if err != nil || !item.Valid() {
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
		}

		err = a.addItem(item)
		if errors.Is(err, models.ErrItemAlreadyExists) {
			http.Error(w, fmt.Sprintf("item with name %s already exists", item.Name), http.StatusConflict)
			return
		} else if errors.Is(err, models.ErrWriteFailed) {
			http.Error(w, "write of valid new item failed for unknown reason", http.StatusInternalServerError)
			return
		} else if err != nil {
			http.Error(w, "unexpected server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
