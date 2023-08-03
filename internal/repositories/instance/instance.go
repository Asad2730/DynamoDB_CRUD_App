package instance

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func GetConnection() (*dynamodb.Client, error) {

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}

	client := dynamodb.NewFromConfig(cfg)
	return client, nil
}
