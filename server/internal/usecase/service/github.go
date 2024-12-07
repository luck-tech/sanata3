package service

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
)

type GitHub struct {
	repo dai.DataAccessInterface
}

func NewGitHub(repo dai.DataAccessInterface) *GitHub {
	return &GitHub{
		repo: repo,
	}
}

type LoginGitHubResult struct {
	AccessToken  string
	RefreshToken string
	UserID       string
	UserName     string
	Icon         string
	IsNewUser    bool
}

func (g *GitHub) Login(ctx context.Context, code string) (*LoginGitHubResult, error) {
	token, err := g.repo.FetchToken(ctx, code)
	if err != nil {
		return nil, err
	}
	isNewUser := false
	fmt.Println(token)

	userInfo, err := g.repo.GetUserByToken(ctx, token.AccessToken)
	if err != nil {
		return nil, err
	}
	user, found, err := g.repo.GetUser(ctx, strconv.Itoa(userInfo.ID))
	if err != nil {
		return nil, err
	}
	if !found {
		isNewUser = true
		newUser := &entity.User{
			ID:    strconv.Itoa(userInfo.ID),
			Email: userInfo.Email,
			Name:  userInfo.Login,
			Icon:  userInfo.AvatarURL,
		}
		if err := g.repo.CreateUser(ctx, newUser); err != nil {
			return nil, err
		}

		// TODO: Neptune
	
	} else {
		user.Name = userInfo.Login
		user.Icon = userInfo.AvatarURL
		if err := g.repo.UpdateUser(ctx, user); err != nil {
			return nil, err
		}
	}

	user, _, err = g.repo.GetUser(ctx, strconv.Itoa(userInfo.ID))
	if err != nil {
		return nil, err
	}

	return &LoginGitHubResult{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		UserID:       user.ID,
		UserName:     user.Name,
		Icon:         user.Icon,
		IsNewUser:    isNewUser,
	}, nil
}

func (g *GitHub) GetUsedLanguage(ctx context.Context, name, token string) (map[string]int, error) {
	languages, err := g.repo.GetUserUseLanguagesByID(ctx, token, name)
	if err != nil {
		return nil, err
	}

	type kv struct {
		Key   string
		Value int
	}
	var orderedLanguages []kv
	for k, v := range languages {
		orderedLanguages = append(orderedLanguages, kv{k, v})
	}

	sort.Slice(orderedLanguages, func(i, j int) bool {
		return orderedLanguages[i].Value > orderedLanguages[j].Value
	})

	var result = make(map[string]int)
	for i, kv := range orderedLanguages {
		if i >= 10 {
			break
		}
		result[kv.Key] = kv.Value
	}

	return result, nil
}
