package main

import (
	"github.com/mattn/go-mastodon"
	"github.com/urfave/cli/v2"
)

// SimpleJSON is a struct for output JSON for data to be simple used
type SimpleJSON struct {
	ID       mastodon.ID `json:"id"`
	Username string      `json:"username"`
	Acct     string      `json:"acct"`
	Avatar   string      `json:"avatar"`
	Content  string      `json:"content"`
}

func checkFlag(f ...bool) bool { _ = "STUB: not implemented"; return false }

func cmdStream(c *cli.Context) error { _ = "STUB: not implemented"; return nil }

// TODO s.displayStatus(c.App.Writer, t.Notification.Status)
