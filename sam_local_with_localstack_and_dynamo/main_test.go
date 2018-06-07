package main

import (
	"testing"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/stretchr/testify/assert"
)

type fakeDynamoDB struct {
	dynamodbiface.DynamoDBAPI
	payload []*string
}

func (f *fakeDynamoDB) ListTables(*dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error) {
	return &dynamodb.ListTablesOutput{
		TableNames: f.payload,
	}, nil
}

func TestGetTableNames(t *testing.T) {
	dynamodb := &fakeDynamoDB{
		payload: []*string{
			aws.String("ala"),
			aws.String("ma"),
			aws.String("kota"),
		},
	}

	actual := GetTableNames(dynamodb)
	expected := []string{"ala", "ma", "kota"}
	assert.EqualValues(t, expected, actual)
}
