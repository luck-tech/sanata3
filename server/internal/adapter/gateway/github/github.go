package github

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	v67 "github.com/google/go-github/v67/github"
	"github.com/murasame29/go-httpserver-template/cmd/config"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
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

func (s *GitHubSerivce) GetUserByToken(ctx context.Context, token *oauth2.Token) (*v67.User, error) {
	return nil, nil
}

var _ dai.GitHubService = (*GitHubSerivce)(nil)
