package interactor

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/framework/contexts"
	"github.com/murasame29/go-httpserver-template/internal/usecase/service"
)

type Room struct {
	_user       *service.User
	_skill      *service.Skill
	_aimSkill   *service.AimSkill
	_room       *service.Room
	_roomMember *service.RoomMember
}

func NewRoom(
	user *service.User,
	skill *service.Skill,
	aimSkill *service.AimSkill,
	room *service.Room,
	roomMember *service.RoomMember,
) *Room {
	return &Room{
		_user:       user,
		_skill:      skill,
		_aimSkill:   aimSkill,
		_room:       room,
		_roomMember: roomMember,
	}
}

type GetRoomResult struct {
	Room    *entity.Room
	AimTags []entity.Skill
	Members []entity.DisplayUser
}

func (i *Room) GetByID(ctx context.Context, roomID string) (*GetRoomResult, error) {
	room, err := i._room.Get(ctx, roomID)
	if err != nil {
		return nil, err
	}

	aimSkills, err := i._aimSkill.List(ctx, roomID)
	if err != nil {
		return nil, err
	}

	aimSkillIDs := make([]int, len(aimSkills))
	for i, aimSkill := range aimSkills {
		aimSkillIDs[i] = aimSkill.SkillID
	}

	skills, err := i._skill.List(ctx, aimSkillIDs)
	if err != nil {
		return nil, err
	}

	members, err := i._roomMember.List(ctx, roomID)
	if err != nil {
		return nil, err
	}

	userIDs := make([]string, len(members))

	for i, member := range members {
		userIDs[i] = member.UserID
	}

	users, err := i._user.List(ctx, userIDs)
	if err != nil {
		return nil, err
	}

	return &GetRoomResult{
		Room:    room,
		AimTags: skills,
		Members: entity.ToDisplayUsers(users),
	}, nil
}

type ListRoomResult struct {
	Rooms []GetRoomResult
}

func (i *Room) List(ctx context.Context) (*ListRoomResult, error) {
	userID := contexts.GetUserID(ctx)
	rooms, err := i._room.List(ctx, userID)
	if err != nil {
		return nil, err
	}

	roomIDs := make([]string, len(rooms))
	for i, room := range rooms {
		roomIDs[i] = room.ID
	}

	aimSkills, skillIDs, err := i._aimSkill.ListByRoomIDs(ctx, roomIDs)
	if err != nil {
		return nil, err
	}

	skills, err := i._skill.List(ctx, skillIDs)
	if err != nil {
		return nil, err
	}

	skillMap := entity.ToSkillMap(skills)

	listMemberResult, err := i._roomMember.ListByRoomIDs(ctx, roomIDs)
	if err != nil {
		return nil, err
	}

	users, err := i._user.List(ctx, listMemberResult.UserIDs)
	if err != nil {
		return nil, err
	}

	userMap := entity.ToUserMap(users)

	var result []GetRoomResult
	for _, room := range rooms {
		skills := make([]entity.Skill, len(aimSkills[room.ID]))
		for i, aimTag := range aimSkills[room.ID] {
			skills[i] = skillMap[aimTag.SkillID]
		}

		members := make([]entity.User, len(listMemberResult.Members[room.ID]))
		for i, member := range listMemberResult.Members[room.ID] {
			members[i] = userMap[member.UserID]
		}

		result = append(result, GetRoomResult{
			Room:    &room,
			AimTags: skills,
			Members: entity.ToDisplayUsers(members),
		})
	}

	return &ListRoomResult{
		Rooms: result,
	}, nil
}

type CreateRoomParam struct {
	Name        string
	Description string
	AimSkills   []string
	CreatedBy   string
}

func (i *Room) Create(ctx context.Context, param CreateRoomParam) (*GetRoomResult, error) {

	roomID, err := i._room.Create(ctx, param.Name, param.Description, param.CreatedBy)
	if err != nil {
		return nil, err
	}

	if err := i._aimSkill.Upsert(ctx, roomID, param.AimSkills); err != nil {
		return nil, err
	}

	if err := i._roomMember.Join(ctx, roomID, param.CreatedBy); err != nil {
		return nil, err
	}

	return i.GetByID(ctx, roomID)
}
