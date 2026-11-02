package pixela

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// CreateGraphParams contains parameters for creating a new pixelation graph.
type CreateGraphParams struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Unit            string   `json:"unit"`
	Type            string   `json:"type"`  // "int" or "float"
	Color           string   `json:"color"` // "shibafu", "momiji", "sora", "ichou", "ajisai", "kuro"
	Timezone        string   `json:"timezone,omitempty"`
	SelfSufficient  string   `json:"selfSufficient,omitempty"` // "increment", "decrement", "none"
	IsSecret        bool     `json:"isSecret,omitempty"`
	PublishOptional bool     `json:"publishOptionalData,omitempty"`
	PurgeCacheURLs  []string `json:"purgeCacheURLs,omitempty"`
}

// UpdateGraphParams contains parameters for updating an existing graph.
type UpdateGraphParams struct {
	Name            string   `json:"name,omitempty"`
	Unit            string   `json:"unit,omitempty"`
	Color           string   `json:"color,omitempty"`
	Timezone        string   `json:"timezone,omitempty"`
	SelfSufficient  string   `json:"selfSufficient,omitempty"`
	IsSecret        *bool    `json:"isSecret,omitempty"`
	PublishOptional *bool    `json:"publishOptionalData,omitempty"`
	PurgeCacheURLs  []string `json:"purgeCacheURLs,omitempty"`
}

// Graph represents a single graph definition.
type Graph struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Unit            string   `json:"unit"`
	Type            string   `json:"type"`
	Color           string   `json:"color"`
	Timezone        string   `json:"timezone"`
	SelfSufficient  string   `json:"selfSufficient"`
	IsSecret        bool     `json:"isSecret"`
	PublishOptional bool     `json:"publishOptionalData"`
	PurgeCacheURLs  []string `json:"purgeCacheURLs"`
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
	Date       string
	Mode       string // "short" or "badge" or empty
	Appearance string // "dark"
}

// CreateGraph creates a new pixelation graph.
func (c *Client) CreateGraph(ctx context.Context, params CreateGraphParams) (*Result, error) {
	if err := validateGraphID(params.ID); err != nil {
		return nil, fmt.Errorf("invalid graph id: %w", err)
	}
	if err := validateUnitType(params.Type); err != nil {
		return nil, fmt.Errorf("invalid unit type: %w", err)
	}
	if err := validateColor(params.Color); err != nil {
		return nil, fmt.Errorf("invalid color: %w", err)
	}
	if err := validateSelfSufficient(params.SelfSufficient); err != nil {
		return nil, fmt.Errorf("invalid self sufficient: %w", err)
	}

	payloadBytes, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	reqURL, err := url.JoinPath(c.baseURL, "v1", "users", c.username, "graphs")
	if err != nil {
		return nil, err
	}

	return c.doResult(ctx, "POST", reqURL, payloadBytes, true)
}

// GetGraph retrieves the graph SVG image as raw bytes.
func (c *Client) GetGraph(ctx context.Context, graphID string, opts ...GetGraphOptions) ([]byte, error) {
	if err := validateGraphID(graphID); err != nil {
		return nil, fmt.Errorf("invalid graph id: %w", err)
	}

	reqURL, err := url.JoinPath(c.baseURL, "v1", "users", c.username, "graphs", graphID)
	if err != nil {
		return nil, err
	}

	if len(opts) > 0 {
		u, err := url.Parse(reqURL)
		if err != nil {
			return nil, err
		}
		q := u.Query()
		if opts[0].Date != "" {
			q.Set("date", opts[0].Date)
		}
		if opts[0].Mode != "" {
			q.Set("mode", opts[0].Mode)
		}
		if opts[0].Appearance != "" {
			q.Set("appearance", opts[0].Appearance)
		}
		u.RawQuery = q.Encode()
		reqURL = u.String()
	}

	return c.do(ctx, "GET", reqURL, nil, true)
}

// GetGraphDefinitions retrieves all graph definitions for the user.
func (c *Client) GetGraphDefinitions(ctx context.Context) (*GraphDefinitions, error) {
	reqURL, err := url.JoinPath(c.baseURL, "v1", "users", c.username, "graphs")
	if err != nil {
		return nil, err
	}

	body, err := c.do(ctx, "GET", reqURL, nil, true)
	if err != nil {
		return nil, err
	}

	var defs GraphDefinitions
	if err := json.Unmarshal(body, &defs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal graph definitions: %w", err)
	}
	return &defs, nil
}

