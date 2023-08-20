package main

import (
	"context"
	"fmt"
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
	nameGenerator := fmt.Sprintf("Random: %s", namegenerator.NewNameGenerator(seed).Generate())

	return api.APIResponse(200, User{time.Now().Unix(), nameGenerator})
}
