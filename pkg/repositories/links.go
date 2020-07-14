package repositories

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/kylerwsm/lambda-links/pkg/entity"
)

// Declare DynamoDB instance which is safe for concurrent use.
var svc = dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-southeast-1"))

const tableName = "Links"

// GetLink gets the Link entry associated with the specified short link from our database.
func GetLink(shortLink string) (entity.Link, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ShortLink": {
				S: aws.String(shortLink),
			},
		},
		TableName: aws.String(tableName),
	}
	result, err := svc.GetItem(input)
	if err != nil {
		return entity.Link{}, err
	}
	var link entity.Link
	err = dynamodbattribute.UnmarshalMap(result.Item, &link)
	return link, err
}
