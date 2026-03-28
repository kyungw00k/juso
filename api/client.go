package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	kozip "github.com/kyungw00k/kozip"
	"golang.org/x/text/unicode/norm"
)

const (
	baseURL   = "https://api.poesis.kr/post/search.php"
	timeout   = 10 * time.Second
	userAgent = "kozip-cli"
)

type Client struct {
	http *http.Client
}

func NewClient() *Client {
	return &Client{
		http: &http.Client{Timeout: timeout},
	}
}

func (c *Client) Search(ctx context.Context, keyword string) ([]kozip.AddressResult, error) {
	normalized := norm.NFC.String(keyword)
	u := fmt.Sprintf("%s?q=%s", baseURL, url.QueryEscape(normalized))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, &APIError{
			Message:    string(body),
			StatusCode: resp.StatusCode,
		}
	}

	var apiResp kozip.ApiResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, err
	}

	if apiResp.Error != "" {
		return nil, &APIError{Message: apiResp.Error}
	}

	results := make([]kozip.AddressResult, len(apiResp.Results))
	for i, r := range apiResp.Results {
		results[i] = r.ToAddressResult()
	}

	return results, nil
}
