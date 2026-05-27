package mastodon

import (
	"context"
)

// Instance holds information for a mastodon instance.
type Instance struct {
	URI            string            `json:"uri"`
	Title          string            `json:"title"`
	Description    string            `json:"description"`
	EMail          string            `json:"email"`
	Version        string            `json:"version,omitempty"`
	Thumbnail      string            `json:"thumbnail,omitempty"`
	URLs           map[string]string `json:"urls,omitempty"`
	Stats          *InstanceStats    `json:"stats,omitempty"`
	Languages      []string          `json:"languages"`
	ContactAccount *Account          `json:"contact_account"`
	Configuration  *InstanceConfig   `json:"configuration"`
}

type InstanceConfigMap map[string]interface{}

// InstanceConfig holds configuration accessible for clients.
type InstanceConfig struct {
	Accounts         *InstanceConfigMap     `json:"accounts"`
	Statuses         *InstanceConfigMap     `json:"statuses"`
	MediaAttachments map[string]interface{} `json:"media_attachments"`
	Polls            *InstanceConfigMap     `json:"polls"`
}

// InstanceStats holds information for mastodon instance stats.
type InstanceStats struct {
	UserCount   int64 `json:"user_count"`
	StatusCount int64 `json:"status_count"`
	DomainCount int64 `json:"domain_count"`
}

// GetInstance returns Instance.
func (c *Client) GetInstance(ctx context.Context) (*Instance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetConfig returns InstanceConfig.
func (c *Instance) GetConfig() *InstanceConfig { _ = "STUB: not implemented"; return nil }

// WeeklyActivity holds information for mastodon weekly activity.
type WeeklyActivity struct {
	Week          Unixtime `json:"week"`
	Statuses      int64    `json:"statuses,string"`
	Logins        int64    `json:"logins,string"`
	Registrations int64    `json:"registrations,string"`
}

// GetInstanceActivity returns instance activity.
func (c *Client) GetInstanceActivity(ctx context.Context) ([]*WeeklyActivity, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetInstancePeers returns instance peers.
func (c *Client) GetInstancePeers(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
