package pixela

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// CreatePixelParams contains parameters for posting a pixel.
type CreatePixelParams struct {
	Date         string `json:"date"`                   // "yyyyMMdd"
	Quantity     string `json:"quantity"`               // int or float string
	OptionalData string `json:"optionalData,omitempty"` // JSON string (< 10KB)
}

// UpdatePixelParams contains parameters for updating an existing pixel.
type UpdatePixelParams struct {
	Quantity     string `json:"quantity"`
	OptionalData string `json:"optionalData,omitempty"`
}

// Pixel represents pixel details.
type Pixel struct {
	Quantity     string `json:"quantity"`
	OptionalData string `json:"optionalData,omitempty"`
}

// CreatePixel records a quantity to the graph for the specified date.
func (c *Client) CreatePixel(ctx context.Context, graphID string, params CreatePixelParams) (*Result, error) {
	if err := validateGraphID(graphID); err != nil {
		return nil, fmt.Errorf("invalid graph id: %w", err)
	}
	if err := validateDate(params.Date); err != nil {
		return nil, fmt.Errorf("invalid date: %w", err)
	}
	if err := validateQuantity(params.Quantity); err != nil {
		return nil, fmt.Errorf("invalid quantity: %w", err)
	}
	if err := validateOptionalData(params.OptionalData); err != nil {
		return nil, fmt.Errorf("invalid optional data: %w", err)
	}

	payloadBytes, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	reqURL, err := url.JoinPath(c.baseURL, "v1", "users", c.username, "graphs", graphID)
	if err != nil {
		return nil, err
	}

	return c.doResult(ctx, "POST", reqURL, payloadBytes, true)
}

// GetPixel retrieves the pixel value for a specified date.
func (c *Client) GetPixel(ctx context.Context, graphID, date string) (*Pixel, error) {
	if err := validateGraphID(graphID); err != nil {
		return nil, fmt.Errorf("invalid graph id: %w", err)
	}
	if err := validateDate(date); err != nil {
		return nil, fmt.Errorf("invalid date: %w", err)
	}

	reqURL, err := url.JoinPath(c.baseURL, "v1", "users", c.username, "graphs", graphID, date)
	if err != nil {
		return nil, err
	}

	body, err := c.do(ctx, "GET", reqURL, nil, true)
	if err != nil {
		return nil, err
	}

	var p Pixel
	if err := json.Unmarshal(body, &p); err != nil {
		return nil, fmt.Errorf("failed to unmarshal pixel: %w", err)
	}
	return &p, nil
}

// UpdatePixel updates an existing pixel for the specified date.
func (c *Client) UpdatePixel(ctx context.Context, graphID, date string, params UpdatePixelParams) (*Result, error) {
	if err := validateGraphID(graphID); err != nil {
		return nil, fmt.Errorf("invalid graph id: %w", err)
	}
	if err := validateDate(date); err != nil {
		return nil, fmt.Errorf("invalid date: %w", err)
	}
	if err := validateQuantity(params.Quantity); err != nil {
		return nil, fmt.Errorf("invalid quantity: %w", err)
	}
	if err := validateOptionalData(params.OptionalData); err != nil {
		return nil, fmt.Errorf("invalid optional data: %w", err)
	}

	payloadBytes, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	reqURL, err := url.JoinPath(c.baseURL, "v1", "users", c.username, "graphs", graphID, date)
	if err != nil {
		return nil, err
	}

	return c.doResult(ctx, "PUT", reqURL, payloadBytes, true)
}

// IncrementPixel increments the today's (or timezone-appropriate) pixel quantity.
func (c *Client) IncrementPixel(ctx context.Context, graphID string) (*Result, error) {
	if err := validateGraphID(graphID); err != nil {
		return nil, fmt.Errorf("invalid graph id: %w", err)
	}

	reqURL, err := url.JoinPath(c.baseURL, "v1", "users", c.username, "graphs", graphID, "increment")
	if err != nil {
		return nil, err
	}

	return c.doResult(ctx, "PUT", reqURL, nil, true)
}

// DecrementPixel decrements the today's (or timezone-appropriate) pixel quantity.
func (c *Client) DecrementPixel(ctx context.Context, graphID string) (*Result, error) {
	if err := validateGraphID(graphID); err != nil {
		return nil, fmt.Errorf("invalid graph id: %w", err)
	}

	reqURL, err := url.JoinPath(c.baseURL, "v1", "users", c.username, "graphs", graphID, "decrement")
	if err != nil {
		return nil, err
	}

	return c.doResult(ctx, "PUT", reqURL, nil, true)
}

// DeletePixel deletes a pixel on the specified date.
func (c *Client) DeletePixel(ctx context.Context, graphID, date string) (*Result, error) {
	if err := validateGraphID(graphID); err != nil {
		return nil, fmt.Errorf("invalid graph id: %w", err)
	}
	if err := validateDate(date); err != nil {
		return nil, fmt.Errorf("invalid date: %w", err)
	}

	reqURL, err := url.JoinPath(c.baseURL, "v1", "users", c.username, "graphs", graphID, date)
	if err != nil {
		return nil, err
	}

	return c.doResult(ctx, "DELETE", reqURL, nil, true)
}
