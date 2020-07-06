package feature

import (
	"errors"
	"example.com/service/models"
	"github.com/tj/assert"
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
		repo *mockItemRepo
	}
	type args struct {
		item models.Item
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		wantErr         error
		wantRepoContent []models.Item
	}{
		{
			name: "normal non-conflicting update works",
			fields: fields{
				repo: &mockItemRepo{
					itemsByName: make(map[string]models.Item),
				},
			},
			args: args{
				item: models.Item{
					ID:    "item_123",
					Name:  "Item Name",
					Price: 49.99,
				},
			},
			wantRepoContent: []models.Item{
				{
					ID:    "item_123",
					Name:  "Item Name",
					Price: 49.99,
				},
			},
		},
		{
			name: "conflicting update returns already exists error and does not change the repo",
			fields: fields{
				repo: &mockItemRepo{
					itemsByName: map[string]models.Item{
						"Item Name": {
							ID:    "item_123",
							Name:  "Item Name",
							Price: 49.99,
						},
					},
				},
			},
			args: args{
				item: models.Item{
					ID:    "item_456",
					Name:  "Item Name",
					Price: 69.99,
				},
			},
			wantErr: models.ErrItemAlreadyExists,
			wantRepoContent: []models.Item{
				{
					ID:    "item_123",
					Name:  "Item Name",
					Price: 49.99,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := itemService{
				repo: tt.fields.repo,
			}
			err := service.addItem(tt.args.item)
			assertErrorIs(t, tt.wantErr, err)
			assertContainsExactly(t, tt.fields.repo.itemsByName, tt.wantRepoContent)
		})
	}
}

func assertErrorIs(t *testing.T, expected error, actual error) {
	if !errors.Is(actual, expected) {
		t.Errorf("expected any error in the actual error chain to match the expected error")
	}
}

func assertContainsExactly(t *testing.T, m map[string]models.Item, values []models.Item) {
	assert.Equal(t, len(m), len(values), "item count should match expected")
	for _, value := range values {
		existing, ok := m[value.Name]
		assert.True(t, ok, "item should be there")
		assert.Equal(t, value, existing)
	}
}
