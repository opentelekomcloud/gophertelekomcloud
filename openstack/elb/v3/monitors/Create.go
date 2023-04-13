package monitors

import "github.com/opentelekomcloud/gophertelekomcloud"

// CreateOptsBuilder allows extensions to add additional parameters to the
// List request.
type CreateOptsBuilder interface {
	ToMonitorCreateMap() (map[string]interface{}, error)
}

// CreateOpts is the common options' struct used in this package's Create
// operation.
type CreateOpts struct {
	// The Pool to Monitor.
	PoolID string `json:"pool_id" required:"true"`

	// Specifies the health check protocol.
	//
	// The value can be TCP, UDP_CONNECT, HTTP, HTTPS, or PING.
	Type Type `json:"type" required:"true"`

	// The time, in seconds, between sending probes to members.
	Delay int `json:"delay" required:"true"`

	// Specifies the maximum time required for waiting for a response from the health check, in seconds.
	// It is recommended that you set the value less than that of parameter delay.
	Timeout int `json:"timeout" required:"true"`

	// Specifies the number of consecutive health checks when the health check result of a backend server changes
	// from OFFLINE to ONLINE. The value ranges from 1 to 10.
	MaxRetries int `json:"max_retries" required:"true"`

	// Specifies the number of consecutive health checks when the health check result of a backend server changes
	// from ONLINE to OFFLINE.
	MaxRetriesDown int `json:"max_retries_down,omitempty"`

	// Specifies the HTTP request path for the health check.
	// The value must start with a slash (/), and the default value is /. This parameter is available only when type is set to HTTP.
	URLPath string `json:"url_path,omitempty"`

	// Specifies the domain name that HTTP requests are sent to during the health check.
	// This parameter is available only when type is set to HTTP.
	DomainName string `json:"domain_name,omitempty"`

	// The HTTP method used for requests by the Monitor. If this attribute
	// is not specified, it defaults to "GET".
	HTTPMethod string `json:"http_method,omitempty"`

	// Expected HTTP codes for a passing HTTP(S) Monitor. You can either specify
	// a single status like "200", or a range like "200-202".
	ExpectedCodes string `json:"expected_codes,omitempty"`

	// ProjectID is the UUID of the project who owns the Monitor.
	// Only administrative users can specify a project UUID other than their own.
	ProjectID string `json:"project_id,omitempty"`

	// The Name of the Monitor.
	Name string `json:"name,omitempty"`

	// The administrative state of the Monitor. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// The Port of the Monitor.
	MonitorPort int `json:"monitor_port,omitempty"`
}

// ToMonitorCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToMonitorCreateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "healthmonitor")
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Create is an operation which provisions a new Health Monitor. There are
// different types of Monitor you can provision: PING, TCP or HTTP(S). Below
// are examples of how to create each one.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToMonitorCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(client.ServiceURL("healthmonitors"), b, &r.Body, nil)
	return
}
