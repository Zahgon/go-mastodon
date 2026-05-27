package mastodon

import (
	"context"
	"time"
)

// Filter is metadata for a filter of users.
type Filter struct {
	ID           ID        `json:"id"`
	Phrase       string    `json:"phrase"`
	Context      []string  `json:"context"`
	WholeWord    bool      `json:"whole_word"`
	ExpiresAt    time.Time `json:"expires_at"`
	Irreversible bool      `json:"irreversible"`
}

type FilterResult struct {
	Filter struct {
		ID           string    `json:"id"`
		Title        string    `json:"title"`
		Context      []string  `json:"context"`
		ExpiresAt    time.Time `json:"expires_at"`
		FilterAction string    `json:"filter_action"`
	} `json:"filter"`
	KeywordMatches []string `json:"keyword_matches"`
	StatusMatches  []string `json:"status_matches"`
}

// GetFilters returns all the filters on the current account.
func (c *Client) GetFilters(ctx context.Context) ([]*Filter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetFilter retrieves a filter by ID.
func (c *Client) GetFilter(ctx context.Context, id ID) (*Filter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateFilter creates a new filter.
func (c *Client) CreateFilter(ctx context.Context, filter *Filter) (*Filter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateFilter updates a filter.
func (c *Client) UpdateFilter(ctx context.Context, id ID, filter *Filter) (*Filter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteFilter removes a filter.
func (c *Client) DeleteFilter(ctx context.Context, id ID) error {
	_ = "STUB: not implemented"
	return nil
}
