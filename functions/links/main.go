package main

import (
	"context"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kylerwsm/lambda-links/pkg/services"
	"github.com/kylerwsm/lambda-links/pkg/util"
)

// Handler is our lambda handler for short link paths.
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	path := request.Path
	if len(path) == 0 {
		return util.ResponseNotFound(), nil
	}
	shortLink := path[1:]
	link, err := services.GetLink(shortLink)
	if err != nil {
		return util.ResponseInternalServerError(), err
	}
	if originalLink := link.OriginalLink; len(strings.TrimSpace(originalLink)) > 0 {
		return util.ResponseFound(link.OriginalLink), nil
	}
	return util.ResponseNotFound(), nil
}

func main() {
	lambda.Start(Handler)
}
