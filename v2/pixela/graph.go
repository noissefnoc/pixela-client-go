package pixela

import "context"

// CreateGraphParams contains parameters for creating a new pixelation graph.
type CreateGraphParams struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Unit           string   `json:"unit"`
	Type           string   `json:"type"` // "int" or "float"
	Color          string   `json:"color"` // "shibafu", "momiji", "sora", "ichou", "ajisai", "kuro"
	Timezone       string   `json:"timezone,omitempty"`
	SelfSufficient string   `json:"selfSufficient,omitempty"` // "increment" or "decrement" or "none"
	IsSecret       bool     `json:"isSecret,omitempty"`
	PublishOptional bool    `json:"publishOptionalData,omitempty"`
	PurgeCacheURLs []string `json:"purgeCacheURLs,omitempty"`
}

// UpdateGraphParams contains parameters for updating an existing graph.
type UpdateGraphParams struct {
	Name           string   `json:"name,omitempty"`
	Unit           string   `json:"unit,omitempty"`
	Color          string   `json:"color,omitempty"`
	Timezone       string   `json:"timezone,omitempty"`
	SelfSufficient string   `json:"selfSufficient,omitempty"`
	IsSecret       *bool    `json:"isSecret,omitempty"`
	PublishOptional *bool   `json:"publishOptionalData,omitempty"`
	PurgeCacheURLs []string `json:"purgeCacheURLs,omitempty"`
}

// Graph represents a single graph definition.
type Graph struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Unit           string   `json:"unit"`
	Type           string   `json:"type"`
	Color          string   `json:"color"`
	Timezone       string   `json:"timezone"`
	SelfSufficient string   `json:"selfSufficient"`
	IsSecret       bool     `json:"isSecret"`
	PublishOptional bool    `json:"publishOptionalData"`
	PurgeCacheURLs []string `json:"purgeCacheURLs"`
}

// GraphDefinitions represents a list of graph definitions.
type GraphDefinitions struct {
	Graphs []Graph `json:"graphs"`
}

// GraphStats represents pixel statistics of a graph.
type GraphStats struct {
	TotalPixelsCount int     `json:"totalPixelsCount"`
	MaxQuantity      int     `json:"maxQuantity"`
	MinQuantity      int     `json:"minQuantity"`
	TotalQuantity    int     `json:"totalQuantity"`
	AvgQuantity      float64 `json:"avgQuantity"`
	TodaysQuantity   int     `json:"todaysQuantity"`
}

// PixelDates represents dates on which pixels were posted.
type PixelDates struct {
	Pixels []string `json:"pixels"`
}

// GetGraphOptions contains optional query parameters for GetGraph SVG.
type GetGraphOptions struct {
	Date     string
	Mode     string // "short" or "badge" or empty
	Appearance string // "dark"
}

// CreateGraph creates a new pixelation graph.
func (c *Client) CreateGraph(ctx context.Context, params CreateGraphParams) (*Result, error) {
	return nil, errNotImplemented
}

// GetGraph retrieves the graph SVG image as raw bytes.
func (c *Client) GetGraph(ctx context.Context, graphID string, opts ...GetGraphOptions) ([]byte, error) {
	return nil, errNotImplemented
}

// GetGraphDefinitions retrieves all graph definitions for the user.
func (c *Client) GetGraphDefinitions(ctx context.Context) (*GraphDefinitions, error) {
	return nil, errNotImplemented
}

// GetGraphDetail retrieves a specific graph definition.
func (c *Client) GetGraphDetail(ctx context.Context, graphID string) (*Graph, error) {
	return nil, errNotImplemented
}

// GetGraphStats retrieves statistics of a graph.
func (c *Client) GetGraphStats(ctx context.Context, graphID string) (*GraphStats, error) {
	return nil, errNotImplemented
}

// GetPixelDates retrieves a list of dates on which pixels are recorded.
func (c *Client) GetPixelDates(ctx context.Context, graphID, from, to string) (*PixelDates, error) {
	return nil, errNotImplemented
}

// UpdateGraph updates an existing graph configuration.
func (c *Client) UpdateGraph(ctx context.Context, graphID string, params UpdateGraphParams) (*Result, error) {
	return nil, errNotImplemented
}

// DeleteGraph deletes a graph.
func (c *Client) DeleteGraph(ctx context.Context, graphID string) (*Result, error) {
	return nil, errNotImplemented
}

// StopGraphStopwatch stops a running pixel stopwatch on a graph.
func (c *Client) StopGraphStopwatch(ctx context.Context, graphID string) (*Result, error) {
	return nil, errNotImplemented
}
