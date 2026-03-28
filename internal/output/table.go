package output

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"
)

type TableFormatter struct {
	Columns []TableColumn
}

func (f *TableFormatter) Format(data interface{}) error {
	rows, err := toMaps(data)
	if err != nil {
		return err
	}
	if len(rows) == 0 {
		return nil
	}

	widths := make([]int, len(f.Columns))
	for i, col := range f.Columns {
		widths[i] = runewidth.StringWidth(col.Header)
	}
	for _, row := range rows {
		for i, col := range f.Columns {
			val := fmt.Sprintf("%v", row[col.Key])
			w := runewidth.StringWidth(val)
			if w > widths[i] {
				widths[i] = w
			}
		}
	}

	headerStyle := lipgloss.NewStyle().Bold(true)

	var headerParts []string
	for i, col := range f.Columns {
		headerParts = append(headerParts, headerStyle.Render(padRight(col.Header, widths[i])))
	}
	fmt.Fprintln(os.Stdout, strings.Join(headerParts, "  "))

	for _, row := range rows {
		var parts []string
		for i, col := range f.Columns {
			val := fmt.Sprintf("%v", row[col.Key])
			parts = append(parts, padRight(val, widths[i]))
		}
		fmt.Fprintln(os.Stdout, strings.Join(parts, "  "))
	}

	return nil
}

func padRight(s string, width int) string {
	sw := runewidth.StringWidth(s)
	if sw >= width {
		return s
	}
	return s + strings.Repeat(" ", width-sw)
}

func toMaps(data interface{}) ([]map[string]interface{}, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var result []map[string]interface{}
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}
	return result, nil
}
