package output

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CSVFormatter struct {
	Columns []TableColumn
}

func (f *CSVFormatter) Format(data interface{}) error {
	rows, err := toMaps(data)
	if err != nil {
		return err
	}

	w := csv.NewWriter(os.Stdout)

	header := make([]string, len(f.Columns))
	for i, col := range f.Columns {
		header[i] = col.Key
	}
	if err := w.Write(header); err != nil {
		return err
	}

	for _, row := range rows {
		record := make([]string, len(f.Columns))
		for i, col := range f.Columns {
			record[i] = fmt.Sprintf("%v", row[col.Key])
		}
		if err := w.Write(record); err != nil {
			return err
		}
	}

	w.Flush()
	return w.Error()
}
