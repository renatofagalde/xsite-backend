package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/goombaio/namegenerator"
	"time"
	"xsite/infra/api"
)

func main() {
	lambda.Start(handler)
}

type User struct {
	ID   int64
	Nome string
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	return api.APIResponse(200, User{time.Now().Unix(), nameGenerator.Generate()})
}
