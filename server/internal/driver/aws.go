package driver

import (
	"context"
	"fmt"

	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/murasame29/go-httpserver-template/cmd/config"
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

type NeptuneClient struct {
	connection *gremlingo.DriverRemoteConnection
}

func NewNeptuneClient() (*NeptuneClient, error) {
	endpoint := fmt.Sprintf("wss://%s:8182/gremlin", config.Config.Neptune.Endpoint)
	driverRemoteConnection, err := gremlingo.NewDriverRemoteConnection(endpoint,
		func(settings *gremlingo.DriverRemoteConnectionSettings) {
			settings.TraversalSource = "g"
		})
	if err != nil {
		return nil, fmt.Errorf("failed to create connection: %w", err)
	}

	return &NeptuneClient{
		connection: driverRemoteConnection,
	}, nil
}

func (client *NeptuneClient) GetTraversal() *gremlingo.GraphTraversalSource {
	return gremlingo.Traversal_().WithRemote(client.connection)
}
