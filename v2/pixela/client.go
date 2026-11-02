package pixela

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
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
	if err := validateUsername(username); err != nil {
		return nil, fmt.Errorf("initialization error: %w", err)
	}
	if err := validateToken(token); err != nil {
		return nil, fmt.Errorf("initialization error: %w", err)
	}

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

func (c *Client) do(ctx context.Context, method, requestURL string, payload []byte, requiresAuth bool) ([]byte, error) {
	var bodyReader io.Reader
	if payload != nil {
		bodyReader = bytes.NewReader(payload)
	}

	req, err := http.NewRequestWithContext(ctx, method, requestURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	} else if method == http.MethodPost || method == http.MethodPut {
		req.Header.Set("Content-Length", "0")
	}

	if requiresAuth && c.token != "" {
		req.Header.Set("X-USER-TOKEN", c.token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		var apiErr Result
		if jsonErr := json.Unmarshal(respBody, &apiErr); jsonErr == nil && apiErr.Message != "" {
			return nil, &APIError{
				StatusCode: resp.StatusCode,
				Message:    apiErr.Message,
				IsSuccess:  apiErr.IsSuccess,
			}
		}
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    string(respBody),
			IsSuccess:  false,
		}
	}

	return respBody, nil
}

func (c *Client) doResult(ctx context.Context, method, requestURL string, payload []byte, requiresAuth bool) (*Result, error) {
	respBody, err := c.do(ctx, method, requestURL, payload, requiresAuth)
	if err != nil {
		return nil, err
	}

	var res Result
	if err := json.Unmarshal(respBody, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &res, nil
}
