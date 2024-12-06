package dai

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
)

type User interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUser(ctx context.Context, id string) (*entity.User, bool, error)
	GetUsers(ctx context.Context, userIDs []string) ([]entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id string) error
}
