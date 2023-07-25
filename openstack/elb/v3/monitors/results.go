package monitors

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/structs"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// Monitor represents a load balancer health monitor. A health monitor is used
// to determine whether back-end members of the VIP's pool are usable
// for processing a request. A pool can have several health monitors associated
// with it. There are different types of health monitors supported:
//
// PING: used to ping the members using ICMP.
// TCP: used to connect to the members using TCP.
// HTTP: used to send an HTTP request to the member.
// HTTPS: used to send a secure HTTP request to the member.
//
// When a pool has several monitors associated with it, each member of the pool
// is monitored by all these monitors. If any monitor declares the member as
// unhealthy, then the member status is changed to INACTIVE and the member
// won't participate in its pool's load balancing. In other words, ALL monitors
// must declare the member to be healthy for it to stay ACTIVE.
type Monitor struct {
	// The unique ID for the Monitor.
	ID string `json:"id"`

	// The Name of the Monitor.
	Name string `json:"name"`

	// Specifies the project ID.
	ProjectID string `json:"project_id"`

	// The type of probe sent by the load balancer to verify the member state,
	// which is PING, TCP, HTTP, or HTTPS.
	Type Type `json:"type"`

	// The time, in seconds, between sending probes to members.
	Delay int `json:"delay"`

	// The maximum number of seconds for a monitor to wait for a connection to be
	// established before it times out. This value must be less than the delay
	// value.
	Timeout int `json:"timeout"`

	// Specifies the number of consecutive health checks when the health check result of a backend server changes
	// from OFFLINE to ONLINE.
	MaxRetries int `json:"max_retries"`

	// Specifies the number of consecutive health checks when the health check result of a backend server changes
	// from ONLINE to OFFLINE.
	MaxRetriesDown int `json:"max_retries_down"`

	// The HTTP method that the monitor uses for requests.
	HTTPMethod string `json:"http_method"`

	// The HTTP path of the request sent by the monitor to test the health of a
	// member. Must be a string beginning with a forward slash (/).
	URLPath string `json:"url_path"`

	// Domain Name.
	DomainName string `json:"domain_name"`

	// Expected HTTP codes for a passing HTTP(S) monitor.
	ExpectedCodes string `json:"expected_codes"`

	// The administrative state of the health monitor, which is up (true) or
	// down (false).
	AdminStateUp bool `json:"admin_state_up"`

	// The Port of the Monitor.
	MonitorPort int `json:"monitor_port"`

	// List of pools that are associated with the health monitor.
	Pools []structs.ResourceRef `json:"pools"`
}

// MonitorPage is the page returned by a pager when traversing over a
// collection of health monitors.
type MonitorPage struct {
	pagination.PageWithInfo
}

// IsEmpty checks whether a MonitorPage struct is empty.
func (r MonitorPage) IsEmpty() (bool, error) {
	is, err := ExtractMonitors(r)
	return len(is) == 0, err
}

// ExtractMonitors accepts a Page struct, specifically a MonitorPage struct,
// and extracts the elements into a slice of Monitor structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractMonitors(r pagination.Page) ([]Monitor, error) {
	var s []Monitor

	err := extract.IntoSlicePtr((r.(MonitorPage)).Body, &s, "healthmonitors")
	if err != nil {
		return nil, err
	}
	return s, nil
}

type commonResult struct {
	golangsdk.Result
}

// Extract is a function that accepts a result and extracts a Monitor.
func (r commonResult) Extract() (*Monitor, error) {
	s := new(Monitor)
	err := r.ExtractIntoStructPtr(s, "healthmonitor")
	if err != nil {
		return nil, err
	}
	return s, nil
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Monitor.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Monitor.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Monitor.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the result succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}
