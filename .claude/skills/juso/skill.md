---
name: juso
description: Korean postal code and address lookup. Search by keyword, get postcode + Korean/English road/jibun addresses + map URLs.
version: 1.0.0
---

# juso — 한국 우편번호 검색

키워드로 한국 우편번호와 주소(도로명/지번, 한/영)를 검색합니다.

## 명령어

### 검색 (기본 동작)

| 파라미터 | 타입 | 필수 | 기본값 | 설명 |
|---------|------|------|--------|------|
| keyword | string | O | - | 검색어 (주소, 건물명, 우편번호) |
| --lang | string | X | ko | 주소 언어: ko, en, all |
| --jibun | bool | X | false | 지번 주소 표시 |
| -o | string | X | auto | 출력 형식: auto, table, json, jsonl, csv |

예시:
- `juso 강남역` — 강남역 관련 주소 검색
- `juso gangnam --lang en` — 영문 주소로 검색
- `juso 역삼동 --jibun` — 지번 주소 출력
- `juso 강남역 -o json` — JSON 형식 출력

### 응답 필드

| 필드 | 설명 |
|------|------|
| postcode5 | 5자리 우편번호 |
| ko_address | 한국어 도로명 주소 |
| ko_jibun | 한국어 지번 주소 |
| en_address | 영문 도로명 주소 |
| en_jibun | 영문 지번 주소 |
| building_name | 건물명 |
| kakao_map_url | 카카오맵 링크 |
| naver_map_url | 네이버지도 링크 |

### 캐시 관리

- `juso cache stats` — 캐시 통계 (건수, 크기)
- `juso cache clear` — 캐시 전체 삭제

### JSON Schema

- `juso tool-schema` — 전체 명령어 JSON Schema
- `juso tool-schema search` — 검색 명령어 Schema만
