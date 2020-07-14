package util

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// ResponseInternalServerError creates a internal server error response event.
func ResponseInternalServerError() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}
}

// ResponseFound creates a redirect response event.
func ResponseFound(location string) events.APIGatewayProxyResponse {
	headers := map[string]string{"Location": location}
	return events.APIGatewayProxyResponse{StatusCode: http.StatusFound, Headers: headers}
}

// ResponseNotFound creates a response for a not found event.
func ResponseNotFound() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
		Body: GetNotFoundHTML(),
	}
}
