package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/utils"
	"go.uber.org/zap"
)

func handle(ctx context.Context, request events.APIGatewayProxyResponse) (events.APIGatewayProxyResponse, error) {
	logger, _ := zap.NewDevelopment()
	headers := utils.NewHeaders()
	// errRes := events.APIGatewayProxyResponse{
	// 	StatusCode: http.StatusInternalServerError,
	// 	Headers:    headers,
	// }
	logger.Debug("Client token: " + request.Headers["Authorization"])

	return events.APIGatewayProxyResponse{
		Headers:    headers,
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(handle)
}
