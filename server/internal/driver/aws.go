package driver

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewAWSConfig() aws.Config {
	awsConfig, err := awsConfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	return awsConfig
}

func NewDynamoDB(awsConfig aws.Config) *dynamodb.Client {
	return dynamodb.NewFromConfig(awsConfig)
}
