package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gobuffalo/packr"
	"github.com/aws/aws-lambda-go/events"
	"context"
)


type Response struct {
	Message string `json:"message"`
}

var html string
func handleRequest(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: html,
		Headers: map[string]string{"Content-Type": "text/html; charset=UTF-8"},
	}, nil
}

func main() {
	box := packr.NewBox("../statics")
	res, err := box.MustString("index.html")
	if err != nil {
		panic(err)
	}
	html = res
	lambda.Start(handleRequest)
}
