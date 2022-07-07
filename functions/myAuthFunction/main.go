package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/phamtrung99/aws_lambda_demo/services"
	"go.uber.org/zap"
)

func generatePolicy(principalId, effect, resource string) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalId}

	if effect != "" && resource != "" {
		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		}
	}

	// Optional output with custom properties of the String, Number or Boolean type.
	authResponse.Context = map[string]interface{}{
		"stringKey":  "stringval",
		"numberKey":  123,
		"booleanKey": true,
	}
	return authResponse
}

func handle(ctx context.Context, request events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	logger, _ := zap.NewDevelopment()
	secretKey := os.Getenv("JWT_SECRET_KEY")

	// clientToken := strings.TrimSpace(strings.TrimPrefix(request.Headers["Authorization"], "Bearer "))
	clientToken := request.AuthorizationToken
	logger.Debug("Client token: " + clientToken)

	tokenService := services.NewTokenService(secretKey)
	_, err := tokenService.Decode(clientToken)
	logger.Debug("Error: " + err.Error())
	if err != nil {
		logger.Sugar().Panic(err)
		return events.APIGatewayCustomAuthorizerResponse{}, err
	}

	return generatePolicy("user", "Allow", request.MethodArn), nil
}

func main() {
	lambda.Start(handle)
}
