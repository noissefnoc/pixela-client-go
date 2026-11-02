package pixela

import (
	"fmt"
	"net/http"
	"time"
)

const (
	defaultBaseURL = "https://pixe.la"
	defaultTimeout = 10 * time.Second
)

// Client is a client for the Pixela API.
type Client struct {
	httpClient *http.Client
	baseURL    string
	username   string
	token      string
	debug      bool
}

// Option configures a Client.
type Option func(*Client)

// WithBaseURL sets a custom base URL for the API.
func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithHTTPClient sets a custom http.Client.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithDebug enables or disables debug logging.
func WithDebug(debug bool) Option {
	return func(c *Client) {
		c.debug = debug
	}
}

// Result represents the standard response from Pixela mutating APIs (POST, PUT, DELETE).
type Result struct {
	Message     string `json:"message"`
	IsSuccess   bool   `json:"isSuccess"`
	WebhookHash string `json:"webhookHash,omitempty"`
}

// APIError represents an error returned by the Pixela API.
type APIError struct {
	StatusCode int
	Message    string
	IsSuccess  bool
}

func (e *APIError) Error() string {
	return fmt.Sprintf("pixela API error (status %d): %s", e.StatusCode, e.Message)
}

// New creates a new Pixela API client.
func New(username, token string, opts ...Option) (*Client, error) {
	c := &Client{
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
		baseURL:  defaultBaseURL,
		username: username,
		token:    token,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c, nil
}

// Username returns the configured username.
func (c *Client) Username() string {
	return c.username
}

// BaseURL returns the configured base URL.
func (c *Client) BaseURL() string {
	return c.baseURL
}

// Mock/stub helper for Step 1 interface definition (implemented in Step 2)
var errNotImplemented = fmt.Errorf("not implemented yet (Step 2)")
