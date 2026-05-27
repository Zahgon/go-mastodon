package mastodon

import (
	"context"
)

// List is metadata for a list of users.
type List struct {
	ID    ID     `json:"id"`
	Title string `json:"title"`
}

// GetLists returns all the lists on the current account.
func (c *Client) GetLists(ctx context.Context) ([]*List, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAccountLists returns the lists containing a given account.
func (c *Client) GetAccountLists(ctx context.Context, id ID) ([]*List, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetListAccounts returns the accounts in a given list.
func (c *Client) GetListAccounts(ctx context.Context, id ID) ([]*Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetList retrieves a list by ID.
func (c *Client) GetList(ctx context.Context, id ID) (*List, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateList creates a new list with a given title.
func (c *Client) CreateList(ctx context.Context, title string) (*List, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RenameList assigns a new title to a list.
func (c *Client) RenameList(ctx context.Context, id ID, title string) (*List, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteList removes a list.
func (c *Client) DeleteList(ctx context.Context, id ID) error {
	_ = "STUB: not implemented"
	return nil
}

// AddToList adds accounts to a list.
//
// Only accounts already followed by the user can be added to a list.
func (c *Client) AddToList(ctx context.Context, list ID, accounts ...ID) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveFromList removes accounts from a list.
func (c *Client) RemoveFromList(ctx context.Context, list ID, accounts ...ID) error {
	_ = "STUB: not implemented"
	return nil
}
