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

// func NewNeptuneClient() (*NeptuneClient, error) {
//     endpoint := fmt.Sprintf("wss://%s:8182/gremlin", config.Config.Neptune.Endpoint)
//     driverRemoteConnection, err := gremlingo.NewDriverRemoteConnection(endpoint,
//         func(settings *gremlingo.DriverRemoteConnectionSettings) {
//             settings.TraversalSource = "g"
//         })
//     if err != nil {
//         return nil, fmt.Errorf("failed to create connection: %w", err)
//     }

//     return &NeptuneClient{
//         connection: driverRemoteConnection,
//     }, nil
// }

func NewNeptuneClient() (*gremlingo.DriverRemoteConnection, error) {
    endpoint := fmt.Sprintf("wss://%s:8182/gremlin", config.Config.Neptune.Endpoint)
    driverRemoteConnection, err := gremlingo.NewDriverRemoteConnection(endpoint,
        func(settings *gremlingo.DriverRemoteConnectionSettings) {
            settings.TraversalSource = "g"
        })
    if err != nil {
        panic(err)
    }

    return driverRemoteConnection, nil
}

func (client *NeptuneClient) Close() error {
    client.connection.Close()
    return nil
}

// AddVertexは、指定されたラベルとプロパティを持つノードをグラフに追加します。
// checkExistenceがtrueの場合、ノードが既に存在するかを確認し、存在する場合は何もしません。
//
// パラメータ:
// - label: ノードのラベル
// - properties: ノードのプロパティを含むマップ
// - checkExistence: ノードが既に存在するかを確認するかどうか
//
// 戻り値:
// - エラーが発生した場合はエラーを返します。成功した場合はnilを返します。
func (client *NeptuneClient) AddVertex(label string, properties map[string]interface{}, checkExistence bool) error {
    g := gremlingo.Traversal_().WithRemote(client.connection)

    if checkExistence {
        // ノードが既に存在するか確認
        existing, err := g.V().HasLabel(label).Has("id", properties["id"]).ToList()
        if err != nil {
            return fmt.Errorf("failed to check vertex existence: %w", err)
        }
        if len(existing) > 0 {
            return nil // 既に存在する場合は何もしない
        }
    }

    // ノードを追加
    promise := g.AddV(label).PropertyMap(properties).Iterate()
    err := <-promise
    if err != nil {
        return fmt.Errorf("failed to add vertex: %w", err)
    }
    return nil
}

// AddEdgeは、指定されたIDのノード間にエッジを追加します。
//
// パラメータ:
// - fromID: エッジの開始ノードのID
// - toID: エッジの終了ノードのID
// - label: エッジのラベル
//
// 戻り値:
// - エラーが発生した場合はエラーを返します。成功した場合はnilを返します。
func (client *NeptuneClient) AddEdge(fromID, toID, label string) error {
    g := gremlingo.Traversal_().WithRemote(client.connection)

    // エッジを追加
    promise := g.V().Has("id", fromID).AddE(label).To(g.V().Has("id", toID)).Iterate()
    err := <-promise
    if err != nil {
        return fmt.Errorf("failed to add edge: %w", err)
    }
    return nil
}
