package mastodon

import (
	"context"
)

// TagInfo gets statistics and information about a tag
func (c *Client) TagInfo(ctx context.Context, tag string) (*FollowedTag, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TagFollow lets you follow a hashtag
func (c *Client) TagFollow(ctx context.Context, tag string) (*FollowedTag, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TagUnfollow unfollows a hashtag.
func (c *Client) TagUnfollow(ctx context.Context, ID string) (*FollowedTag, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TagsFollowed returns a list of hashtags you follow.
func (c *Client) TagsFollowed(ctx context.Context, pg *Pagination) ([]*FollowedTag, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
