package output

import (
	"encoding/json"
	"fmt"
)

type JSONLFormatter struct{}

func (f *JSONLFormatter) Format(data interface{}) error {
	rows, err := toMaps(data)
	if err != nil {
		return err
	}
	for _, row := range rows {
		b, err := json.Marshal(row)
		if err != nil {
			return err
		}
		fmt.Println(string(b))
	}
	return nil
}
