// SPDX-FileCopyrightText: 2020 inovex GmbH <https://www.inovex.de>
// 
// SPDX-License-Identifier: MIT
package feature

import (
	"context"
	"example.com/service/models"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/guregu/dynamo"
)

type dynamoRepo struct {
	itemTable dynamo.Table
}

// newRepoWithClient returns an instance of dynamoRepo that is using the provided interface
// for all DB operations. It operates on the table provided by itemTableName.
func newRepoWithClient(itemTableName string, client dynamodbiface.DynamoDBAPI) *dynamoRepo {
	db := dynamo.NewFromIface(client)
	return &dynamoRepo{
		itemTable: db.Table(itemTableName),
	}
}

// newDynamoRepo returns an instance of dynamoRepo that uses a default AWS dynamo client
// and operates on the table provided by itemTableName
func newDynamoRepo(itemTableName string) *dynamoRepo {
	client := dynamodb.New(session.Must(session.NewSessionWithOptions(session.Options{})))
	return newRepoWithClient(itemTableName, client)
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
