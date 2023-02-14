# Http-tester

```
$ go run main.go -n 1 -c 1 -u https://site.codenut.org/some/path
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

### Standard usage

Put load on a system and check that no errors are returned from the server

```
$ http-tester -u https://site.codenut.org/some/path -n 10 -c 12
..........................[502].............................................................................................

Finished 120 requests in 1.007123483s [119.17 rps]

Response Status Distribution:

200: 119
502: 1

Response Time Distribution:

   0ms -   50ms: 79
  50ms -  100ms: 29
 450ms -  500ms: 9
 500ms -  550ms: 3
```

Check the details of HTTP calls

```
$ http-tester -u https://www.google.be -o detail
  ---
  threadid: 1
  requestid: 1
  method: GET
  url: https://www.google.be
  start: 2020-02-12T14:51:15.704097+01:00
  status: 200
  duration: 445.030204
  tlshandshake: true
  headers:
  - 'X-Xss-Protection: [0]'
  - 'X-Frame-Options: [SAMEORIGIN]'
  - 'Expires: [-1]'
  - 'Cache-Control: [private, max-age=0]'
  - 'Server: [gws]'
  - 'Set-Cookie: [1P_JAR=2020-02-12-13; expires=Fri, 13-Mar-2020 13:51:16 GMT; path=/;
    domain=.google.be; Secure NID=197=GbGhAXkSOoA1pQuFs9JfJVDWSLIttZxUWjVlm8gRoKd31ZwV-10gnIEsD-21Rkyqxloz-U1jGLdSoH4SbtqqOwiHZkB4eFrK4IPOo4pt8NAg4GF8hiVWUjNBHoUcP_FKuSYJ85QevY6p-44Yjjx8RDHBC8sFgCLywDp3gpREmQ8;
    expires=Thu, 13-Aug-2020 13:51:16 GMT; path=/; domain=.google.be; HttpOnly]'
  - 'Alt-Svc: [quic=":443"; ma=2592000; v="46,43",h3-Q050=":443"; ma=2592000,h3-Q049=":443";
    ma=2592000,h3-Q048=":443"; ma=2592000,h3-Q046=":443"; ma=2592000,h3-Q043=":443";
    ma=2592000]'
  - 'Date: [Wed, 12 Feb 2020 13:51:16 GMT]'
  - 'Content-Type: [text/html; charset=ISO-8859-1]'
  - 'P3p: [CP="This is not a P3P policy! See g.co/p3phelp for more info."]'
  data: !!binary |
    PCFkb2N0eXBlIGh0bWw+PGh0bWwgaXRlbXNjb3BlPSIiIGl0ZW10eXBlPSJodHRwOi8vc2
    NoZW1hLm9yZy9XZWJQYWdlIiBsYW5nPSJubC1CRSI+PGhlYWQ+PG1ldGEgY29udGVudD0i
    dGV4dC9odG1sOyBjaGFyc2V0PVVURi04IiBodHRwLWVxdWl2PSJDb250ZW50LVR5cGUiPj
    xtZXRhIGNvbnRlbnQ9Ii9pbWFnZXMvYnJhbmRpbmcvZ29vZ2xlZy8xeC9nb29nbGVnX3N0
    ...
    NceDIyOltdLFx4MjJzYnBsXHgyMjoxNixceDIyc2Jwclx4MjI6MTYsXHgyMnNjZFx4MjI6
    MTAsXHgyMnN0b2tceDIyOlx4MjJYc3lfbmRxVUM5LVFxOHdCcTVja3dqUHZ5ZFFceDIyLF
    x4MjJ1aGRlXHgyMjpmYWxzZX19Jztnb29nbGUucG1jPUpTT04ucGFyc2UocG1jKTt9KSgp
    Ozwvc2NyaXB0PiAgICAgICAgPC9ib2R5PjwvaHRtbD4=
  actions:
  - name: Trace started
    params: []
    duration: 0
    total: 0.002275
  - name: GetConn
    params:
    - www.google.be:443
    duration: 0.21214
    total: 0.214415
  - name: DNSStart
    params:
    - www.google.be
    duration: 0.123115
    total: 0.33753
  - name: DNSDone
    params:
    - - ip: 74.125.193.94
        zone: ""
      - ip: 2a00:1450:400b:c01::5e
        zone: ""
    duration: 58.076564
    total: 58.414094
  - name: ConnectStart
    params:
    - tcp
    - 74.125.193.94:443
    duration: 0.13853800000000405
    total: 58.552632
  - name: ConnectDone
    params:
    - tcp
    - 74.125.193.94:443
    - null
    duration: 78.642994
    total: 137.195626
  - name: TLSHandshakeStart
    params: []
    duration: 0.3156919999999843
    total: 137.511318
  - name: TLSHandshakeDone
    params:
    - '{Version:772 HandshakeComplete:true DidResume:false CipherSuite:4865 NegotiatedProtocol:h2
      NegotiatedProtocolIsMutual:true ServerName: PeerCertificates:[0xc000166000 0xc000166580]
      VerifiedChains:[[0xc000166000 0xc000166580 0xc000208b00]] SignedCertificateTimestamps:[]
      OCSPResponse:[] ekm:0x12390e0 TLSUnique:[]}'
    - null
    duration: 216.98250700000003
    total: 354.493825
  - name: GotConn
    params:
    - conn: {}
      reused: false
      wasidle: false
      idletime: 0s
    duration: 0.12201799999996865
    total: 354.615843
  - name: WroteHeaderField
    params:
    - :authority
    - '[www.google.be]'
    duration: 0.03249700000003486
    total: 354.64834
  - name: WroteHeaderField
    params:
    - :method
    - '[GET]'
    duration: 0.002579999999966276
    total: 354.65092
  - name: WroteHeaderField
    params:
    - :path
    - '[/]'
    duration: 0.0014869999999973516
    total: 354.652407
  - name: WroteHeaderField
    params:
    - :scheme
    - '[https]'
    duration: 0.0014219999999909305
    total: 354.653829
  - name: WroteHeaderField
    params:
    - accept-encoding
    - '[gzip]'
    duration: 0.0028390000000513282
    total: 354.656668
  - name: WroteHeaderField
    params:
    - user-agent
    - '[Go-http-client/2.0]'
    duration: 0.003042999999991025
    total: 354.659711
  - name: WroteHeaders
    params: []
    duration: 0.024038999999959287
    total: 354.68375
  - name: WroteRequest
    params:
    - err: null
    duration: 0.0006490000000098917
    total: 354.684399
  - name: GotFirstResponseByte
    params: []
    duration: 88.62813900000003
    total: 443.312538
  - name: Trace finished
    params: []
    duration: 1.7176660000000084
    total: 445.030204
  

  Finished 1 requests in 450.297678ms [2.22 rps]
  
  Response Status Distribution:
  
  200: 1
  
  Response Time Distribution:
  
   400ms -  450ms: 1
