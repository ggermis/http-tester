package interpolator

import (
	"fmt"
	"strings"
)

type Interpolator interface {
	Register(string, string)
	Interpolate(string) string
}

type defaultInterpolator struct {
	registry map[string]string
}

func (i defaultInterpolator) Register(key, value string) {
	i.registry[key] = value
}

func (i defaultInterpolator) Interpolate(text string) string {
	for key, value := range i.registry {
		text = strings.ReplaceAll(text, fmt.Sprintf("${%s}", key), value)
	}
	return text
}

func NewInterpolator() Interpolator {
	return defaultInterpolator{registry: make(map[string]string)}
}
