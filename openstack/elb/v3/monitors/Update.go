package monitors

import "github.com/opentelekomcloud/gophertelekomcloud"

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToMonitorUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts is the common options' struct used in this package's Update
// operation.
type UpdateOpts struct {
	// The time, in seconds, between sending probes to members.
	Delay int `json:"delay,omitempty"`

	// Maximum number of seconds for a Monitor to wait for a ping reply
	// before it times out. The value must be less than the delay value.
	Timeout int `json:"timeout,omitempty"`

	// Number of permissible ping failures before changing the member's
	// status to INACTIVE. Must be a number between 1 and 10.
	MaxRetries int `json:"max_retries,omitempty"`

	MaxRetriesDown int `json:"max_retries_down,omitempty"`

	// URI path that will be accessed if Monitor type is HTTP or HTTPS.
	// Required for HTTP(S) types.
	URLPath string `json:"url_path,omitempty"`

	// Domain Name.
	DomainName string `json:"domain_name,omitempty"`

	// The HTTP method used for requests by the Monitor. If this attribute
	// is not specified, it defaults to "GET". Required for HTTP(S) types.
	HTTPMethod string `json:"http_method,omitempty"`

	// Expected HTTP codes for a passing HTTP(S) Monitor. You can either specify
	// a single status like "200", or a range like "200-202". Required for HTTP(S)
	// types.
	ExpectedCodes string `json:"expected_codes,omitempty"`

	// The Name of the Monitor.
	Name string `json:"name,omitempty"`

	// The administrative state of the Monitor. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// The Port of the Monitor.
	MonitorPort int `json:"monitor_port,omitempty"`

	// The type of probe, which is PING, TCP, HTTP, or HTTPS, that is
	// sent by the load balancer to verify the member state.
	Type string `json:"type,omitempty"`
}

// Update is an operation which modifies the attributes of the specified
// Monitor.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToMonitorUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(client.ServiceURL("healthmonitors", id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return
}

// ToMonitorUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToMonitorUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "healthmonitor")
}
