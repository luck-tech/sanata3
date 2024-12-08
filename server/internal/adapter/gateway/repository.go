package gateway

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/murasame29/go-httpserver-template/internal/adapter/gateway/aws"
	"github.com/murasame29/go-httpserver-template/internal/adapter/gateway/github"
	"github.com/murasame29/go-httpserver-template/internal/adapter/gateway/repository"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
	"github.com/uptrace/bun"
)

type Repository struct {
	*repository.UserRepository
	*repository.SessionRepository
	*repository.SkillRepository
	*repository.UsedSkillRepository
	*repository.WantLearnSkillRepository
	*repository.AimSkillRepository
	*repository.RoomRepository
	*repository.RoomMemberRepository
	*github.GitHubSerivce
	*aws.DynamoRepository
}

func NewRepository(db bun.IDB, dynamo *dynamodb.Client) *Repository {
	return &Repository{
		UserRepository:           repository.NewUserRepository(db),
		SessionRepository:        repository.NewSessionRepository(db),
		SkillRepository:          repository.NewSkillRepository(db),
		UsedSkillRepository:      repository.NewUsedSkillRepository(db),
		WantLearnSkillRepository: repository.NewWantLearnSkillRepository(db),
		AimSkillRepository:       repository.NewAimSkillRepository(db),
		RoomRepository:           repository.NewRoomRepository(db),
		RoomMemberRepository:     repository.NewRoomMemberRepository(db),
		GitHubSerivce:            github.NewGitHubSerivce(),
		DynamoRepository:         aws.NewDynamoRepository(dynamo),
	}
}

var _ dai.DataAccessInterface = (*Repository)(nil)
