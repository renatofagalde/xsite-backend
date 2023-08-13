package main

import (
	"fmt"
	"hello-world/model"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var greeting string
	sourceIP := request.RequestContext.Identity.SourceIP

	if sourceIP == "" {
		greeting = "Hello, world!\n"
	} else {
		greeting = fmt.Sprintf("Hello, %s!\n", sourceIP)
	}

	config := model.EnvConfig{}
	erro := env.Parse(&config)
	if erro != nil {
		fmt.Println("Erro while read parsing env variables")
	}
	log.Info(fmt.Sprintf(" helloooo modulo: golang-lambda - %s", config))
	log.Info(fmt.Sprintf(" linux time - %d", time.Now().UnixNano()))
	log.Info("Inside handler - info level message")
	log.Debug("Inside handler - debug level message")

	return events.APIGatewayProxyResponse{
		Body:       greeting,
		StatusCode: 200,
	}, nil
}

func main() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	log.Info("Inside main - info level message")
	log.Debug("Inside main - debug level message")

	lambda.Start(handler)
}
