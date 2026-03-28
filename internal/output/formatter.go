package output

import (
	"os"

	"golang.org/x/term"
)

type TableColumn struct {
	Header string
	Key    string
}

type Formatter interface {
	Format(data interface{}) error
}

func IsTTY() bool {
	return term.IsTerminal(int(os.Stdout.Fd()))
}

// AllColumns is the full set of columns for CSV/JSONL export.
var AllColumns = []TableColumn{
	{Header: "postcode5", Key: "postcode5"},
	{Header: "postcode6", Key: "postcode6"},
	{Header: "ko_address", Key: "ko_address"},
	{Header: "ko_jibun", Key: "ko_jibun"},
	{Header: "en_address", Key: "en_address"},
	{Header: "en_jibun", Key: "en_jibun"},
	{Header: "building_name", Key: "building_name"},
	{Header: "kakao_map_url", Key: "kakao_map_url"},
	{Header: "naver_map_url", Key: "naver_map_url"},
}

func NewFormatter(format string, columns []TableColumn) Formatter {
	switch format {
	case "table":
		return &TableFormatter{Columns: columns}
	case "json":
		return &JSONFormatter{}
	case "jsonl":
		return &JSONLFormatter{}
	case "csv":
		return &CSVFormatter{Columns: AllColumns}
	case "auto":
		if IsTTY() {
			return &TableFormatter{Columns: columns}
		}
		return &JSONFormatter{}
	default:
		return &JSONFormatter{}
	}
}
