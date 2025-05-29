package common

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func PrintJson(data any) {
	body, _ := json.Marshal(data)
	var str bytes.Buffer
	_ = json.Indent(&str, body, "", "    ")
	fmt.Println(str.String())
}
