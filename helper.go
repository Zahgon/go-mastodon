package mastodon

import (
	"net/http"
	"os"
)

type APIError struct {
	prefix     string
	Message    string
	StatusCode int
}

func (e *APIError) Error() string { _ = "STUB: not implemented"; return "" }

// Base64EncodeFileName returns the base64 data URI format string of the file with the file name.
func Base64EncodeFileName(filename string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Base64Encode returns the base64 data URI format string of the file.
func Base64Encode(file *os.File) (string, error) { _ = "STUB: not implemented"; return "", nil }

// String is a helper function to get the pointer value of a string.
func String(v string) *string { _ = "STUB: not implemented"; return nil }

func parseAPIError(prefix string, resp *http.Response) error { _ = "STUB: not implemented"; return nil }
