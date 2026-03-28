package api

import "fmt"

type APIError struct {
	Message    string
	StatusCode int
}

func (e *APIError) Error() string {
	if e.StatusCode > 0 {
		return fmt.Sprintf("API error (HTTP %d): %s", e.StatusCode, e.Message)
	}
	return fmt.Sprintf("API error: %s", e.Message)
}

func (e *APIError) ExitCode() int {
	switch {
	case e.StatusCode == 429:
		return 3
	case e.StatusCode >= 400 && e.StatusCode < 500:
		return 4
	default:
		return 1
	}
}

func AsAPIError(err error) (*APIError, bool) {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr, true
	}
	return nil, false
}
