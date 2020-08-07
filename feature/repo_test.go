// SPDX-FileCopyrightText: 2020 inovex GmbH <https://www.inovex.de>
// 
// SPDX-License-Identifier: MIT
package feature

import (
	"context"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/tj/assert"
	"testing"
)

type mockItemDynamoClient struct {
	dynamodbiface.DynamoDBAPI
	calledIndexName string
}

func (m *mockItemDynamoClient) Query(input *dynamodb.QueryInput) (*dynamodb.GetItemOutput, error) {
	m.calledIndexName = *input.IndexName
	return nil, nil
}

// This is not a good test, but it exemplifies one of the use cases of interface composition
func Test_dynamoRepo_getItemByName_indexNameCorrect(t *testing.T) {
	client := &mockItemDynamoClient{}
	repo := newRepoWithClient("items", client)
	_, _ = repo.getItemByName(context.Background(), "some item")
	assert.Equal(t, "IndexName", client.calledIndexName)
}
