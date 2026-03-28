package output

import (
	"encoding/json"
	"fmt"
)

type JSONFormatter struct{}

func (f *JSONFormatter) Format(data interface{}) error {
	var b []byte
	var err error
	if IsTTY() {
		b, err = json.MarshalIndent(data, "", "  ")
	} else {
		b, err = json.Marshal(data)
	}
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
