package interactor

import (
	"context"
	"fmt"
	"time"

	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/framework/contexts"
	"github.com/murasame29/go-httpserver-template/internal/framework/serrors"
	"github.com/murasame29/go-httpserver-template/internal/usecase/service"
)

type Chat struct {
	_room       *service.Room
	_roomMember *service.RoomMember
	_user       *service.User
	_chat       *service.Chat
}

func NewChat(
	room *service.Room,
	roomMember *service.RoomMember,
	user *service.User,
	chat *service.Chat,
) *Chat {
	return &Chat{
		_room:       room,
		_roomMember: roomMember,
		_user:       user,
		_chat:       chat,
	}
}

type GetChatParam struct {
	RoomID   string
	LastTime time.Time
}

func (i *Chat) Get(ctx context.Context, param GetChatParam) ([]entity.Chat, error) {
	userID := contexts.GetUserID(ctx)
	found, err := i._roomMember.Find(ctx, param.RoomID, userID)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, serrors.ErrPermissionNotFound
	}

	chats, err := i._chat.Get(ctx, param.RoomID, param.LastTime)
	if err != nil {
		return nil, err
	}

	fmt.Println(chats)
	return chats, nil
}

type PostChatParam struct {
	RoomID  string
	Message string
}

func (i *Chat) Post(ctx context.Context, param PostChatParam) error {
	userID := contexts.GetUserID(ctx)
	found, err := i._roomMember.Find(ctx, param.RoomID, userID)
	if err != nil {
		return err
	}

	if !found {
		return serrors.ErrPermissionNotFound
	}

	if err := i._chat.Post(ctx, param.RoomID, userID, param.Message); err != nil {
		return err
	}

	return nil
}

type UpdateChatParam struct {
	RoomID  string
	ChatID  string
	Message string
}

func (i *Chat) Edit(ctx context.Context, param UpdateChatParam) error {
	userID := contexts.GetUserID(ctx)
	found, err := i._roomMember.Find(ctx, param.RoomID, userID)
	if err != nil {
		return err
	}

	if !found {
		return serrors.ErrPermissionNotFound
	}

	chat, err := i._chat.GetbyID(ctx, param.ChatID)
	if err != nil {
		return err
	}

	if chat.AutherID != userID {
		return serrors.ErrPermissionNotFound
	}

	chat.Message = param.Message
	chat.UpdatedAt = time.Now()
	if err := i._chat.Update(ctx, chat); err != nil {
		return err
	}

	return nil
}

type DeleteChatParam struct {
	RoomID string
	ChatID string
}

func (i *Chat) Delete(ctx context.Context, param DeleteChatParam) error {
	userID := contexts.GetUserID(ctx)
	found, err := i._roomMember.Find(ctx, param.RoomID, userID)
	if err != nil {
		return err
	}

	if !found {
		return serrors.ErrPermissionNotFound
	}

	chat, err := i._chat.GetbyID(ctx, param.ChatID)
	if err != nil {
		return err
	}

	if chat.AutherID != userID {
		return serrors.ErrPermissionNotFound
	}

	now := time.Now()
	chat.UpdatedAt = now
	chat.DeletedAt = &now

	if err := i._chat.Update(ctx, chat); err != nil {
		return err
	}

	return nil
}
