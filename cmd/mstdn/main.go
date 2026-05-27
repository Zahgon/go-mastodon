package main

import (
	"bufio"
	"io"
	"os"

	"github.com/mattn/go-mastodon"
	"github.com/urfave/cli/v2"
)

func readFile(filename string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func textContent(s string) string { _ = "STUB: not implemented"; return "" }

var (
	readUsername = func() (string, error) {
		b, _, err := bufio.NewReader(os.Stdin).ReadLine()
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	readPassword func() (string, error)
)

func prompt() (string, string, error) { _ = "STUB: not implemented"; return "", "", nil }

func configFile(c *cli.Context) (string, error) { _ = "STUB: not implemented"; return "", nil }

func getConfig(c *cli.Context) (string, *mastodon.Config, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func authenticate(client *mastodon.Client, config *mastodon.Config, file string) error {
	_ = "STUB: not implemented"
	return nil
}

func argstr(c *cli.Context) string { _ = "STUB: not implemented"; return "" }

func fatalIf(err error) { _ = "STUB: not implemented"; return }

func makeApp() *cli.App { _ = "STUB: not implemented"; return nil }

type screen struct {
	host string
}

func newScreen(config *mastodon.Config) *screen { _ = "STUB: not implemented"; return nil }

func (s *screen) acct(a string) string { _ = "STUB: not implemented"; return "" }

func (s *screen) displayError(w io.Writer, e error) { _ = "STUB: not implemented"; return }

func (s *screen) displayStatus(w io.Writer, t *mastodon.Status) { _ = "STUB: not implemented"; return }

func run() int { _ = "STUB: not implemented"; return 0 }

func main() {
	os.Exit(run())
}
