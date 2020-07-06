package feature

import (
	"example.com/service/models"
	"testing"
)

type mockItemRepo struct {
	itemsByName   map[string]models.Item
	putShouldFail bool
}

func (m mockItemRepo) getItemByName(name string) (models.Item, error) {
	i, ok := m.itemsByName[name]
	if !ok {
		return models.Item{}, models.ErrItemNotFound
	}
	return i, nil
}

func (m *mockItemRepo) putItem(item models.Item) error {
	if m.putShouldFail {
		return models.ErrWriteFailed
	}
	m.itemsByName[item.Name] = item
	return nil
}

func Test_itemService_addItem(t *testing.T) {
	type fields struct {
		repo itemRepository
	}
	type args struct {
		item models.Item
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := itemService{
				repo: tt.fields.repo,
			}
			if err := service.addItem(tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("addItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
