package mastodon

import (
	"context"
	"time"
)

// Poll holds information for mastodon polls.
type Poll struct {
	ID          ID           `json:"id"`
	ExpiresAt   time.Time    `json:"expires_at"`
	Expired     bool         `json:"expired"`
	Multiple    bool         `json:"multiple"`
	VotesCount  int64        `json:"votes_count"`
	VotersCount int64        `json:"voters_count"`
	Options     []PollOption `json:"options"`
	Voted       bool         `json:"voted"`
	OwnVotes    []int        `json:"own_votes"`
	Emojis      []Emoji      `json:"emojis"`
}

// Poll holds information for a mastodon poll option.
type PollOption struct {
	Title      string `json:"title"`
	VotesCount int64  `json:"votes_count"`
}

// GetPoll returns poll specified by id.
func (c *Client) GetPoll(ctx context.Context, id ID) (*Poll, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PollVote votes on a poll specified by id, choices is the Poll.Options index to vote on
func (c *Client) PollVote(ctx context.Context, id ID, choices ...int) (*Poll, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
