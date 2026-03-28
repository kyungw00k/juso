# juso

키워드로 한국 우편번호와 주소를 검색하는 CLI 도구이자 Go 라이브러리입니다.

## 설치

### Homebrew

```bash
brew install kyungw00k/cli/juso
```

### Go

```bash
go install github.com/kyungw00k/juso/cmd/juso@latest
```

### Shell Script

```bash
curl -fsSL https://kyungw00k.dev/juso/install.sh | sh
```

## 사용법

```bash
juso 강남역                      # 우편번호 검색
juso 강남역 --lang en            # 영문 주소 출력
juso 강남역 --lang all           # 한/영 동시 출력
juso 역삼동 --jibun              # 지번 주소 출력
juso 강남역 -o json              # JSON 출력
juso 강남역 -o csv > results.csv # CSV 내보내기
```

### 출력 예시

```
$ juso 강남역
우편번호  주소
06252     서울특별시 강남구 강남대로 328
06232     서울특별시 강남구 강남대로 지하 396
06253     서울특별시 강남구 강남대로66길 14
06234     서울특별시 강남구 테헤란로10길 10 (강남역 우정에쉐르)

$ juso 강남역 --lang en
우편번호  주소
06252     328, Gangnam-daero, Gangnam-gu, Seoul
06232     Jiha 396, Gangnam-daero, Gangnam-gu, Seoul
06253     14, Gangnam-daero 66-gil, Gangnam-gu, Seoul
```

## CLI 옵션

| 플래그 | 기본값 | 설명 |
|--------|--------|------|
| `-o, --output` | `auto` | 출력 형식: `auto`, `table`, `json`, `jsonl`, `csv` |
| `--lang` | `ko` | 주소 언어: `ko`, `en`, `all` (ASCII 키워드는 자동으로 `en`) |
| `--jibun` | `false` | 지번 주소 출력 (기본: 도로명) |

### 출력 형식

| 형식 | 설명 |
|------|------|
| `auto` | 터미널이면 `table`, 파이프면 `json` |
| `table` | 사람이 읽기 좋은 테이블 |
| `json` | JSON (터미널: pretty, 파이프: compact) |
| `jsonl` | JSON Lines (스트리밍용) |
| `csv` | CSV (전체 필드 포함) |

### 서브커맨드

```bash
juso cache stats        # 캐시 통계 (건수, 크기)
juso cache clear        # 캐시 전체 삭제
juso tool-schema        # AI Agent용 JSON Schema 출력
juso tool-schema search # 검색 명령어 Schema만
juso update             # 최신 버전으로 업데이트
juso update --check     # 업데이트 확인만
```

## 캐시

검색 결과는 SQLite에 로컬 캐시됩니다 (TTL: 24시간).

캐시 경로 우선순위:
1. `$XDG_CACHE_HOME/juso/cache.db`
2. `~/.cache/juso/cache.db` (`~/.cache` 존재 시)
3. `~/.juso/cache.db`

## 라이브러리로 사용

```go
import "github.com/kyungw00k/juso"

results, err := juso.Search(ctx, "강남역")
for _, r := range results {
    fmt.Println(r.Postcode5, r.KoAddress)
    fmt.Println(r.EnAddress)
    fmt.Println(r.KakaoMapURL)
}
```

### 옵션 지정

```go
results, err := juso.SearchWithOptions(ctx, "강남역", &juso.Options{
    Timeout: 5 * time.Second,
})
```

### 응답 필드

| 필드 | 설명 |
|------|------|
| `Postcode5` | 5자리 우편번호 |
| `Postcode6` | 6자리 우편번호 (구) |
| `KoAddress` | 한국어 도로명 주소 |
| `KoJibun` | 한국어 지번 주소 |
| `EnAddress` | 영문 도로명 주소 |
| `EnJibun` | 영문 지번 주소 |
| `BuildingName` | 건물명 |
| `KakaoMapURL` | 카카오맵 검색 링크 |
| `NaverMapURL` | 네이버지도 검색 링크 |

## AI Agent 연동

`tool-schema` 명령으로 JSON Schema를 내보내 LLM/MCP 도구로 사용할 수 있습니다.

```bash
juso tool-schema
```

## 문서

[https://kyungw00k.dev/juso/](https://kyungw00k.dev/juso/)

## Credits

주소 데이터는 [Postcodify](https://www.poesis.dev/postcodify/) API를 통해 제공됩니다.
Postcodify는 [poesis](https://github.com/poesis/postcodify)가 개발한 한국 우편번호 검색 엔진입니다.

## License

MIT

juso는 Postcodify의 HTTP API를 호출하며, Postcodify 소스코드를 포함하거나 배포하지 않습니다.
