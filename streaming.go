package mastodon

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

// UpdateEvent is a struct for passing status event to app.
type UpdateEvent struct {
	Status *Status `json:"status"`
}

func (e *UpdateEvent) event() {
	_ = "STUB: not implemented"

	// UpdateEditEvent is a struct for passing status edit event to app.
	return
}

type UpdateEditEvent struct {
	Status *Status `json:"status"`
}

func (e *UpdateEditEvent) event() {
	_ = "STUB: not implemented"

	// NotificationEvent is a struct for passing notification event to app.
	return
}

type NotificationEvent struct {
	Notification *Notification `json:"notification"`
}

func (e *NotificationEvent) event() {
	_ = "STUB: not implemented"

	// DeleteEvent is a struct for passing deletion event to app.
	return
}

type DeleteEvent struct{ ID ID }

func (e *DeleteEvent) event() {
	_ = "STUB: not implemented"

	// ConversationEvent is a struct for passing conversationevent to app.
	return
}

type ConversationEvent struct {
	Conversation *Conversation `json:"conversation"`
}

func (e *ConversationEvent) event() {
	_ = "STUB: not implemented"

	// ErrorEvent is a struct for passing errors to app.
	return
}

type ErrorEvent struct{ Err error }

func (e *ErrorEvent) event()        { _ = "STUB: not implemented"; return }
func (e *ErrorEvent) Error() string { _ = "STUB: not implemented"; return "" }

// Event is an interface passing events to app.
type Event interface {
	event()
}

func handleReader(q chan Event, r io.Reader) error { _ = "STUB: not implemented"; return nil }

func (c *Client) streaming(ctx context.Context, p string, params url.Values) (chan Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) doStreaming(req *http.Request, q chan Event) { _ = "STUB: not implemented"; return }

// StreamingUser returns a channel to read events on home.
func (c *Client) StreamingUser(ctx context.Context) (chan Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StreamingPublic returns a channel to read events on public.
func (c *Client) StreamingPublic(ctx context.Context, isLocal bool) (chan Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StreamingHashtag returns a channel to read events on tagged timeline.
func (c *Client) StreamingHashtag(ctx context.Context, tag string, isLocal bool) (chan Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StreamingList returns a channel to read events on a list.
func (c *Client) StreamingList(ctx context.Context, id ID) (chan Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StreamingDirect returns a channel to read events on a direct messages.
func (c *Client) StreamingDirect(ctx context.Context) (chan Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
