package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"encoding/json"
)

type Response struct {
	Message string `json:"message"`
}

func Handler(_ events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body, err := json.Marshal(&Response{Message: "Hello Manning!!"})
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Unprocessable Entity", StatusCode: 422}, err
	}
	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
