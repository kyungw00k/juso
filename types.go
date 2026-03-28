package juso

import (
	"fmt"
	"net/url"
)

// ApiResponse is the raw response from poesis.kr API.
type ApiResponse struct {
	Error   string      `json:"error,omitempty"`
	Count   int         `json:"count"`
	Results []ApiResult `json:"results"`
}

// ApiResult is a single result from the API.
type ApiResult struct {
	Postcode5    string `json:"postcode5"`
	Postcode6    string `json:"postcode6"`
	KoCommon     string `json:"ko_common"`
	KoDoro       string `json:"ko_doro"`
	KoJibeon     string `json:"ko_jibeon"`
	EnCommon     string `json:"en_common"`
	EnDoro       string `json:"en_doro"`
	EnJibeon     string `json:"en_jibeon"`
	BuildingName string `json:"building_name,omitempty"`
}

// AddressResult is the enriched result used for output.
type AddressResult struct {
	Postcode5    string `json:"postcode5"`
	Postcode6    string `json:"postcode6"`
	KoAddress    string `json:"ko_address"`
	KoJibun      string `json:"ko_jibun"`
	EnAddress    string `json:"en_address"`
	EnJibun      string `json:"en_jibun"`
	BuildingName string `json:"building_name,omitempty"`
	KakaoMapURL  string `json:"kakao_map_url"`
	NaverMapURL  string `json:"naver_map_url"`
}

// ToAddressResult converts an ApiResult to an AddressResult.
func (r ApiResult) ToAddressResult() AddressResult {
	buildingSuffix := ""
	if r.BuildingName != "" {
		buildingSuffix = fmt.Sprintf(" (%s)", r.BuildingName)
	}

	koAddr := fmt.Sprintf("%s %s%s", r.KoCommon, r.KoDoro, buildingSuffix)
	koJibun := fmt.Sprintf("%s %s%s", r.KoCommon, r.KoJibeon, buildingSuffix)
	enAddr := fmt.Sprintf("%s, %s", r.EnDoro, r.EnCommon)
	enJibun := fmt.Sprintf("%s, %s", r.EnJibeon, r.EnCommon)

	return AddressResult{
		Postcode5:    r.Postcode5,
		Postcode6:    r.Postcode6,
		KoAddress:    koAddr,
		KoJibun:      koJibun,
		EnAddress:    enAddr,
		EnJibun:      enJibun,
		BuildingName: r.BuildingName,
		KakaoMapURL:  fmt.Sprintf("https://map.kakao.com/link/search/%s", url.QueryEscape(koAddr)),
		NaverMapURL:  fmt.Sprintf("https://map.naver.com/v5/search/%s", url.QueryEscape(koAddr)),
	}
}
