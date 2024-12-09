package neptune

import (
	"context"
	"fmt"
	"sort"

	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/murasame29/go-httpserver-template/cmd/config"
)

var __ = gremlingo.T__
var P = gremlingo.P

const (
	WantsWeight       = 0.6
	HasWeight         = 0.3
	RoomWeight        = 0.1
	MaxRecommendRooms = 10
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

// when get recommend rooms
func GetRecommendRoomNodes(ctx context.Context, userID string) ([]string, error) {
	driverRemoteConnection, err := NewNeptuneClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create neptune client: %w", err)
	}
	defer driverRemoteConnection.Close()

	g := gremlingo.Traversal_().WithRemote(driverRemoteConnection)

	// ユーザーの技術・資格を取得
	wantsSkills, err := getUserSkills(g, userID, "wants")
	if err != nil {
		return nil, fmt.Errorf("failed to get wants skills: %w", err)
	}
	hasSkills, err := getUserSkills(g, userID, "has")
	if err != nil {
		return nil, fmt.Errorf("failed to get has skills: %w", err)
	}

	// ユーザーが所属しているルームの技術・資格を取得
	roomSkills, err := getUserRoomSkills(g, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user room skills: %w", err)
	}

	// 関連するルームを取得
	relatedRooms, err := getRelatedRooms(g, wantsSkills, hasSkills, roomSkills)
	if err != nil {
		return nil, fmt.Errorf("failed to get related rooms: %w", err)
	}

	// ユーザーが既に参加しているルームを除外
	participatingRooms, err := getUserParticipatingRooms(g, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get participating rooms: %w", err)
	}
	createdRooms, err := getUserCreatedRooms(g, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get created rooms: %w", err)
	}

	// フィルタリング
	recommendedRooms := filterRooms(relatedRooms, participatingRooms, createdRooms)

	// 重みを計算してソート
	weightedRooms := calculateWeights(g, recommendedRooms, wantsSkills, hasSkills, roomSkills)
	sort.Slice(weightedRooms, func(i, j int) bool {
		return weightedRooms[i].Weight > weightedRooms[j].Weight
	})

	// ルームIDのリストを返却
	roomIDs := make([]string, 0, MaxRecommendRooms)
	for i, room := range weightedRooms {
		if i >= MaxRecommendRooms {
			break
		}
		roomIDs = append(roomIDs, room.ID)
	}

	return roomIDs, nil
}

// GetRecommendRoomNodes function's utils below
type Room struct {
	ID     string
	Weight float64
}

func getUserSkills(g *gremlingo.GraphTraversalSource, userID, edgeLabel string) ([]string, error) {
	skills := []string{}
	result, err := g.V().HasLabel("user").Has("id", userID).Out(edgeLabel).Values("name").ToList()
	if err != nil {
		return nil, err
	}
	for _, skill := range result {
		skills = append(skills, skill.GetString())
	}
	return skills, nil
}

func getUserRoomSkills(g *gremlingo.GraphTraversalSource, userID string) ([]string, error) {
	skills := []string{}
	result, err := g.V().HasLabel("user").Has("id", userID).Out("member_of", "creates").In("relatedTo").Values("name").ToList()
	if err != nil {
		return nil, err
	}
	for _, skill := range result {
		skills = append(skills, skill.GetString())
	}
	return skills, nil
}

func getRelatedRooms(g *gremlingo.GraphTraversalSource, wantsSkills, hasSkills, roomSkills []string) ([]Room, error) {
	rooms := []Room{}
	skillNames := append(wantsSkills, hasSkills...)
	skillNames = append(skillNames, roomSkills...)
	skillNamesInterface := make([]interface{}, len(skillNames))
	for i, v := range skillNames {
		skillNamesInterface[i] = v
	}
	result, err := g.V().HasLabel("skill").Has("name", P.Within(skillNamesInterface...)).In("relatedTo").HasLabel("room").Values("id").ToList()
	if err != nil {
		return nil, err
	}
	for _, room := range result {
		rooms = append(rooms, Room{ID: room.GetString()})
	}
	return rooms, nil
}

func getUserParticipatingRooms(g *gremlingo.GraphTraversalSource, userID string) ([]Room, error) {
	rooms := []Room{}
	result, err := g.V().HasLabel("user").Has("id", userID).Out("participatesIn").Values("id").ToList()
	if err != nil {
		return nil, err
	}
	for _, room := range result {
		rooms = append(rooms, Room{ID: room.GetString()})
	}
	return rooms, nil
}

func getUserCreatedRooms(g *gremlingo.GraphTraversalSource, userID string) ([]Room, error) {
	rooms := []Room{}
	result, err := g.V().HasLabel("user").Has("id", userID).Out("created").Values("id").ToList()
	if err != nil {
		return nil, err
	}
	for _, room := range result {
		rooms = append(rooms, Room{ID: room.GetString()})
	}
	return rooms, nil
}

func filterRooms(relatedRooms, participatingRooms, createdRooms []Room) []Room {
	participatingRoomIDs := make(map[string]struct{})
	for _, room := range participatingRooms {
		participatingRoomIDs[room.ID] = struct{}{}
	}
	createdRoomIDs := make(map[string]struct{})
	for _, room := range createdRooms {
		createdRoomIDs[room.ID] = struct{}{}
	}

	filteredRooms := []Room{}
	for _, room := range relatedRooms {
		if _, participating := participatingRoomIDs[room.ID]; !participating {
			if _, created := createdRoomIDs[room.ID]; !created {
				filteredRooms = append(filteredRooms, room)
			}
		}
	}
	return filteredRooms
}

func calculateWeights(g *gremlingo.GraphTraversalSource, rooms []Room, wantsSkills, hasSkills, roomSkills []string) []Room {
	for i, room := range rooms {
		wantsCount := 0
		hasCount := 0
		roomCount := 0
		// ルームの関連技術・資格を取得
		roomSkills, err := getRoomSkills(g, room.ID)
		if err != nil {
			continue
		}
		for _, skill := range roomSkills {
			if contains(wantsSkills, skill) {
				wantsCount++
			}
			if contains(hasSkills, skill) {
				hasCount++
			}
			if contains(roomSkills, skill) {
				roomCount++
			}
		}
		rooms[i].Weight = WantsWeight*float64(wantsCount) + HasWeight*float64(hasCount) + RoomWeight*float64(roomCount)
	}
	return rooms
}

func getRoomSkills(g *gremlingo.GraphTraversalSource, roomID string) ([]string, error) {
	skills := []string{}
	result, err := g.V().HasLabel("room").Has("id", roomID).Out("relatedTo").Values("name").ToList()
	if err != nil {
		return nil, err
	}
	for _, skill := range result {
		skills = append(skills, skill.GetString())
	}
	return skills, nil
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
