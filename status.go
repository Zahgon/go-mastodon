package mastodon

import (
	"context"
	"io"
	"time"
)

// Status is struct to hold status.
type Status struct {
	ID                 ID              `json:"id"`
	URI                string          `json:"uri"`
	URL                string          `json:"url"`
	Account            Account         `json:"account"`
	InReplyToID        interface{}     `json:"in_reply_to_id"`
	InReplyToAccountID interface{}     `json:"in_reply_to_account_id"`
	Reblog             *Status         `json:"reblog"`
	Content            string          `json:"content"`
	CreatedAt          time.Time       `json:"created_at"`
	EditedAt           time.Time       `json:"edited_at"`
	Emojis             []Emoji         `json:"emojis"`
	RepliesCount       int64           `json:"replies_count"`
	ReblogsCount       int64           `json:"reblogs_count"`
	FavouritesCount    int64           `json:"favourites_count"`
	Reblogged          interface{}     `json:"reblogged"`
	Favourited         interface{}     `json:"favourited"`
	Bookmarked         interface{}     `json:"bookmarked"`
	Muted              interface{}     `json:"muted"`
	Sensitive          bool            `json:"sensitive"`
	SpoilerText        string          `json:"spoiler_text"`
	Visibility         string          `json:"visibility"`
	MediaAttachments   []Attachment    `json:"media_attachments"`
	Mentions           []Mention       `json:"mentions"`
	Tags               []Tag           `json:"tags"`
	Card               *Card           `json:"card"`
	Poll               *Poll           `json:"poll"`
	Application        Application     `json:"application"`
	Language           string          `json:"language"`
	Pinned             interface{}     `json:"pinned"`
	ScheduledParams    ScheduledParams `json:"params"`
	Filtered           []FilterResult  `json:"filtered"`
}

// StatusHistory is a struct to hold status history data.
type StatusHistory struct {
	Content          string       `json:"content"`
	SpoilerText      string       `json:"spoiler_text"`
	Account          Account      `json:"account"`
	Sensitive        bool         `json:"sensitive"`
	CreatedAt        time.Time    `json:"created_at"`
	Emojis           []Emoji      `json:"emojis"`
	MediaAttachments []Attachment `json:"media_attachments"`
}

// ScheduledStatus holds information returned when ScheduledAt is set on a status
type ScheduledParams struct {
	ApplicationID ID          `json:"application_id"`
	Idempotency   string      `json:"idempotency"`
	InReplyToID   interface{} `json:"in_reply_to_id"`
	MediaIDs      []ID        `json:"media_ids"`
	Poll          *Poll       `json:"poll"`
	ScheduledAt   *time.Time  `json:"scheduled_at,omitempty"`
	Sensitive     bool        `json:"sensitive"`
	SpoilerText   string      `json:"spoiler_text"`
	Text          string      `json:"text"`
	Visibility    string      `json:"visibility"`
}

// Context holds information for a mastodon context.
type Context struct {
	Ancestors   []*Status `json:"ancestors"`
	Descendants []*Status `json:"descendants"`
}

