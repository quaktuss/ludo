package database

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Client wraps DynamoDB with app-specific helpers.
type Client struct {
	db        *dynamodb.DynamoDB
	tableName string
}

// New returns a DynamoDB client. If accessKey/secretKey are empty, the default
// AWS credential chain (IAM role, env vars, ~/.aws) is used.
func New(region, tableName, accessKeyID, secretAccessKey string) (*Client, error) {
	var creds *credentials.Credentials
	if accessKeyID != "" && secretAccessKey != "" {
		creds = credentials.NewStaticCredentials(accessKeyID, secretAccessKey, "")
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: creds,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %w", err)
	}

	return &Client{
		db:        dynamodb.New(sess),
		tableName: tableName,
	}, nil
}

// TODO: implement CRUD methods (SaveServer, GetServer, ListServers, DeleteServer)
