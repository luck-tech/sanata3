package neptune

import (
	"context"
	"fmt"

	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/murasame29/go-httpserver-template/cmd/config"
)

var __ = gremlingo.T__

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

// when create new user
func AddUserNode(ctx context.Context, userID string) error {
	driverRemoteConnection, err := NewNeptuneClient()
	if err != nil {
		return fmt.Errorf("failed to create neptune client: %w", err)
	}
	defer driverRemoteConnection.Close()

	g := gremlingo.Traversal_().WithRemote(driverRemoteConnection)

	_, err = g.AddV("user").Property("id", userID).Next()
	if err != nil {
		return fmt.Errorf("failed to add user node: %w", err)
	}

	return nil
}

// when update user
func UpdateUserSkillEdge(ctx context.Context, userID string, havingSkills []string, wantingSkills []string) error {
	driverRemoteConnection, err := NewNeptuneClient()
	if err != nil {
		return fmt.Errorf("failed to create neptune client: %w", err)
	}
	defer driverRemoteConnection.Close()

	g := gremlingo.Traversal_().WithRemote(driverRemoteConnection)

	skills := append(havingSkills, wantingSkills...)
	for _, skill := range skills {
		_, err = g.AddV("skill").Property("name", skill).Next()
		if err != nil {
			return fmt.Errorf("failed to add skill node: %w", err)
		}
	}

	_, err = g.V().HasLabel("user").Has("id", userID).OutE("has").Drop().Next()
	if err != nil {
		return fmt.Errorf("failed to drop has edge: %w", err)
	}
	for _, skill := range havingSkills {
		_, err = g.V().HasLabel("user").Has("id", userID).AddE("has").To(g.V().HasLabel("skill").Has("name", skill)).Next()
		if err != nil {
			return fmt.Errorf("failed to add has edge: %w", err)
		}
	}

	_, err = g.V().HasLabel("user").Has("id", userID).OutE("wants").Drop().Next()
	if err != nil {
		return fmt.Errorf("failed to drop wants edge: %w", err)
	}
	for _, skill := range wantingSkills {
		_, err = g.V().HasLabel("user").Has("id", userID).AddE("wants").To(g.V().HasLabel("skill").Has("name", skill)).Next()
		if err != nil {
			return fmt.Errorf("failed to add wants edge: %w", err)
		}
	}

	return nil
}

// when create room
func CreateRoomNodeAndEdge(ctx context.Context, roomID string, creatorID string, relatingSkills []string) error {
	driverRemoteConnection, err := NewNeptuneClient()
	if err != nil {
		return fmt.Errorf("failed to create neptune client: %w", err)
	}
	defer driverRemoteConnection.Close()

	g := gremlingo.Traversal_().WithRemote(driverRemoteConnection)

	_, err = g.AddV("room").Property("id", roomID).Next()
	if err != nil {
		return fmt.Errorf("failed to add room node: %w", err)
	}

	_, err = g.V().HasLabel("user").Has("id", creatorID).AddE("create").To(g.V().HasLabel("room").Has("id", roomID)).Next()
	if err != nil {
		return fmt.Errorf("failed to add create edge: %w", err)
	}

	for _, skill := range relatingSkills {
		_, err = g.AddV("skill").Property("name", skill).Next()
		if err != nil {
			return fmt.Errorf("failed to add skill node: %w", err)
		}

		_, err = g.V().HasLabel("room").Has("id", roomID).AddE("relates").To(g.V().HasLabel("skill").Has("name", skill)).Next()
		if err != nil {
			return fmt.Errorf("failed to add relates edge: %w", err)
		}
	}

	return nil
}

// when update room
func UpdateRoomSkillEdge(ctx context.Context, roomID string, relatingSkills []string) error {
	driverRemoteConnection, err := NewNeptuneClient()
	if err != nil {
		return fmt.Errorf("failed to create neptune client: %w", err)
	}
	defer driverRemoteConnection.Close()

	g := gremlingo.Traversal_().WithRemote(driverRemoteConnection)

	_, err = g.V().HasLabel("room").Has("id", roomID).OutE("relates").Drop().Next()
	if err != nil {
		return fmt.Errorf("failed to drop relates edge: %w", err)
	}

	for _, skill := range relatingSkills {
		_, err = g.AddV("skill").Property("name", skill).Next()
		if err != nil {
			return fmt.Errorf("failed to add skill node: %w", err)
		}

		_, err = g.V().HasLabel("room").Has("id", roomID).AddE("relates").To(g.V().HasLabel("skill").Has("name", skill)).Next()
		if err != nil {
			return fmt.Errorf("failed to add relates edge: %w", err)
		}
	}

	return nil
}

// when delete room
func DeleteRoomNodeAndEdge(ctx context.Context, roomID string) error {
	driverRemoteConnection, err := NewNeptuneClient()
	if err != nil {
		return fmt.Errorf("failed to create neptune client: %w", err)
	}
	defer driverRemoteConnection.Close()

	g := gremlingo.Traversal_().WithRemote(driverRemoteConnection)

	// 辺も一緒に消える
	_, err = g.V().HasLabel("room").Has("id", roomID).Drop().Next()
	if err != nil {
		return fmt.Errorf("failed to drop room node: %w", err)
	}

	return nil
}

// when join room
func AddRoomMemberEdge(ctx context.Context, roomID string, memberID string) error {
	driverRemoteConnection, err := NewNeptuneClient()
	if err != nil {
		return fmt.Errorf("failed to create neptune client: %w", err)
	}
	defer driverRemoteConnection.Close()

	g := gremlingo.Traversal_().WithRemote(driverRemoteConnection)

	_, err = g.V().HasLabel("user").Has("id", memberID).AddE("participatesIn").To(g.V().HasLabel("room").Has("id", roomID)).Next()
	if err != nil {
		return fmt.Errorf("failed to add participatesIn edge: %w", err)
	}

	return nil
}

// when leave room
func DeleteRoomMemberEdge(ctx context.Context, roomID string, memberID string) error {
	driverRemoteConnection, err := NewNeptuneClient()
	if err != nil {
		return fmt.Errorf("failed to create neptune client: %w", err)
	}
	defer driverRemoteConnection.Close()

	g := gremlingo.Traversal_().WithRemote(driverRemoteConnection)

	_, err = g.V().HasLabel("user").Has("id", memberID).OutE("member_of").Where(__.OtherV().HasLabel("room").Has("id", roomID)).Drop().Next()
	if err != nil {
		return fmt.Errorf("failed to drop participatesIn edge: %w", err)
	}

	return nil
}
