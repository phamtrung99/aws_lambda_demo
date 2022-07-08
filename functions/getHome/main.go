package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/phamtrung99/aws_lambda_demo/services"
	"github.com/phamtrung99/aws_lambda_demo/utils"

	"go.uber.org/zap"
)

func handle(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger, _ := zap.NewDevelopment()
	headers := utils.NewHeaders()
	secretKey := os.Getenv("JWT_SECRET_KEY")
	logger.Debug("HOME secretKey: " + secretKey)
	logger.Sugar().Info(request)

	errRes := events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Headers:    headers,
	}

	clientToken := strings.TrimSpace(strings.TrimPrefix(request.Headers["Authorization"], "Bearer "))
	logger.Debug("HOME Client token: " + clientToken)

	tokenService := services.NewTokenService(secretKey)
	payload, err := tokenService.Decode(clientToken)
	if err != nil {
		errRes.Body = err.Error()
		errRes.StatusCode = http.StatusBadRequest
		return errRes, nil
	}

	body, err := json.Marshal(payload)
	return events.APIGatewayProxyResponse{
		Headers:    headers,
		StatusCode: http.StatusOK,
		Body:       string(body),
	}, nil
}

func main() {
	lambda.Start(handle)
}
