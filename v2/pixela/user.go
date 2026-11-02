package pixela

import "context"

// CreateUserParams contains parameters for creating a new user.
type CreateUserParams struct {
	AgreeTermsOfService bool   `json:"agreeTermsOfService"`
	NotMinor            bool   `json:"notMinor"`
	ThanksCode          string `json:"thanksCode,omitempty"`
}

// UpdateUserParams contains parameters for updating an existing user's token.
type UpdateUserParams struct {
	NewToken   string `json:"newToken"`
	ThanksCode string `json:"thanksCode,omitempty"`
}

// CreateUser creates a new Pixela user account.
func (c *Client) CreateUser(ctx context.Context, params CreateUserParams) (*Result, error) {
	return nil, errNotImplemented
}

// UpdateUser updates the secret user token.
func (c *Client) UpdateUser(ctx context.Context, params UpdateUserParams) (*Result, error) {
	return nil, errNotImplemented
}

// DeleteUser deletes the user account.
func (c *Client) DeleteUser(ctx context.Context) (*Result, error) {
	return nil, errNotImplemented
}
