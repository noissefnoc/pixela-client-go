package pixela

import "context"

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
	return nil, errNotImplemented
}

// GetPixel retrieves the pixel value for a specified date.
func (c *Client) GetPixel(ctx context.Context, graphID, date string) (*Pixel, error) {
	return nil, errNotImplemented
}

// UpdatePixel updates an existing pixel for the specified date.
func (c *Client) UpdatePixel(ctx context.Context, graphID, date string, params UpdatePixelParams) (*Result, error) {
	return nil, errNotImplemented
}

// IncrementPixel increments the today's (or timezone-appropriate) pixel quantity.
func (c *Client) IncrementPixel(ctx context.Context, graphID string) (*Result, error) {
	return nil, errNotImplemented
}

// DecrementPixel decrements the today's (or timezone-appropriate) pixel quantity.
func (c *Client) DecrementPixel(ctx context.Context, graphID string) (*Result, error) {
	return nil, errNotImplemented
}

// DeletePixel deletes a pixel on the specified date.
func (c *Client) DeletePixel(ctx context.Context, graphID, date string) (*Result, error) {
	return nil, errNotImplemented
}
