# Http-tester

```
$ go run main.go trace -n 1 -c 1 -u https://site.codenut.org/some/path
```

Uses `httptrace.ClientTrace` to track progress from the HTTP call


## Install

Installs the tool into your `$GOPATH`. Make sure you add it to your `$PATH`

```
make install
```

## Uninstall

Removes the tool from your `$GOPATH`
```
make uninstall
```

## Usage

### Multiple Requests

To perform 5 sequential calls

```
http-tester trace -u https://site.codenut.org/some/path -n 5

```

To perform 5 sequential calls per thread using 3 threads (15 calls in total)

```
http-tester trace -u https://site.codenut.org/some/path -n 5 -c 3
```

Perform requests at randomized intervals to distribute better. The following example will perform 2500 HTTP calls to 
the `/some/path` target on `site.codenut.org`. To distribute traffic better (instead of having 50 requests done simultaneously 
and all wait and retrigger at the same time), specify the `-r` option. This takes a value in ms to wait between requests. In 
this case it will wait anywhere from 0-500ms on each request. Randomized for each
request

```
http-tester trace -u https://site.codenut.org/some/path -n 50 -c 50 -r 500
```

### Output

Default outputter is `dot`

* `null`: Do not print anything 
* `dot`: Prints a `.` for each successfull request (HTTP status 200). It will print the response code otherwise
* `csv`: Print CSV data per request. Can be used for graphing
* `detail`: Print the full trace of every HTTP request

Example

```
http-tester trace -u https://site.codenut.org/some/path -n 2 -o detail
```

### Log only errors

all detail outputs that do not return the expected HTTP status code are written to stderr to make it easier to log them
to a file for examination

```
http-tester trace -u https://site.codenut.org/some/path -n 342 -c 121 -o detail 2>/tmp/errors.log
```

### Custom headers

```
http-tester trace -u https://site.codenut.org/some/path -o detail -H X-A=test -H X-B='string with spaces'
```

### Scenario

It is possible to list HTTP calls in a YAML file. The semantics of the `-n` flag changes slightly. It will run all calls defined in the
YAML file per request

```
- method: GET
  url: https://site.codenut.org/some/path
  headers:
   X-A: a
   X-B: c d e
  data: |-
    some data
    to
    send

- method: PUT
  url: https://site.codenut.org/some/other/path
  headers:
   X-B: a
   X-C: adfc d e
  data: |-
    some other data
    to
    send
```

The following command wil perform 4 actual HTTP requests (2 from the YAML file * request count)

```
$ http-tester trace -f scenario.yml -n 2 -o csv
```