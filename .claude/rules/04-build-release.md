# Build & Release

## Commands

```bash
make build    # Build to ./build/juso
make test     # Run all tests
make install  # Install to ~/.local/bin/
make lint     # golangci-lint
make clean    # Remove build artifacts
```

## Release Pipeline

- Conventional Commits → go-semantic-release → GoReleaser
- `feat:` → minor, `fix:` → patch, `BREAKING CHANGE:` → major
- Homebrew: `brew install kyungw00k/cli/juso`
- Binaries: linux/darwin/windows × amd64/arm64

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 3 | Rate limit (HTTP 429) |
| 4 | Validation error (HTTP 4xx) |

## Demo GIF

Cast file: `docs/demo.cast` (asciinema v2)
Convert: `agg docs/demo.cast docs/demo.gif --font-size 16 --cols 100 --rows 35`

## API

Data source: Postcodify (https://www.poesis.dev/postcodify/)
Endpoint: https://api.poesis.kr/post/search.php?q={keyword}
