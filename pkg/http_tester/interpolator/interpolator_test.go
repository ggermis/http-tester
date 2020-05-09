package interpolator_test

import (
	"testing"

	"github.com/ggermis/http-tester/pkg/http_tester/interpolator"
)

func TestInterpolation(t *testing.T) {
	i := interpolator.NewInterpolator()
	i.Register("x", "aaa")
	i.Register("y", "bbb")

	test := map[string]string{
		"":            "",
		"x":           "x",
		"${x}":        "aaa",
		"x${x}x":      "xaaax",
		"x${xy}x":     "x${xy}x",
		"x${x} ${x}y": "xaaa aaay",
		"x${x} ${y}y": "xaaa bbby",
	}

	for original, expected := range test {
		result := i.Interpolate(original)
		if result != expected {
			t.Errorf("Expected '%s' but got '%s'", expected, result)
		}
	}
}
