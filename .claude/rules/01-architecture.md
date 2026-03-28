# Architecture

## Package Layout

```
cmd/juso/main.go        → Entrypoint (cli.Execute + exit codes)
internal/cli/            → Cobra commands (CLI-only, not importable)
internal/i18n/           → ko/en message translations (CLI-only)
internal/output/         → Multi-format output: table/json/jsonl/csv (CLI-only)
api/                     → HTTP API client for Postcodify (public, importable)
cache/                   → SQLite cache with TTL (public, importable)
juso.go + types.go       → Library facade + data types (public root package)
```

## Dependency Flow

```
cmd/juso → internal/cli → api/
                        → cache/
                        → internal/i18n
                        → internal/output
                        → juso (root package types)

External users:
  import "github.com/kyungw00k/juso"       → juso.Search(), juso.AddressResult
  import "github.com/kyungw00k/juso/api"   → api.NewClient(), api.Client.Search()
  import "github.com/kyungw00k/juso/cache"  → cache.Open(), cache.Cache
```

## Data Flow

1. User enters keyword → root command
2. Cache lookup (NFC normalized + lowercase key)
3. Cache miss → api.Client.Search() → Postcodify API (https://api.poesis.kr)
4. ApiResult.ToAddressResult() → AddressResult (with map URLs)
5. Cache store → output formatter → stdout