```


### Multiple Requests

To perform 5 sequential calls

```
$ http-tester -u https://site.codenut.org/some/path -n 5
```

To perform 5 sequential calls per thread using 3 threads (15 calls in total)

```
$ http-tester -u https://site.codenut.org/some/path -n 5 -c 3
```

Perform requests at randomized intervals to distribute better. The following example will perform 2500 HTTP calls to 
the `/some/path` target on `site.codenut.org`. To distribute traffic better (instead of having 50 requests done simultaneously 
and all wait and retrigger at the same time), specify the `-r` option. This takes a value in ms to wait between requests. In 
this case it will wait anywhere from 0-500ms on each request. Randomized for each
request

```
$ http-tester -u https://site.codenut.org/some/path -n 50 -c 50 -r 500
```

### Output

Default outputter is `dot`

* `null`: Do not print anything 
* `dot`: Prints a `.` for each successfull request (HTTP status 200). It will print the response code otherwise
* `csv`: Print CSV data per request. Can be used for graphing
* `detail`: Print the full trace of every HTTP request
* `split`: Split total response time into connection and response

Example

```
$ http-tester -u https://site.codenut.org/some/path -n 2 -o detail
```

### Log only errors

all detail outputs that do not return the expected HTTP status code are written to stderr to make it easier to log them
to a file for examination

```
$ http-tester -u https://site.codenut.org/some/path -n 342 -c 121 -o detail 2>/tmp/errors.log
```

### Custom headers

```
$ http-tester -u https://site.codenut.org/some/path -o detail -H X-A=test -H X-B='string with spaces'
```

### Scenario

It is possible to list HTTP calls in a YAML file. The semantics of the `-n` flag changes slightly. It will run all calls defined in the
YAML file per request

When a `JSON` response is detected, the data is parsed and can be referenced in subsequent calls using `jsonpath` syntax. In the below 
example we do an oauth call and use the `access_token` that was returned in the next call

```
- method: POST
  url: https://api.example.com/oauth/token
  headers:
    Authorization: Basic dGVzdDp0ZXN0
    Content-Type: application/x-www-form-urlencoded
  data: |-
    grant_type=password&username=my.user@example.com&password=mypassword
  variables:
    access_token: $.access_token
- method: GET
  url: https://api.example.com/mobile/users/me
  headers:
    Authorization: Bearer ${access_token}
    Content-Type: application/json
```

The following command wil perform 4 actual HTTP requests (2 from the YAML file * request count)

```
$ http-tester -f scenario.yml -n 2 -o csv
```