package golangsdk

import (
	"io"
	"net/http"
	"strings"
)

// ServiceClient stores details required to interact with a specific service API implemented by a provider.
// Generally, you'll acquire these by calling the appropriate `New` method on a ProviderClient.
type ServiceClient struct {
	// ProviderClient is a reference to the provider that implements this service.
	*ProviderClient

	// Endpoint is the base URL of the service's API, acquired from a service catalog.
	// It MUST end with a /.
	Endpoint string

	// ResourceBase is the base URL shared by the resources within a service's API. It should include
	// the API version and, like Endpoint, MUST end with a / if set. If not set, the Endpoint is used
	// as-is, instead.
	ResourceBase string

	// This is the service client type (e.g. compute, sharev2).
	// NOTE: FOR INTERNAL USE ONLY. DO NOT SET. GOPHERCLOUD WILL SET THIS.
	// It is only exported because it gets set in a different package.
	Type string

	// The microversion of the service to use. Set this to use a particular microversion.
	Microversion string
}

// ResourceBaseURL returns the base URL of any resources used by this service. It MUST end with a /.
func (client *ServiceClient) ResourceBaseURL() string {
	if client.ResourceBase != "" {
		return client.ResourceBase
	}
	return client.Endpoint
}

// ServiceURL constructs a URL for a resource belonging to this provider.
func (client *ServiceClient) ServiceURL(parts ...string) string {
	return client.ResourceBaseURL() + strings.Join(parts, "/")
}

func (client *ServiceClient) initReqOpts(JSONBody map[string]interface{}, JSONResponse *io.Reader, opts *RequestOpts) {
	opts.JSONBody = JSONBody

	if JSONResponse != nil {
		opts.JSONResponse = JSONResponse
	}

	if opts.MoreHeaders == nil {
		opts.MoreHeaders = make(map[string]string)
	}

	if client.Microversion != "" {
		client.setMicroversionHeader(opts)
	}
}

// Get calls `Request` with the "GET" HTTP verb. Def 200
// JSONResponse Deprecated
func (client *ServiceClient) Get(url string, JSONResponse *io.Reader, opts *RequestOpts) (*http.Response, error) {
	if opts == nil {
		opts = new(RequestOpts)
	}
	client.initReqOpts(nil, JSONResponse, opts)
	return client.Request("GET", url, opts)
}

// Post calls `Request` with the "POST" HTTP verb. Def 201, 202
// JSONResponse Deprecated
func (client *ServiceClient) Post(url string, JSONBody map[string]interface{}, JSONResponse *io.Reader, opts *RequestOpts) (*http.Response, error) {
	if opts == nil {
		opts = new(RequestOpts)
	}
	client.initReqOpts(JSONBody, JSONResponse, opts)
	return client.Request("POST", url, opts)
}

// Put calls `Request` with the "PUT" HTTP verb. Def 201, 202
// JSONResponse Deprecated
func (client *ServiceClient) Put(url string, JSONBody map[string]interface{}, JSONResponse *io.Reader, opts *RequestOpts) (*http.Response, error) {
	if opts == nil {
		opts = new(RequestOpts)
	}
	client.initReqOpts(JSONBody, JSONResponse, opts)
	return client.Request("PUT", url, opts)
}

// Patch calls `Request` with the "PATCH" HTTP verb. Def 200, 204
// JSONResponse Deprecated
func (client *ServiceClient) Patch(url string, JSONBody map[string]interface{}, JSONResponse *io.Reader, opts *RequestOpts) (*http.Response, error) {
	if opts == nil {
		opts = new(RequestOpts)
	}
	client.initReqOpts(JSONBody, JSONResponse, opts)
	return client.Request("PATCH", url, opts)
}

// Delete calls `Request` with the "DELETE" HTTP verb. Def 202, 204
func (client *ServiceClient) Delete(url string, opts *RequestOpts) (*http.Response, error) {
	if opts == nil {
		opts = new(RequestOpts)
	}
	client.initReqOpts(nil, nil, opts)
	return client.Request("DELETE", url, opts)
}

// DeleteWithBody calls `Request` with the "DELETE" HTTP verb. Def 202, 204
func (client *ServiceClient) DeleteWithBody(url string, JSONBody map[string]interface{}, opts *RequestOpts) (*http.Response, error) {
	if opts == nil {
		opts = new(RequestOpts)
	}
	client.initReqOpts(JSONBody, nil, opts)
	return client.Request("DELETE", url, opts)
}

// DeleteWithResponse calls `Request` with the "DELETE" HTTP verb. Def 202, 204
// Deprecated
func (client *ServiceClient) DeleteWithResponse(url string, JSONResponse *io.Reader, opts *RequestOpts) (*http.Response, error) {
	if opts == nil {
		opts = new(RequestOpts)
	}
	client.initReqOpts(nil, JSONResponse, opts)
	return client.Request("DELETE", url, opts)
}

// DeleteWithBodyResp calls `Request` with the "DELETE" HTTP verb. Def 202, 204
// Deprecated
func (client *ServiceClient) DeleteWithBodyResp(url string, JSONBody map[string]interface{}, JSONResponse *io.Reader, opts *RequestOpts) (*http.Response, error) {
	if opts == nil {
		opts = new(RequestOpts)
	}
	client.initReqOpts(JSONBody, JSONResponse, opts)
	return client.Request("DELETE", url, opts)
}

func (client *ServiceClient) setMicroversionHeader(opts *RequestOpts) {
	switch client.Type {
	case "compute":
		opts.MoreHeaders["X-OpenStack-Nova-API-Version"] = client.Microversion
	case "sharev2":
		opts.MoreHeaders["X-OpenStack-Manila-API-Version"] = client.Microversion
	case "volume":
		opts.MoreHeaders["X-OpenStack-Volume-API-Version"] = client.Microversion
	}

	if client.Type != "" {
		opts.MoreHeaders["OpenStack-API-Version"] = client.Type + " " + client.Microversion
	}
}
