package http_tester

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"codenut.org/http-tester/pkg/http_tester/input"
	"codenut.org/http-tester/pkg/http_tester/trace"
)

func httpRequestFactory() []*http.Request {
	var requests []*http.Request
	for _, scenario := range input.LoadScenario() {
		req, err := http.NewRequest(scenario.Method, scenario.Url, strings.NewReader(scenario.Data))
		if err != nil {
			log.Panic(err)
		}
		for key, value := range scenario.Headers {
			req.Header.Set(key, value)
		}
		requests = append(requests, req)
	}
	return requests
}

func doTracedHttpRequest(req *http.Request, cap *trace.Capture) *trace.Capture {
	res, err := client.Do(cap.StartTrace(req))
	if err != nil {
		log.Panic(err)
	}
	for key, value := range res.Header {
		cap.Headers = append(cap.Headers, fmt.Sprintf("%s: %s", key, value))
	}
	if data, err := ioutil.ReadAll(res.Body); err == nil {
		cap.Data = string(data)
	}
	_ = res.Body.Close()
	cap.StopTrace(res.StatusCode)

	return cap
}
