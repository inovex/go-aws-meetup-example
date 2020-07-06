package feature

import (
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := itemService{
				repo: tt.fields.repo,
			}
			err := service.addItem(tt.args.item)
			assert.IsType(t, tt.wantErr, err, "expected and returned errors should match")
			assertContainsExactly(t, tt.fields.repo.itemsByName, tt.wantRepoContent)
		})
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
