package i18n

const (
	// Root
	MsgRootShort Key = "MsgRootShort"
	MsgRootLong  Key = "MsgRootLong"

	// Flags
	FlagOutputUsage Key = "FlagOutputUsage"
	FlagLangUsage   Key = "FlagLangUsage"
	FlagJibunUsage  Key = "FlagJibunUsage"

	// Groups
	GroupSearch Key = "GroupSearch"
	GroupCache  Key = "GroupCache"
	GroupUtil   Key = "GroupUtil"

	// Table headers
	HdrPostcode  Key = "HdrPostcode"
	HdrAddress   Key = "HdrAddress"
	HdrKoAddress Key = "HdrKoAddress"
	HdrEnAddress Key = "HdrEnAddress"

	// Cache
	MsgCacheShort   Key = "MsgCacheShort"
	MsgCacheClear   Key = "MsgCacheClear"
	MsgCacheStats   Key = "MsgCacheStats"
	MsgCacheCleared Key = "MsgCacheCleared"
	MsgCacheEntries Key = "MsgCacheEntries"
	MsgCacheSize    Key = "MsgCacheSize"

	// Errors
	ErrNoKeyword Key = "ErrNoKeyword"
	ErrAPIFailed Key = "ErrAPIFailed"
	ErrNoResults Key = "ErrNoResults"

	// Tool Schema
	MsgToolSchemaShort Key = "MsgToolSchemaShort"

	// Update
	MsgUpdateShort Key = "MsgUpdateShort"
)

var ko = map[Key]string{
	MsgRootShort:       "한국 우편번호 검색 CLI",
	MsgRootLong:        "키워드로 한국 우편번호와 주소를 검색합니다.",
	FlagOutputUsage:    "출력 형식: auto, table, json, jsonl, csv",
	FlagLangUsage:      "주소 언어: ko, en, all",
	FlagJibunUsage:     "지번 주소 출력 (기본: 도로명)",
	GroupSearch:        "검색:",
	GroupCache:         "캐시:",
	GroupUtil:          "유틸리티:",
	HdrPostcode:        "우편번호",
	HdrAddress:         "주소",
	HdrKoAddress:       "한국어 주소",
	HdrEnAddress:       "영문 주소",
	MsgCacheShort:      "캐시 관리",
	MsgCacheClear:      "캐시 전체 삭제",
	MsgCacheStats:      "캐시 통계",
	MsgCacheCleared:    "캐시가 삭제되었습니다.",
	MsgCacheEntries:    "캐시 항목: %d건",
	MsgCacheSize:       "캐시 크기: %s",
	ErrNoKeyword:       "검색어를 입력하세요.",
	ErrAPIFailed:       "API 호출 실패: %s",
	ErrNoResults:       "검색 결과가 없습니다.",
	MsgToolSchemaShort: "AI Agent용 JSON Schema 출력",
	MsgUpdateShort:     "최신 버전으로 업데이트",
}

var en = map[Key]string{
	MsgRootShort:       "Korean postal code lookup CLI",
	MsgRootLong:        "Search Korean postal codes and addresses by keyword.",
	FlagOutputUsage:    "Output format: auto, table, json, jsonl, csv",
	FlagLangUsage:      "Address language: ko, en, all",
	FlagJibunUsage:     "Show jibun address (default: road address)",
	GroupSearch:        "Search:",
	GroupCache:         "Cache:",
	GroupUtil:          "Utility:",
	HdrPostcode:        "POSTCODE",
	HdrAddress:         "ADDRESS",
	HdrKoAddress:       "KOREAN ADDRESS",
	HdrEnAddress:       "ENGLISH ADDRESS",
	MsgCacheShort:      "Manage cache",
	MsgCacheClear:      "Clear all cached data",
	MsgCacheStats:      "Show cache statistics",
	MsgCacheCleared:    "Cache cleared.",
	MsgCacheEntries:    "Cache entries: %d",
	MsgCacheSize:       "Cache size: %s",
	ErrNoKeyword:       "Please enter a search keyword.",
	ErrAPIFailed:       "API call failed: %s",
	ErrNoResults:       "No results found.",
	MsgToolSchemaShort: "Export JSON Schema for AI agents",
	MsgUpdateShort:     "Update to the latest version",
}
