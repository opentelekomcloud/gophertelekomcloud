package monitors

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToMonitorListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the Monitor attributes you want to see returned. SortKey allows you to
// sort by a particular Monitor attribute. SortDir sets the direction, and is
// either `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	ID            string `q:"id"`
	Name          string `q:"name"`
	TenantID      string `q:"tenant_id"`
	ProjectID     string `q:"project_id"`
	PoolID        string `q:"pool_id"`
	Type          string `q:"type"`
	Delay         int    `q:"delay"`
	Timeout       int    `q:"timeout"`
	MaxRetries    int    `q:"max_retries"`
	HTTPMethod    string `q:"http_method"`
	URLPath       string `q:"url_path"`
	ExpectedCodes string `q:"expected_codes"`
	AdminStateUp  *bool  `q:"admin_state_up"`
	Status        string `q:"status"`
	Limit         int    `q:"limit"`
	Marker        string `q:"marker"`
	SortKey       string `q:"sort_key"`
	SortDir       string `q:"sort_dir"`
}

// ToMonitorListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToMonitorListQuery() (string, error) {
	q, err := build.QueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

// List returns a Pager which allows you to iterate over a collection of
// health monitors. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those health monitors that are owned by the
// tenant who submits the request, unless an admin user submits the request.
func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(client)
	if opts != nil {
		query, err := opts.ToMonitorListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.Pager{
		Client:     client,
		InitialURL: url,
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return MonitorPage{PageWithInfo: pagination.NewPageWithInfo(r)}
		},
	}
}

type Type string

// Constants that represent approved monitoring types.
const (
	TypePING  Type = "PING"
	TypeTCP   Type = "TCP"
	TypeHTTP  Type = "HTTP"
	TypeHTTPS Type = "HTTPS"
)

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
	_, r.Err = client.Post(rootURL(client), b, &r.Body, nil)
	return
}

// Get retrieves a particular Health Monitor based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(resourceURL(client, id), &r.Body, nil)
	return
}

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

// ToMonitorUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToMonitorUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "healthmonitor")
}

// Update is an operation which modifies the attributes of the specified
// Monitor.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToMonitorUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(resourceURL(client, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return
}

// Delete will permanently delete a particular Monitor based on its unique ID.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(resourceURL(client, id), nil)
	return
}
