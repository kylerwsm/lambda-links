package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kylerwsm/lambda-links/pkg/constants"
)

// Handler is our lambda handler to serve static files for the base path.
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{StatusCode: http.StatusFound, Headers: map[string]string{
		"Location": constants.HomePage,
	}}, nil
}

func main() {
	lambda.Start(Handler)
}
