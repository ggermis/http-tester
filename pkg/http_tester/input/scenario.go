package input

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/ggermis/http-tester/pkg/http_tester/cli"
	"github.com/ggermis/http-tester/pkg/http_tester/interpolator"
	"gopkg.in/yaml.v2"
)

type Task struct {
	Method    string
	Url       string
	Headers   map[string]string
	Data      string
	Variables map[string]string
}

func (t Task) AsRequest(scenario *Scenario) *http.Request {
	scenario.interpolate(&t)
	req, err := http.NewRequest(t.Method, t.Url, strings.NewReader(t.Data))
	if err != nil {
		panic(err)
	}
	for key, value := range t.Headers {
		req.Header.Set(key, value)
	}
	return req
}

type Scenario struct {
	Interpolator interpolator.Interpolator
	Tasks        []*Task
}

func (s Scenario) interpolate(t *Task) {
	t.Method = s.Interpolator.Interpolate(t.Method)
	t.Url = s.Interpolator.Interpolate(t.Url)
	t.Data = s.Interpolator.Interpolate(t.Data)
	for key := range t.Headers {
		t.Headers[key] = s.Interpolator.Interpolate(t.Headers[key])
	}
}

func LoadScenario() Scenario {
	var scenario Scenario
	if cli.Option.InputFile == "" {
		scenario = Scenario{
			Tasks: []*Task{
				{Method: cli.Option.Method, Url: cli.Option.Url, Headers: cli.Option.HeadersAsMap(), Data: cli.Option.Data},
			},
		}
	} else {
		scenario = loadScenarioFromFile(cli.Option.InputFile)
	}
	scenario.Interpolator = interpolator.NewInterpolator()
	return scenario
}

func loadScenarioFromFile(filename string) Scenario {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var scenario Scenario
	if err := yaml.Unmarshal(data, &scenario.Tasks); err != nil {
		log.Fatal(err)
	}
	return scenario
}
