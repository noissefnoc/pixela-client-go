package pixela

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// CreateWebhookParams contains parameters for creating a webhook.
type CreateWebhookParams struct {
	GraphID string `json:"graphID"`
	Type    string `json:"type"` // "increment" or "decrement"
}

// Webhook represents a single webhook definition.
type Webhook struct {
	WebhookHash string `json:"webhookHash"`
	GraphID     string `json:"graphID"`
	Type        string `json:"type"`
}

// WebhookDefinitions represents a list of webhook definitions.
type WebhookDefinitions struct {
	Webhooks []Webhook `json:"webhooks"`
}

// CreateWebhook creates a new webhook for invoking actions without authentication.
func (c *Client) CreateWebhook(ctx context.Context, params CreateWebhookParams) (*Result, error) {
	if err := validateGraphID(params.GraphID); err != nil {
		return nil, fmt.Errorf("invalid graph id: %w", err)
	}
	if err := validateWebhookType(params.Type); err != nil {
		return nil, fmt.Errorf("invalid webhook type: %w", err)
	}

	payloadBytes, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	reqURL, err := url.JoinPath(c.baseURL, "v1", "users", c.username, "webhooks")
	if err != nil {
		return nil, err
	}

	return c.doResult(ctx, "POST", reqURL, payloadBytes, true)
}

// GetWebhooks retrieves all webhook definitions for the user.
func (c *Client) GetWebhooks(ctx context.Context) (*WebhookDefinitions, error) {
	reqURL, err := url.JoinPath(c.baseURL, "v1", "users", c.username, "webhooks")
	if err != nil {
		return nil, err
	}

	body, err := c.do(ctx, "GET", reqURL, nil, true)
	if err != nil {
		return nil, err
	}

	var defs WebhookDefinitions
	if err := json.Unmarshal(body, &defs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal webhook definitions: %w", err)
	}
	return &defs, nil
}

// InvokeWebhook triggers a webhook action using its webhook hash without user authentication token.
func (c *Client) InvokeWebhook(ctx context.Context, webhookHash string) (*Result, error) {
	reqURL, err := url.JoinPath(c.baseURL, "v1", "users", c.username, "webhooks", webhookHash)
	if err != nil {
		return nil, err
	}

	return c.doResult(ctx, "POST", reqURL, nil, false)
}

// DeleteWebhook deletes a webhook by its webhook hash.
func (c *Client) DeleteWebhook(ctx context.Context, webhookHash string) (*Result, error) {
	reqURL, err := url.JoinPath(c.baseURL, "v1", "users", c.username, "webhooks", webhookHash)
	if err != nil {
		return nil, err
	}

	return c.doResult(ctx, "DELETE", reqURL, nil, true)
}
