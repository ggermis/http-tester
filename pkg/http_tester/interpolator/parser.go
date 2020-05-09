package interpolator

import (
	"encoding/json"
	"fmt"

	"github.com/yalp/jsonpath"
)

type Parser interface {
	Parse(string, map[string]string) map[string]string
}

func NewParser(mimeType string) Parser {
	switch mimeType {
	case "application/json":
		return &JSONParser{}
	}
	return &DefaultParser{}
}

type DefaultParser struct{}

func (p DefaultParser) Parse(_ string, _ map[string]string) map[string]string {
	return make(map[string]string)
}

type JSONParser struct{}

func (p JSONParser) Parse(text string, variables map[string]string) map[string]string {
	var data interface{}
	err := json.Unmarshal([]byte(text), &data)
	if err != nil {
		panic(err)
	}

	result := make(map[string]string)
	for key, expr := range variables {
		value, err := jsonpath.Read(data, expr)
		if err != nil {
			fmt.Printf("Couldn't find '%s' in response body. Ignoring.", key)
		}
		result[key] = fmt.Sprintf("%v", value)
	}
	return result
}
