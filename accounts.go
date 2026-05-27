package mastodon

import (
	"context"
	"iter"
	"time"
)

// Account holds information for a mastodon account.
type Account struct {
	ID             ID             `json:"id"`
	Username       string         `json:"username"`
	Acct           string         `json:"acct"`
	DisplayName    string         `json:"display_name"`
	Locked         bool           `json:"locked"`
	CreatedAt      time.Time      `json:"created_at"`
	FollowersCount int64          `json:"followers_count"`
	FollowingCount int64          `json:"following_count"`
	StatusesCount  int64          `json:"statuses_count"`
	Note           string         `json:"note"`
	URL            string         `json:"url"`
	Avatar         string         `json:"avatar"`
	AvatarStatic   string         `json:"avatar_static"`
	Header         string         `json:"header"`
	HeaderStatic   string         `json:"header_static"`
	Emojis         []Emoji        `json:"emojis"`
	Moved          *Account       `json:"moved"`
	Fields         []Field        `json:"fields"`
	Bot            bool           `json:"bot"`
	Discoverable   bool           `json:"discoverable"`
	Source         *AccountSource `json:"source"`
	FollowedTag    []FollowedTag  `json:"followed_tags"`
}

// Field is a Mastodon account profile field.
type Field struct {
	Name       string    `json:"name"`
	Value      string    `json:"value"`
	VerifiedAt time.Time `json:"verified_at"`
}

// AccountSource is a Mastodon account profile field.
type AccountSource struct {
	Privacy   *string  `json:"privacy"`
	Sensitive *bool    `json:"sensitive"`
	Language  *string  `json:"language"`
	Note      *string  `json:"note"`
	Fields    *[]Field `json:"fields"`
}

// UnixTimeString represents a time in a Unix Epoch string
type UnixTimeString struct {
	time.Time
}

func (u *UnixTimeString) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// History is the history of a followed tag
type FollowedTagHistory struct {
	Day      UnixTimeString `json:"day,omitempty"`
	Accounts int            `json:"accounts,string,omitempty"`
	Uses     int            `json:"uses,string,omitempty"`
}

// FollowedTag is a Hash Tag followed by the user
type FollowedTag struct {
	Name      string               `json:"name,omitempty"`
	URL       string               `json:"url,omitempty"`
	History   []FollowedTagHistory `json:"history,omitempty"`
	Following bool                 `json:"following,omitempty"`
}

// GetAccount return Account.
func (c *Client) GetAccount(ctx context.Context, id ID) (*Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAccountCurrentUser returns the Account of current user.
func (c *Client) GetAccountCurrentUser(ctx context.Context) (*Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AccountLookup returns the Account of specified acct uri.
func (c *Client) AccountLookup(ctx context.Context, acct string) (*Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Profile is a struct for updating profiles.
type Profile struct {
	// If it is nil it will not be updated.
	// If it is empty, update it with empty.
	DisplayName *string
	Note        *string
	Locked      *bool
	Fields      *[]Field
	Source      *AccountSource

	// Set the base64 encoded character string of the image.
	Avatar string
	Header string
}

// AccountUpdate updates the information of the current user.
func (c *Client) AccountUpdate(ctx context.Context, profile *Profile) (*Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AccountStatuses iterates over statuses for the provided account id.
func (c *Client) AccountStatuses(ctx context.Context, id ID, pg *Pagination) iter.Seq2[*Status, error] {
	_ = "STUB: not implemented"
	return nil
}

// GetAccountStatuses return statuses by specified account.
func (c *Client) GetAccountStatuses(ctx context.Context, id ID, pg *Pagination) ([]*Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAccountPinnedStatuses returns statuses pinned by specified accuont.
func (c *Client) GetAccountPinnedStatuses(ctx context.Context, id ID) ([]*Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAccountFollowers returns followers list.
func (c *Client) GetAccountFollowers(ctx context.Context, id ID, pg *Pagination) ([]*Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAccountFollowing returns following list.
func (c *Client) GetAccountFollowing(ctx context.Context, id ID, pg *Pagination) ([]*Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBlocks returns block list.
func (c *Client) GetBlocks(ctx context.Context, pg *Pagination) ([]*Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetEndorsements return accounts that the user is currently featuring on their profile.
func (c *Client) GetEndorsements(ctx context.Context, pg *Pagination) ([]*Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Relationship holds information for relationship to the account.
type Relationship struct {
	ID                  ID   `json:"id"`
	Following           bool `json:"following"`
	FollowedBy          bool `json:"followed_by"`
	Blocking            bool `json:"blocking"`
	Muting              bool `json:"muting"`
	MutingNotifications bool `json:"muting_notifications"`
	Requested           bool `json:"requested"`
	DomainBlocking      bool `json:"domain_blocking"`
	ShowingReblogs      bool `json:"showing_reblogs"`
	Endorsed            bool `json:"endorsed"`
}

// AccountFollow follows the account.
func (c *Client) AccountFollow(ctx context.Context, id ID) (*Relationship, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AccountUnfollow unfollows the account.
func (c *Client) AccountUnfollow(ctx context.Context, id ID) (*Relationship, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AccountBlock blocks the account.
func (c *Client) AccountBlock(ctx context.Context, id ID) (*Relationship, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AccountUnblock unblocks the account.
func (c *Client) AccountUnblock(ctx context.Context, id ID) (*Relationship, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AccountMute mutes the account.
func (c *Client) AccountMute(ctx context.Context, id ID) (*Relationship, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AccountUnmute unmutes the account.
func (c *Client) AccountUnmute(ctx context.Context, id ID) (*Relationship, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAccountRelationships returns relationship for the account.
func (c *Client) GetAccountRelationships(ctx context.Context, ids []string) ([]*Relationship, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AccountsSearch searches accounts by query.
func (c *Client) AccountsSearch(ctx context.Context, q string, limit int64) ([]*Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) AccountsSearchResolve(ctx context.Context, q string, limit int64, resolve bool) ([]*Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FollowRemoteUser sends follow-request.
func (c *Client) FollowRemoteUser(ctx context.Context, uri string) (*Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetFollowRequests returns follow requests.
func (c *Client) GetFollowRequests(ctx context.Context, pg *Pagination) ([]*Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FollowRequestAuthorize authorizes the follow request of user with id.
func (c *Client) FollowRequestAuthorize(ctx context.Context, id ID) error {
	_ = "STUB: not implemented"
	return nil
}

// FollowRequestReject rejects the follow request of user with id.
func (c *Client) FollowRequestReject(ctx context.Context, id ID) error {
	_ = "STUB: not implemented"
	return nil
}

// GetMutes returns the list of users muted by the current user.
func (c *Client) GetMutes(ctx context.Context, pg *Pagination) ([]*Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetFollowedTags returns the list of Hashtags followed by the user.
func (c *Client) GetFollowedTags(ctx context.Context, pg *Pagination) ([]*FollowedTag, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
