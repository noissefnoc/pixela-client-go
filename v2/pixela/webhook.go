package pixela

import "context"

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
	return nil, errNotImplemented
}

// GetWebhooks retrieves all webhook definitions for the user.
func (c *Client) GetWebhooks(ctx context.Context) (*WebhookDefinitions, error) {
	return nil, errNotImplemented
}

// InvokeWebhook triggers a webhook action using its webhook hash.
func (c *Client) InvokeWebhook(ctx context.Context, webhookHash string) (*Result, error) {
	return nil, errNotImplemented
}

// DeleteWebhook deletes a webhook by its webhook hash.
func (c *Client) DeleteWebhook(ctx context.Context, webhookHash string) (*Result, error) {
	return nil, errNotImplemented
}
