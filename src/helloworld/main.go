package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func helloworld(ctx context.Context,
	e events.APIGatewayV2HTTPRequest) (
	events.APIGatewayV2HTTPResponse, error) {

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       string("helloworld"),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(helloworld)
}
