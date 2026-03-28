---
name: juso
description: "한국 우편번호/주소 검색. 'juso', '우편번호', 'postal code', 'Korean address', '주소 검색', '주소 찾기' 키워드에서 트리거. juso CLI가 설치되어 있을 때 사용."
---

# juso — 한국 우편번호 검색

`juso` CLI로 한국 우편번호와 주소를 검색합니다.

## 사전 조건

juso CLI가 설치되어 있어야 합니다:
```bash
brew install kyungw00k/cli/juso
# 또는
go install github.com/kyungw00k/juso/cmd/juso@latest
```

## 사용법

### 기본 검색

키워드로 우편번호와 도로명 주소를 검색합니다:

```bash
juso 테헤란로
juso 판교역로
juso 강남대로
```

### 영문 주소

`--lang en` 플래그로 영문 주소를 출력합니다. ASCII 키워드는 자동으로 영문 출력됩니다:

```bash
juso 테헤란로 --lang en
```

### 한/영 동시 출력

```bash
juso 판교역로 --lang all
```

### 지번 주소

도로명 대신 지번 주소를 출력합니다:

```bash
juso 테헤란로 --jibun
```

### 출력 형식

| 플래그 | 설명 | 용도 |
|--------|------|------|
| `-o auto` (기본) | TTY → 테이블, 파이프 → JSON | 기본 사용 |
| `-o table` | 테이블 | 사람이 읽을 때 |
| `-o json` | JSON | 프로그래밍/파이프 |
| `-o jsonl` | JSON Lines | 스트리밍 처리 |
| `-o csv` | CSV (전체 필드) | 스프레드시트 내보내기 |

AI 에이전트에서는 `-o json`을 사용하세요:

```bash
juso 테헤란로 -o json
```

### 응답 필드 (JSON)

| 필드 | 설명 |
|------|------|
| `postcode5` | 5자리 우편번호 |
| `postcode6` | 6자리 우편번호 (구) |
| `ko_address` | 한국어 도로명 주소 |
| `ko_jibun` | 한국어 지번 주소 |
| `en_address` | 영문 도로명 주소 |
| `en_jibun` | 영문 지번 주소 |
| `building_name` | 건물명 |
| `kakao_map_url` | 카카오맵 검색 링크 |
| `naver_map_url` | 네이버지도 검색 링크 |

### 캐시

검색 결과는 SQLite에 24시간 캐시됩니다:

```bash
juso cache stats   # 캐시 통계
juso cache clear   # 캐시 삭제
```

### AI Tool Schema

JSON Schema를 내보내 다른 AI 도구에서 사용 가능:

```bash
juso tool-schema          # 전체 스키마
juso tool-schema search   # 검색 스키마만
```

## 사용 예시 시나리오

### "테헤란로 주소를 찾아줘"
```bash
juso 테헤란로 -o json
```

### "판교역 근처 우편번호를 CSV로 내보내줘"
```bash
juso 판교역로 -o csv > pangyo.csv
```

### "강남대로 주소를 영문으로 알려줘"
```bash
juso 강남대로 --lang en -o json
```
