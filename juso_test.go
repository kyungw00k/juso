package juso

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearch(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := ApiResponse{
			Count: 1,
			Results: []ApiResult{
				{
					Postcode5: "06236",
					KoCommon:  "서울특별시 강남구",
					KoDoro:    "강남대로 396",
					KoJibeon:  "역삼동 826-5",
					EnCommon:  "Gangnam-gu, Seoul",
					EnDoro:    "396, Gangnam-daero",
					EnJibeon:  "826-5, Yeoksam-dong",
				},
			},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	results, err := SearchWithOptions(context.Background(), "강남역", &Options{
		BaseURL: server.URL,
	})
	if err != nil {
		t.Fatalf("Search() error: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("len(results) = %d, want 1", len(results))
	}
	if results[0].Postcode5 != "06236" {
		t.Errorf("Postcode5 = %q, want %q", results[0].Postcode5, "06236")
	}
	if results[0].KoAddress == "" {
		t.Error("KoAddress is empty")
	}
	if results[0].KakaoMapURL == "" {
		t.Error("KakaoMapURL is empty")
	}
}

func TestSearch_EmptyResults(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(ApiResponse{Count: 0, Results: []ApiResult{}})
	}))
	defer server.Close()

	results, err := SearchWithOptions(context.Background(), "nonexistent", &Options{
		BaseURL: server.URL,
	})
	if err != nil {
		t.Fatalf("Search() error: %v", err)
	}
	if len(results) != 0 {
		t.Errorf("len(results) = %d, want 0", len(results))
	}
}
