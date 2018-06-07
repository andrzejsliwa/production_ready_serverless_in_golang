package main

import (
	"log"

	"os"

	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"context"
)

func handleRequest(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	endpoint := os.Getenv("AWS_DYNAMODB_ENDPOINT")

	fmt.Println("All env variables:")
	for _, pair := range os.Environ() {
		fmt.Println(pair)
	}

	sess, err := session.NewSession(&aws.Config{Endpoint: aws.String(endpoint)})
	fmt.Println("ENDPOINT = ", endpoint)

	if err != nil {
		log.Fatal(err)
	}
	svc := dynamodb.New(sess)

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("%s", GetTableNames(svc)),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}



func GetTableNames(db dynamodbiface.DynamoDBAPI) []string {
	out, err := db.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
	var tableNames []string
	for _, name := range out.TableNames {
		tableNames = append(tableNames, *name)
	}
	return tableNames
}




