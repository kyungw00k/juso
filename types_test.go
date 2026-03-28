package kozip

import (
	"testing"
)

func TestToAddressResult(t *testing.T) {
	r := ApiResult{
		Postcode5:    "06236",
		Postcode6:    "135080",
		KoCommon:     "서울특별시 강남구",
		KoDoro:       "강남대로 396",
		KoJibeon:     "역삼동 826-5",
		EnCommon:     "Gangnam-gu, Seoul",
		EnDoro:       "396, Gangnam-daero",
		EnJibeon:     "826-5, Yeoksam-dong",
		BuildingName: "강남빌딩",
	}

	result := r.ToAddressResult()

	if result.Postcode5 != "06236" {
		t.Errorf("Postcode5 = %q, want %q", result.Postcode5, "06236")
	}
	if result.KoAddress != "서울특별시 강남구 강남대로 396 (강남빌딩)" {
		t.Errorf("KoAddress = %q", result.KoAddress)
	}
	if result.EnAddress != "396, Gangnam-daero, Gangnam-gu, Seoul" {
		t.Errorf("EnAddress = %q", result.EnAddress)
	}
	if result.KakaoMapURL == "" {
		t.Error("KakaoMapURL is empty")
	}
	if result.NaverMapURL == "" {
		t.Error("NaverMapURL is empty")
	}
}

func TestToAddressResult_NoBuildingName(t *testing.T) {
	r := ApiResult{
		Postcode5: "06236",
		KoCommon:  "서울특별시 강남구",
		KoDoro:    "강남대로 396",
		KoJibeon:  "역삼동 826-5",
		EnCommon:  "Gangnam-gu, Seoul",
		EnDoro:    "396, Gangnam-daero",
		EnJibeon:  "826-5, Yeoksam-dong",
	}

	result := r.ToAddressResult()

	if result.KoAddress != "서울특별시 강남구 강남대로 396" {
		t.Errorf("KoAddress = %q, want no building suffix", result.KoAddress)
	}
}