// GetGraphDetail retrieves a specific graph definition.
func (c *Client) GetGraphDetail(ctx context.Context, graphID string) (*Graph, error) {
	if err := validateGraphID(graphID); err != nil {
		return nil, fmt.Errorf("invalid graph id: %w", err)
	}

	reqURL, err := url.JoinPath(c.baseURL, "v1", "users", c.username, "graphs", graphID, "graph-def")
	if err != nil {
		return nil, err
	}

	body, err := c.do(ctx, "GET", reqURL, nil, true)
	if err != nil {
		return nil, err
	}

	var g Graph
	if err := json.Unmarshal(body, &g); err != nil {
		return nil, fmt.Errorf("failed to unmarshal graph detail: %w", err)
	}
	return &g, nil
}

// GetGraphStats retrieves statistics of a graph.
func (c *Client) GetGraphStats(ctx context.Context, graphID string) (*GraphStats, error) {
	if err := validateGraphID(graphID); err != nil {
		return nil, fmt.Errorf("invalid graph id: %w", err)
	}

	reqURL, err := url.JoinPath(c.baseURL, "v1", "users", c.username, "graphs", graphID, "stats")
	if err != nil {
		return nil, err
	}

	body, err := c.do(ctx, "GET", reqURL, nil, true)
	if err != nil {
		return nil, err
	}

	var stats GraphStats
	if err := json.Unmarshal(body, &stats); err != nil {
		return nil, fmt.Errorf("failed to unmarshal graph stats: %w", err)
	}
	return &stats, nil
}

// GetPixelDates retrieves a list of dates on which pixels are recorded.
func (c *Client) GetPixelDates(ctx context.Context, graphID, from, to string) (*PixelDates, error) {
	if err := validateGraphID(graphID); err != nil {
		return nil, fmt.Errorf("invalid graph id: %w", err)
	}
	if from != "" {
		if err := validateDate(from); err != nil {
			return nil, fmt.Errorf("invalid from date: %w", err)
		}
	}
	if to != "" {
		if err := validateDate(to); err != nil {
			return nil, fmt.Errorf("invalid to date: %w", err)
		}
	}

	reqURL, err := url.JoinPath(c.baseURL, "v1", "users", c.username, "graphs", graphID, "pixels")
	if err != nil {
		return nil, err
	}

	if from != "" || to != "" {
		u, err := url.Parse(reqURL)
		if err != nil {
			return nil, err
		}
		q := u.Query()
		if from != "" {
			q.Set("from", from)
		}
		if to != "" {
			q.Set("to", to)
		}
		u.RawQuery = q.Encode()
		reqURL = u.String()
	}

	body, err := c.do(ctx, "GET", reqURL, nil, true)
	if err != nil {
		return nil, err
	}

	var dates PixelDates
	if err := json.Unmarshal(body, &dates); err != nil {
		return nil, fmt.Errorf("failed to unmarshal pixel dates: %w", err)
	}
	return &dates, nil
}

// UpdateGraph updates an existing graph configuration.
func (c *Client) UpdateGraph(ctx context.Context, graphID string, params UpdateGraphParams) (*Result, error) {
	if err := validateGraphID(graphID); err != nil {
		return nil, fmt.Errorf("invalid graph id: %w", err)
	}
	if params.Color != "" {
		if err := validateColor(params.Color); err != nil {
			return nil, fmt.Errorf("invalid color: %w", err)
		}
	}
	if params.SelfSufficient != "" {
		if err := validateSelfSufficient(params.SelfSufficient); err != nil {
			return nil, fmt.Errorf("invalid self sufficient: %w", err)
		}
	}

	payloadBytes, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	reqURL, err := url.JoinPath(c.baseURL, "v1", "users", c.username, "graphs", graphID)
	if err != nil {
		return nil, err
	}

	return c.doResult(ctx, "PUT", reqURL, payloadBytes, true)
}

// DeleteGraph deletes a graph.
func (c *Client) DeleteGraph(ctx context.Context, graphID string) (*Result, error) {
	if err := validateGraphID(graphID); err != nil {
		return nil, fmt.Errorf("invalid graph id: %w", err)
	}

	reqURL, err := url.JoinPath(c.baseURL, "v1", "users", c.username, "graphs", graphID)
	if err != nil {
		return nil, err
	}

	return c.doResult(ctx, "DELETE", reqURL, nil, true)
}

// StopGraphStopwatch stops a running pixel stopwatch on a graph.
func (c *Client) StopGraphStopwatch(ctx context.Context, graphID string) (*Result, error) {
	if err := validateGraphID(graphID); err != nil {
		return nil, fmt.Errorf("invalid graph id: %w", err)
	}

	reqURL, err := url.JoinPath(c.baseURL, "v1", "users", c.username, "graphs", graphID, "stopwatch")
	if err != nil {
		return nil, err
	}

	return c.doResult(ctx, "POST", reqURL, nil, true)
}
