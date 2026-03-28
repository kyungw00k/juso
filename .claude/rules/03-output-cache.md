# Output & Cache

## Output Formats

| Format | Flag | Behavior |
|--------|------|----------|
| auto | `-o auto` (default) | TTY → table, pipe → json |
| table | `-o table` | lipgloss styled, CJK-aware columns (go-runewidth) |
| json | `-o json` | TTY → pretty, pipe → compact |
| jsonl | `-o jsonl` | One JSON object per line |
| csv | `-o csv` | Full fields (AllColumns), nil → empty string |

## Table Column Selection (--lang)

| --lang | Columns |
|--------|---------|
| ko (default) | 우편번호, 주소 (ko_address) |
| en | POSTCODE, ADDRESS (en_address) |
| all | 우편번호, 한국어 주소, 영문 주소 |

CSV/JSON/JSONL always include all fields regardless of --lang.

## SQLite Cache

- Engine: modernc.org/sqlite (pure Go, no CGO)
- TTL: 24 hours
- Key: NFC-normalized + lowercased keyword
- Path priority: $XDG_CACHE_HOME/juso/ → ~/.cache/juso/ → ~/.juso/

## AddressResult Fields

| Field | Description |
|-------|-------------|
| Postcode5 | 5-digit postal code |
| Postcode6 | Legacy 6-digit code |
| KoAddress | Korean road address |
| KoJibun | Korean lot-number address |
| EnAddress | English road address |
| EnJibun | English lot-number address |
| BuildingName | Building name |
| KakaoMapURL | Kakao Map search link |
| NaverMapURL | Naver Map search link |
