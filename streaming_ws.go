package mastodon

import (
	"context"
	"net/url"

	"github.com/gorilla/websocket"
)

// WSClient is a WebSocket client.
type WSClient struct {
	websocket.Dialer
	client *Client
}

// NewWSClient return WebSocket client.
func (c *Client) NewWSClient() *WSClient { _ = "STUB: not implemented"; return nil }

// Stream is a struct of data that flows in streaming.
type Stream struct {
	Event   string      `json:"event"`
	Payload interface{} `json:"payload"`
}

// StreamingWSUser return channel to read events on home using WebSocket.
func (c *WSClient) StreamingWSUser(ctx context.Context) (chan Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StreamingWSPublic return channel to read events on public using WebSocket.
func (c *WSClient) StreamingWSPublic(ctx context.Context, isLocal bool) (chan Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StreamingWSHashtag return channel to read events on tagged timeline using WebSocket.
func (c *WSClient) StreamingWSHashtag(ctx context.Context, tag string, isLocal bool) (chan Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StreamingWSList return channel to read events on a list using WebSocket.
func (c *WSClient) StreamingWSList(ctx context.Context, id ID) (chan Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StreamingWSDirect return channel to read events on a direct messages using WebSocket.
func (c *WSClient) StreamingWSDirect(ctx context.Context) (chan Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *WSClient) streamingWS(ctx context.Context, stream, tag string) (chan Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *WSClient) handleWS(ctx context.Context, rawurl string, q chan Event) error {
	_ = "STUB: not implemented"
	return nil
}

// End.

// Close the WebSocket when the context is canceled.

// End.

// Reconnect.

func (c *WSClient) dialRedirect(rawurl string) (conn *websocket.Conn, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *WSClient) dial(rawurl string) (*websocket.Conn, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func changeWebSocketScheme(rawurl string) (*url.URL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
