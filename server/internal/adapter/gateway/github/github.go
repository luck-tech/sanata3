package github

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/murasame29/go-httpserver-template/cmd/config"
	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/framework/requests"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
	"github.com/sourcegraph/conc"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type GitHubSerivce struct {
	oac oauth2.Config
}

func NewGitHubSerivce() *GitHubSerivce {
	return &GitHubSerivce{
		oac: oauth2.Config{
			ClientID:     config.Config.GitHub.ClientID,
			ClientSecret: config.Config.GitHub.ClientSecret,
			Scopes:       []string{},
			Endpoint: oauth2.Endpoint{
				AuthURL:  github.Endpoint.AuthURL,
				TokenURL: github.Endpoint.TokenURL,
			},
			RedirectURL: config.Config.GitHub.RedirectURI,
		},
	}
}

func (s *GitHubSerivce) FetchToken(ctx context.Context, code string) (*oauth2.Token, error) {
	// Set us the request body as JSON
	requestBodyMap := map[string]string{
		"client_id":     s.oac.ClientID,
		"client_secret": s.oac.ClientSecret,
		"code":          code,
	}
	requestJSON, err := json.Marshal(requestBodyMap)
	if err != nil {
		return nil, err
	}

	// POST request to set URL
	req, err := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		bytes.NewBuffer(requestJSON),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Get the response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Response body converted to stringified JSON
	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Convert stringified JSON to a struct object of type githubAccessTokenResponse
	var ghresp *oauth2.Token
	json.Unmarshal(respbody, &ghresp)

	return ghresp, nil
}

func (s *GitHubSerivce) GetUserByToken(ctx context.Context, accessToken string) (*entity.GitHubUser, error) {
	user, err := requests.RequestWithAccessToken[entity.GitHubUser](ctx, "https://api.github.com/user", accessToken)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *GitHubSerivce) GetUserUseLanguagesByID(ctx context.Context, accessToken, username string) (map[string]int, error) {
	repos, err := requests.RequestWithAccessToken[[]entity.GitHubRepo](ctx, fmt.Sprintf("https://api.github.com/users/%s/repos", username), accessToken)
	if err != nil {
		return nil, err
	}

	var wg conc.WaitGroup
	var languageMaps *entity.Language

	for _, repo := range *repos {
		wg.Go(func() {
			language, err := requests.RequestWithAccessToken[map[string]int](ctx, fmt.Sprintf("https://api.github.com/repos/%s/%s/languages", username, repo.Name), accessToken)
			if err != nil {
				return
			}

			for k, v := range *language {
				languageMaps.Store(k, v)
			}
		})
	}
	wg.Wait()

	return languageMaps.Count, nil
}

var _ dai.GitHubService = (*GitHubSerivce)(nil)
