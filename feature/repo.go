package feature

import (
	"context"
	"example.com/service/models"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/guregu/dynamo"
)

type dynaMock struct {
	dynamodbiface.DynamoDBAPI
}

type dynamoRepo struct {
	itemTable dynamo.Table
	mock      dynaMock
}

func newDynamoRepo(itemTableName string) *dynamoRepo {
	mock := dynaMock{}
	db := dynamo.NewFromIface(mock)
	return &dynamoRepo{
		itemTable: db.Table(itemTableName),
		mock:      dynaMock{},
	}
}

func (d dynamoRepo) getItemByName(ctx context.Context, name string) (models.Item, error) {
	var item models.Item
	err := d.itemTable.
		Get("Name", name).
		Index("IndexName").
		One(&item)
	return item, err
}

func (d dynamoRepo) putItem(ctx context.Context, item models.Item) error {
	return d.itemTable.Put(item).Run()
}

func (d dynamoRepo) getItems(ctx context.Context) ([]models.Item, error) {
	var result []models.Item
	err := d.itemTable.Scan().All(&result)
	return result, err
}
