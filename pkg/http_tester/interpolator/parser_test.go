package interpolator_test

import (
	"testing"

	"github.com/ggermis/http-tester/pkg/http_tester/interpolator"
)

func TestDefaultParser(t *testing.T) {
	parser := interpolator.NewParser("default")
	variables := parser.Parse("some random string", nil)
	if len(variables) > 0 {
		t.Errorf("Default parser should not return a map")
	}
}

func TestJSONParser(t *testing.T) {
	json := `
{
	"a": {
		"b": "c"
    },
    "d": "e"
}
`
	extract := map[string]string{
		"x": "$.a.b",
		"y": "$.d",
	}

	parser := interpolator.NewParser("application/json")
	vars := parser.Parse(json, extract)
	if vars["x"] != "c" {
		t.Errorf("Expected 'x' to have value 'c'")
	}
	if vars["y"] != "e" {
		t.Errorf("Expected 'y' to have value 'e")
	}

}
