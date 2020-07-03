package feature

import (
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

func (d dynamoRepo) getItemByName(id string) (models.Item, error) {
	panic("implement me")
}

func (d dynamoRepo) putItem(item models.Item) error {
	panic("implement me")
}

func (d dynamoRepo) getItems() ([]models.Item, error) {
	panic("implement me")
}
