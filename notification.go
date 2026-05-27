package mastodon

import (
	"context"
	"crypto/ecdsa"
	"iter"
	"time"
)

// Notification holds information for a mastodon notification.
type Notification struct {
	ID        ID        `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	Account   Account   `json:"account"`
	Status    *Status   `json:"status"`
}

type PushSubscription struct {
	ID        ID          `json:"id"`
	Endpoint  string      `json:"endpoint"`
	ServerKey string      `json:"server_key"`
	Alerts    *PushAlerts `json:"alerts"`
}

type PushAlerts struct {
	Follow    *Sbool `json:"follow"`
	Favourite *Sbool `json:"favourite"`
	Reblog    *Sbool `json:"reblog"`
	Mention   *Sbool `json:"mention"`
}

// NotificationFilter customizes how a notification query is submitted to
// the remote Mastodon server.
// See:
//
//	https://docs.joinmastodon.org/methods/notifications/
type NotificationFilter struct {
	Includes []string // list of notifications types to include
	Excludes []string // list of notifications types to exclude
}

// Notifications iterate over notifications.
func (c *Client) Notifications(ctx context.Context, filter *NotificationFilter, pg *Pagination) iter.Seq2[*Notification, error] {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) notificationsFilter(ctx context.Context, qry *NotificationFilter, pg *Pagination) iter.Seq2[*Notification, error] {
	_ = "STUB: not implemented"
	return nil
}

// GetNotifications returns notifications.
func (c *Client) GetNotifications(ctx context.Context, pg *Pagination) ([]*Notification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetNotificationsExclude returns notifications with excluded notifications
func (c *Client) GetNotificationsExclude(ctx context.Context, exclude *[]string, pg *Pagination) ([]*Notification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) getNotificationsFilter(ctx context.Context, qry *NotificationFilter, pg *Pagination) ([]*Notification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetNotification returns notification.
func (c *Client) GetNotification(ctx context.Context, id ID) (*Notification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DismissNotification deletes a single notification.
func (c *Client) DismissNotification(ctx context.Context, id ID) error {
	_ = "STUB: not implemented"
	return nil
}

// ClearNotifications clears notifications.
func (c *Client) ClearNotifications(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// AddPushSubscription adds a new push subscription.
func (c *Client) AddPushSubscription(ctx context.Context, endpoint string, public ecdsa.PublicKey, shared []byte, alerts PushAlerts) (*PushSubscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdatePushSubscription updates which type of notifications are sent for the active push subscription.
func (c *Client) UpdatePushSubscription(ctx context.Context, alerts *PushAlerts) (*PushSubscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RemovePushSubscription deletes the active push subscription.
func (c *Client) RemovePushSubscription(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// GetPushSubscription retrieves information about the active push subscription.
func (c *Client) GetPushSubscription(ctx context.Context) (*PushSubscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
