package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/phamtrung99/aws_lambda_demo/utils"
	"go.uber.org/zap"
)

func handle(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger, _ := zap.NewDevelopment()
	headers := utils.NewHeaders()
	errRes := events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Headers:    headers,
	}
	fmt.Println("oke")
	logger.Debug("Client token: " + request.Headers["Authorization"])

	return events.APIGatewayProxyResponse{
		Headers:    headers,
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(handle)
}
