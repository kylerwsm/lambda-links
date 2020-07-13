package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kylerwsm/kyler-bot/pkg/services"
)

// Handler is our lambda handler.
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	shortLink := request.Path[1:]
	link, err := services.GetLink(shortLink)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}
	headers := map[string]string{"Location": link.OriginalLink}
	return events.APIGatewayProxyResponse{StatusCode: http.StatusFound, Headers: headers}, nil
}

func main() {
	lambda.Start(Handler)
}
