# Tips

## Implementing default logging and re-authentication attempts

You can implement custom logging and/or limit re-auth attempts by creating a custom HTTP client
like the following and setting it as the provider client's HTTP Client (via the
`golangsdk.ProviderClient.HTTPClient` field):

```go
//...

// LogRoundTripper satisfies the http.RoundTripper interface and is used to
// customize the default gophertelekomcloud RoundTripper to allow for logging.
type LogRoundTripper struct {
	rt                http.RoundTripper
	numReauthAttempts int
}

// newHTTPClient return a custom HTTP client that allows for logging relevant
// information before and after the HTTP request.
func newHTTPClient() http.Client {
	return http.Client{
		Transport: &LogRoundTripper{
			rt: http.DefaultTransport,
		},
	}
}

// RoundTrip performs a round-trip HTTP request and logs relevant information about it.
func (lrt *LogRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	glog.Infof("Request URL: %s\n", request.URL)

	response, err := lrt.rt.RoundTrip(request)
	if response == nil {
		return nil, err
	}

	if response.StatusCode == http.StatusUnauthorized {
		if lrt.numReauthAttempts == 3 {
			return response, fmt.Errorf("Tried to re-authenticate 3 times with no success.")
		}
		lrt.numReauthAttempts++
	}

	glog.Debugf("Response Status: %s\n", response.Status)

	return response, nil
}

endpoint := "https://127.0.0.1/auth"
pc := openstack.NewClient(endpoint)
pc.HTTPClient = newHTTPClient()

//...
```


## Implementing custom objects

T-Cloud Public (former OpenTelekomCloud) request and response objects may differ among services, API
versions, and API operations. New code should follow the operation-per-file
style used by services such as `apigw` and `fgs`: put the operation function,
its request options, and its response type in the same file unless the types are
shared widely in the package.

### Custom request objects

For a request body, define an operation-specific `Opts` struct with JSON tags.
Use `json:"-"` for values that are used only to build the URL and must not be
sent in the body:

```go
type CreateOpts struct {
	GatewayID   string `json:"-"`
	Name        string `json:"name" required:"true"`
	Description string `json:"remark,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*AppResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(
		client.ServiceURL("apigw", "instances", opts.GatewayID, "apps"),
		b,
		nil,
		&golangsdk.RequestOpts{OkCodes: []int{201}},
	)
	if err != nil {
		return nil, err
	}

	var res AppResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
```

For query parameters, define a `ListOpts` or other operation-specific options
struct with `q` tags and pass it to `golangsdk.BuildQueryString(&opts)`.

### Custom response objects

Define response structs with JSON tags that match the API payload and decode the
SDK response body with `extract.Into` or `extract.IntoSlicePtr`:

```go
type AppResp struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"remark"`
}

func Get(client *golangsdk.ServiceClient, gatewayID, appID string) (*AppResp, error) {
	raw, err := client.Get(
		client.ServiceURL("apigw", "instances", gatewayID, "apps", appID),
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}

	var res AppResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
```

## Overriding default `UnmarshalJSON` method

For some response objects, a field may be a custom type or may be allowed to take on
different types. In these cases, overriding the default `UnmarshalJSON` method may be
necessary. To do this, declare the JSON `struct` field tag as "-" and create an `UnmarshalJSON`
method on the type:

```go
type FunctionEvent struct {
	ID          string    `json:"id"`
	LastUpdate time.Time `json:"-"`
}

func (r *FunctionEvent) UnmarshalJSON(b []byte) error {
	type tmp FunctionEvent
	var s struct {
		tmp
		LastUpdate golangsdk.JSONRFC3339MilliNoZ `json:"last_update"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	*r = FunctionEvent(s.tmp)
	r.LastUpdate = time.Time(s.LastUpdate)

	return err
}
```
