package cache

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
	"golang.org/x/text/unicode/norm"
)

const ttl = 24 * time.Hour

type Cache struct {
	db   *sqlx.DB
	path string
}

func dbPath() string {
	if xdg := os.Getenv("XDG_CACHE_HOME"); xdg != "" {
		return filepath.Join(xdg, "kozip", "cache.db")
	}
	home, _ := os.UserHomeDir()
	dotCache := filepath.Join(home, ".cache")
	if info, err := os.Stat(dotCache); err == nil && info.IsDir() {
		return filepath.Join(dotCache, "kozip", "cache.db")
	}
	return filepath.Join(home, ".kozip", "cache.db")
}

func Open() (*Cache, error) {
	p := dbPath()
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		return nil, err
	}

	db, err := sqlx.Open("sqlite", p)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS search_cache (
		keyword    TEXT PRIMARY KEY,
		data       TEXT NOT NULL,
		expires_at INTEGER NOT NULL
	)`)
	if err != nil {
		db.Close()
		return nil, err
	}

	return &Cache{db: db, path: p}, nil
}

func (c *Cache) Close() error {
	return c.db.Close()
}

func normalizeKey(keyword string) string {
	return strings.ToLower(norm.NFC.String(keyword))
}

func (c *Cache) Get(keyword string) ([]byte, bool) {
	key := normalizeKey(keyword)
	var data string
	var expiresAt int64

	err := c.db.QueryRow(
		"SELECT data, expires_at FROM search_cache WHERE keyword = ?", key,
	).Scan(&data, &expiresAt)

	if err != nil {
		return nil, false
	}

	if time.Now().Unix() > expiresAt {
		c.db.Exec("DELETE FROM search_cache WHERE keyword = ?", key)
		return nil, false
	}

	return []byte(data), true
}

func (c *Cache) Set(keyword string, results interface{}) error {
	key := normalizeKey(keyword)
	b, err := json.Marshal(results)
	if err != nil {
		return err
	}

	expiresAt := time.Now().Add(ttl).Unix()
	_, err = c.db.Exec(
		`INSERT OR REPLACE INTO search_cache (keyword, data, expires_at) VALUES (?, ?, ?)`,
		key, string(b), expiresAt,
	)
	return err
}

func (c *Cache) Clear() error {
	_, err := c.db.Exec("DELETE FROM search_cache")
	return err
}

type Stats struct {
	Entries int
	Size    int64
}

func (c *Cache) Stats() (Stats, error) {
	var count int
	if err := c.db.QueryRow("SELECT COUNT(*) FROM search_cache").Scan(&count); err != nil {
		return Stats{}, err
	}

	info, err := os.Stat(c.path)
	if err != nil {
		return Stats{}, err
	}

	return Stats{Entries: count, Size: info.Size()}, nil
}

func FormatSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
