package input

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
	"codenut.org/http-tester/pkg/http_tester/cli"
)

type Task struct {
	Method  string
	Url     string
	Headers map[string]string
	Data    string
}

type Scenario []Task

func LoadScenario() Scenario {
	var scenario Scenario
	if cli.Option.InputFile == "" {
		scenario = Scenario{Task{Method: cli.Option.Method, Url: cli.Option.Url, Headers: cli.Option.HeadersAsMap(), Data: cli.Option.Data}}
	} else {
		scenario = loadScenarioFromFile(cli.Option.InputFile)
	}
	return scenario
}

func loadScenarioFromFile(filename string) Scenario {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var scenario Scenario
	if err := yaml.Unmarshal(data, &scenario); err != nil {
		log.Fatal(err)
	}
	return scenario
}
