package github

import (
	"context"
	"fmt"
	"github.com/google/go-github/v74/github"
	"golang.org/x/oauth2"
)

type Client struct {
	client *github.Client
	token  string
}

// NewClient creates a new GitHub client
func NewClient(token string) *Client {
	if token == "" {
		return nil
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return &Client{
		client: client,
		token:  token,
	}
}

// ValidateToken validates the GitHub token by making a test API call
func (c *Client) ValidateToken(ctx context.Context) error {
	user, _, err := c.client.Users.Get(ctx, "")
	if err != nil {
		return fmt.Errorf("failed to validate GitHub token: %w", err)
	}

	if user.Login == nil {
		return fmt.Errorf("invalid GitHub token: no user information returned")
	}

	return nil
}

func (c *Client) GetAuthenticatedUser(ctx context.Context) (*github.User, error) {
	user, _, err := c.client.Users.Get(ctx, "")
	if err != nil {
		return nil, fmt.Errorf("failed to get authenticated user: %w", err)
	}
	return user, nil
}
