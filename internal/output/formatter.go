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

func NewFormatter(format string, columns []TableColumn) Formatter {
	switch format {
	case "table":
		return &TableFormatter{Columns: columns}
	case "json":
		return &JSONFormatter{}
	case "jsonl":
		return &JSONLFormatter{}
	case "csv":
		return &CSVFormatter{Columns: columns}
	case "auto":
		if IsTTY() {
			return &TableFormatter{Columns: columns}
		}
		return &JSONFormatter{}
	default:
		return &JSONFormatter{}
	}
}
