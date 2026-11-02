package pixela

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// CreateUserParams contains parameters for creating a new user.
type CreateUserParams struct {
	AgreeTermsOfService bool   `json:"-"`
	NotMinor            bool   `json:"-"`
	ThanksCode          string `json:"thanksCode,omitempty"`
}

type createUserPayload struct {
	Token               string `json:"token"`
	Username            string `json:"username"`
	AgreeTermsOfService string `json:"agreeTermsOfService"`
	NotMinor            string `json:"notMinor"`
	ThanksCode          string `json:"thanksCode,omitempty"`
}

// UpdateUserParams contains parameters for updating an existing user's token.
type UpdateUserParams struct {
	NewToken   string `json:"newToken"`
	ThanksCode string `json:"thanksCode,omitempty"`
}

// CreateUser creates a new Pixela user account.
func (c *Client) CreateUser(ctx context.Context, params CreateUserParams) (*Result, error) {
	agree := "no"
	if params.AgreeTermsOfService {
		agree = "yes"
	}
	notMinor := "no"
	if params.NotMinor {
		notMinor = "yes"
	}

	payload := createUserPayload{
		Token:               c.token,
		Username:            c.username,
		AgreeTermsOfService: agree,
		NotMinor:            notMinor,
		ThanksCode:          params.ThanksCode,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	reqURL, err := url.JoinPath(c.baseURL, "v1", "users")
	if err != nil {
		return nil, err
	}

	return c.doResult(ctx, "POST", reqURL, payloadBytes, false)
}

// UpdateUser updates the secret user token.
func (c *Client) UpdateUser(ctx context.Context, params UpdateUserParams) (*Result, error) {
	if err := validateToken(params.NewToken); err != nil {
		return nil, fmt.Errorf("invalid new token: %w", err)
	}

	payloadBytes, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	reqURL, err := url.JoinPath(c.baseURL, "v1", "users", c.username)
	if err != nil {
		return nil, err
	}

	return c.doResult(ctx, "PUT", reqURL, payloadBytes, true)
}

// DeleteUser deletes the user account.
func (c *Client) DeleteUser(ctx context.Context) (*Result, error) {
	reqURL, err := url.JoinPath(c.baseURL, "v1", "users", c.username)
	if err != nil {
		return nil, err
	}

	return c.doResult(ctx, "DELETE", reqURL, nil, true)
}
