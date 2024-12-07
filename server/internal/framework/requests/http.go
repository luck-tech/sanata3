package requests

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2"
)

func RequestWithAccessToken[T any](ctx context.Context, endpoint, accessToken string) (*T, error) {
	fmt.Println("token:", accessToken, "endpoint", endpoint)
	tc := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken, TokenType: "Bearer"},
	))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := tc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("resp body:", string(data))
	var body T
	if err := json.Unmarshal(data, &body); err != nil {
		return nil, err
	}

	return &body, nil
}

func Request[T any](ctx context.Context, endpoint string) (*T, error) {
	c := http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	fmt.Println("endpoint", endpoint)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var body T
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}
