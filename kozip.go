package kozip

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/text/unicode/norm"
)

const (
	defaultBaseURL = "https://api.poesis.kr/post/search.php"
	defaultTimeout = 10 * time.Second
)

// Search searches for Korean postal codes and addresses by keyword.
// This is the main library entry point for external users.
func Search(ctx context.Context, keyword string) ([]AddressResult, error) {
	return SearchWithOptions(ctx, keyword, nil)
}

// Options configures the search behavior.
type Options struct {
	BaseURL string
	Timeout time.Duration
}

// SearchWithOptions searches with custom options.
func SearchWithOptions(ctx context.Context, keyword string, opts *Options) ([]AddressResult, error) {
	baseURL := defaultBaseURL
	timeout := defaultTimeout

	if opts != nil {
		if opts.BaseURL != "" {
			baseURL = opts.BaseURL
		}
		if opts.Timeout > 0 {
			timeout = opts.Timeout
		}
	}

	normalized := norm.NFC.String(keyword)
	u := fmt.Sprintf("%s?q=%s", baseURL, url.QueryEscape(normalized))

	client := &http.Client{Timeout: timeout}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "kozip-go")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error (HTTP %d): %s", resp.StatusCode, string(body))
	}

	var apiResp ApiResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, err
	}

	if apiResp.Error != "" {
		return nil, fmt.Errorf("API error: %s", apiResp.Error)
	}

	results := make([]AddressResult, len(apiResp.Results))
	for i, r := range apiResp.Results {
		results[i] = r.ToAddressResult()
	}

	return results, nil
}
