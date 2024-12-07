package neptune

import (
	"fmt"

	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/murasame29/go-httpserver-template/cmd/config"
)

func NewNeptuneClient() (*gremlingo.DriverRemoteConnection, error) {
	endpoint := fmt.Sprintf("wss://%s:8182/gremlin", config.Config.Neptune.Endpoint)
	driverRemoteConnection, err := gremlingo.NewDriverRemoteConnection(endpoint,
		func(settings *gremlingo.DriverRemoteConnectionSettings) {
			settings.TraversalSource = "g"
		})
	if err != nil {
		return nil, fmt.Errorf("failed to create driver remote connection: %w", err)
	}

	return driverRemoteConnection, nil
}
