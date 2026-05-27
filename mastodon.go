// Package mastodon provides functions and structs for accessing the mastodon API.
package mastodon

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Config is a setting for access mastodon APIs.
type Config struct {
	Server       string
	ClientID     string
	ClientSecret string
	AccessToken  string
}

type WriterResetter interface {
	io.Writer
	Reset()
}

// Client is a API client for mastodon.
type Client struct {
	http.Client
	Config     *Config
	UserAgent  string
	JSONWriter io.Writer
}

func (c *Client) doAPI(ctx context.Context, method string, uri string, params interface{}, res interface{}, pg *Pagination) error {
	_ = "STUB: not implemented"
	return nil
}

// handle status code 429, which indicates the server is throttling
// our requests. Do an exponential backoff and retry the request.

// NewClient returns a new mastodon API client.
func NewClient(config *Config) *Client { _ = "STUB: not implemented"; return nil }

// Authenticate gets access-token to the API.
// DEPRECATED: Authenticating with username and password is no longer supported, please use
// GetAppAccessToken() or GetUserAccessToken() instead
func (c *Client) Authenticate(ctx context.Context, username, password string) error {
	_ = "STUB: not implemented"
	return nil
}

// AuthenticateApp logs in using client credentials.
// DEPRECATED: use GetAppAccessToken() instead
func (c *Client) AuthenticateApp(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// GetAppAccessToken exchanges API Credentials for an application Access Token
// https://docs.joinmastodon.org/api/oauth-tokens/#app-tokens
func (c *Client) GetAppAccessToken(ctx context.Context, redirectURI string) error {
	_ = "STUB: not implemented"
	return nil
}

// AuthenticateToken logs in using a grant token returned by Application.AuthURI.
// redirectURI should be the same as Application.RedirectURI.
// DEPRECATED:  Use GetUserAccessToken() instead
func (c *Client) AuthenticateToken(ctx context.Context, authCode, redirectURI string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetUserAccessToken exhanges a user provided authorization code for an User Access Token
// https://docs.joinmastodon.org/api/oauth-tokens/#user-tokens
func (c *Client) GetUserAccessToken(ctx context.Context, authCode, redirectURI string) error {
	_ = "STUB: not implemented"
	return nil
}

// DEPRECATED: Use getAccessToken() instead
func (c *Client) authenticate(ctx context.Context, params url.Values) error {
	_ = "STUB: not implemented"
	return nil
}

// Exchanges credentials for an access token to be used by applications and sets the access token in the client config
func (c *Client) getAccessToken(ctx context.Context, params url.Values) error {
	_ = "STUB: not implemented"
	return nil
}

// Convenience constants for Toot.Visibility
const (
	VisibilityPublic        = "public"
	VisibilityUnlisted      = "unlisted"
	VisibilityFollowersOnly = "private"
	VisibilityDirectMessage = "direct"
)

// Toot is a struct to post status.
type Toot struct {
	Status      string     `json:"status"`
	InReplyToID ID         `json:"in_reply_to_id"`
	MediaIDs    []ID       `json:"media_ids"`
	Sensitive   bool       `json:"sensitive"`
	SpoilerText string     `json:"spoiler_text"`
	Visibility  string     `json:"visibility"`
	Language    string     `json:"language"`
	ScheduledAt *time.Time `json:"scheduled_at,omitempty"`
	Poll        *TootPoll  `json:"poll"`
}

// TootPoll holds information for creating a poll in Toot.
type TootPoll struct {
	Options          []string `json:"options"`
	ExpiresInSeconds int64    `json:"expires_in"`
	Multiple         bool     `json:"multiple"`
	HideTotals       bool     `json:"hide_totals"`
}

// Mention hold information for mention.
type Mention struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Acct     string `json:"acct"`
	ID       ID     `json:"id"`
}

// Tag hold information for tag.
type Tag struct {
	Name    string    `json:"name"`
	URL     string    `json:"url"`
	History []History `json:"history"`
}

// History hold information for history.
type History struct {
	Day      string `json:"day"`
	Uses     string `json:"uses"`
	Accounts string `json:"accounts"`
}

// Attachment hold information for attachment.
type Attachment struct {
	ID          ID             `json:"id"`
	Type        string         `json:"type"`
	URL         string         `json:"url"`
	RemoteURL   string         `json:"remote_url"`
	PreviewURL  string         `json:"preview_url"`
	TextURL     string         `json:"text_url"`
	Description string         `json:"description"`
	BlurHash    string         `json:"blurhash"`
	Meta        AttachmentMeta `json:"meta"`
}

// AttachmentMeta holds information for attachment metadata.
type AttachmentMeta struct {
	Original AttachmentSize  `json:"original"`
	Small    AttachmentSize  `json:"small"`
	Focus    AttachmentFocus `json:"focus"`
}

// AttachmentSize holds information for attatchment size.
type AttachmentSize struct {
	Width  int64   `json:"width"`
	Height int64   `json:"height"`
	Size   string  `json:"size"`
	Aspect float64 `json:"aspect"`
}

// AttachmentSize holds information for attatchment size.
type AttachmentFocus struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Emoji hold information for CustomEmoji.
type Emoji struct {
	ShortCode       string `json:"shortcode"`
	StaticURL       string `json:"static_url"`
	URL             string `json:"url"`
	VisibleInPicker bool   `json:"visible_in_picker"`
}

// Results hold information for search result.
type Results struct {
	Accounts []*Account `json:"accounts"`
	Statuses []*Status  `json:"statuses"`
	Hashtags []*Tag     `json:"hashtags"`
}

// Pagination is a struct for specifying the get range.
type Pagination struct {
	MaxID   ID
	SinceID ID
	MinID   ID
	Limit   int64
}

func newPaginationPrevNext(rawlink string) (prev, next Pagination, err error) {
	_ = "STUB: not implemented"
	return *new(Pagination), *new(Pagination), nil
}

func paginationFromLink(link string) (p Pagination, err error) {
	_ = "STUB: not implemented"
	return *new(Pagination), nil
}

func getPaginationID(rawurl, key string) (ID, error) {
	_ = "STUB: not implemented"
	return *new(ID), nil
}

func (p *Pagination) fromValues(vs url.Values) error { _ = "STUB: not implemented"; return nil }

func (p *Pagination) toValues() url.Values { _ = "STUB: not implemented"; return *new(url.Values) }

func (p *Pagination) setValues(params url.Values) url.Values {
	_ = "STUB: not implemented"
	return *new(url.Values)
}