// Card holds information for a mastodon card.
type Card struct {
	URL          string `json:"url"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Image        string `json:"image"`
	Type         string `json:"type"`
	AuthorName   string `json:"author_name"`
	AuthorURL    string `json:"author_url"`
	ProviderName string `json:"provider_name"`
	ProviderURL  string `json:"provider_url"`
	HTML         string `json:"html"`
	Width        int64  `json:"width"`
	Height       int64  `json:"height"`
}

// Source holds source properties so a status can be edited.
type Source struct {
	ID          ID     `json:"id"`
	Text        string `json:"text"`
	SpoilerText string `json:"spoiler_text"`
}

// Conversation holds information for a mastodon conversation.
type Conversation struct {
	ID         ID         `json:"id"`
	Accounts   []*Account `json:"accounts"`
	Unread     bool       `json:"unread"`
	LastStatus *Status    `json:"last_status"`
}

// Media is struct to hold media.
type Media struct {
	File        io.Reader
	Thumbnail   io.Reader
	Description string
	Focus       string
}

type TagData struct {
	Any  []string
	All  []string
	None []string
}

func (m *Media) bodyAndContentType() (io.Reader, string, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), "", nil
}

// GetFavourites returns the favorite list of the current user.
func (c *Client) GetFavourites(ctx context.Context, pg *Pagination) ([]*Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBookmarks returns the bookmark list of the current user.
func (c *Client) GetBookmarks(ctx context.Context, pg *Pagination) ([]*Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetStatus returns status specified by id.
func (c *Client) GetStatus(ctx context.Context, id ID) (*Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetStatusContext returns status specified by id.
func (c *Client) GetStatusContext(ctx context.Context, id ID) (*Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetStatusCard returns status specified by id.
func (c *Client) GetStatusCard(ctx context.Context, id ID) (*Card, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetStatusSource returns source data specified by id.
func (c *Client) GetStatusSource(ctx context.Context, id ID) (*Source, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetStatusHistory returns the status history specified by id.
func (c *Client) GetStatusHistory(ctx context.Context, id ID) ([]*StatusHistory, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetRebloggedBy returns the account list of the user who reblogged the toot of id.
func (c *Client) GetRebloggedBy(ctx context.Context, id ID, pg *Pagination) ([]*Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetFavouritedBy returns the account list of the user who liked the toot of id.
func (c *Client) GetFavouritedBy(ctx context.Context, id ID, pg *Pagination) ([]*Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reblog reblogs the toot of id and returns status of reblog.
func (c *Client) Reblog(ctx context.Context, id ID) (*Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unreblog unreblogs the toot of id and returns status of the original toot.
func (c *Client) Unreblog(ctx context.Context, id ID) (*Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Favourite favourites the toot of id and returns status of the favourite toot.
func (c *Client) Favourite(ctx context.Context, id ID) (*Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unfavourite unfavourites the toot of id and returns status of the unfavourite toot.
func (c *Client) Unfavourite(ctx context.Context, id ID) (*Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Bookmark bookmarks the toot of id and returns status of the bookmark toot.
func (c *Client) Bookmark(ctx context.Context, id ID) (*Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unbookmark is unbookmark the toot of id and return status of the unbookmark toot.
func (c *Client) Unbookmark(ctx context.Context, id ID) (*Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetTimelineHome return statuses from home timeline.
func (c *Client) GetTimelineHome(ctx context.Context, pg *Pagination) ([]*Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetTrendingStatuses return statuses from explore timeline.
func (c *Client) GetTrendingStatuses(ctx context.Context, pg *Pagination) ([]*Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetTimelinePublic return statuses from public timeline.
func (c *Client) GetTimelinePublic(ctx context.Context, isLocal bool, pg *Pagination) ([]*Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetTimelineHashtag return statuses from tagged timeline.
func (c *Client) GetTimelineHashtag(ctx context.Context, tag string, isLocal bool, pg *Pagination) ([]*Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetTimelineHashtagMultiple return statuses from tagged timeline.
func (c *Client) GetTimelineHashtagMultiple(ctx context.Context, tag string, isLocal bool, td *TagData, pg *Pagination) ([]*Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetTimelineList return statuses from a list timeline.
func (c *Client) GetTimelineList(ctx context.Context, id ID, pg *Pagination) ([]*Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetTimelineMedia return statuses from media timeline.
// NOTE: This is an experimental feature of pawoo.net.
func (c *Client) GetTimelineMedia(ctx context.Context, isLocal bool, pg *Pagination) ([]*Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PostStatus post the toot.
func (c *Client) PostStatus(ctx context.Context, toot *Toot) (*Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateStatus updates the toot.
func (c *Client) UpdateStatus(ctx context.Context, toot *Toot, id ID) (*Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) postStatus(ctx context.Context, toot *Toot, update bool, updateID ID) (*Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Can't use Media and Poll at the same time.

// DeleteStatus delete the toot.
func (c *Client) DeleteStatus(ctx context.Context, id ID) error {
	_ = "STUB: not implemented"
	return nil
}

// Search search content with query.
func (c *Client) Search(ctx context.Context, q string, resolve bool) (*Results, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UploadMedia upload a media attachment from a file.
func (c *Client) UploadMedia(ctx context.Context, file string) (*Attachment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UploadMediaFromBytes uploads a media attachment from a byte slice.
func (c *Client) UploadMediaFromBytes(ctx context.Context, b []byte) (*Attachment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UploadMediaFromReader uploads a media attachment from an io.Reader.
func (c *Client) UploadMediaFromReader(ctx context.Context, reader io.Reader) (*Attachment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UploadMediaFromMedia uploads a media attachment from a Media struct.
func (c *Client) UploadMediaFromMedia(ctx context.Context, media *Media) (*Attachment, error) {
	_ = "STUB: not implemented"
	return nil,

		// "/api/v2/media" USING V2 TODO: IMPORTANT
		nil
}

// GetMediaStatus checks the status of a media attachment.
func (c *Client) GetMediaStatus(ctx context.Context, attachment *Attachment) error {
	_ = "STUB: not implemented"
	return nil
}

// GetTimelineDirect return statuses from direct timeline.
func (c *Client) GetTimelineDirect(ctx context.Context, pg *Pagination) ([]*Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetConversations return direct conversations.
func (c *Client) GetConversations(ctx context.Context, pg *Pagination) ([]*Conversation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteConversation delete the conversation specified by id.
func (c *Client) DeleteConversation(ctx context.Context, id ID) error {
	_ = "STUB: not implemented"
	return nil
}

// MarkConversationAsRead mark the conversation as read.
func (c *Client) MarkConversationAsRead(ctx context.Context, id ID) error {
	_ = "STUB: not implemented"
	return nil
}
