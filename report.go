package mastodon

import (
	"context"
)

// Report holds information for a mastodon report.
type Report struct {
	ID          ID   `json:"id"`
	ActionTaken bool `json:"action_taken"`
}

// GetReports returns report of the current user.
func (c *Client) GetReports(ctx context.Context) ([]*Report, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Report reports the report
func (c *Client) Report(ctx context.Context, accountID ID, ids []ID, comment string) (*Report, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
