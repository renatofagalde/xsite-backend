package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"xsite/infra/api"
)

func main() {
	lambda.Start(handler)
}

type User struct {
	ID   uint
	nome string
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return api.APIResponse(200, User{3, "nome"})
}
