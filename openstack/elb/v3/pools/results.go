package pools

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/structs"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// Pool represents a logical set of devices, such as web servers, that you
// group together to receive and process traffic. The load balancing function
// chooses a Member of the Pool according to the configured load balancing
// method to handle the new requests or connections received on the VIP address.
type Pool struct {
	// The load-balancer algorithm, which is round-robin, least-connections, and
	// so on. This value, which must be supported, is dependent on the provider.
	// Round-robin must be supported.
	LBMethod string `json:"lb_algorithm"`

	// The protocol of the Pool, which is TCP, HTTP, or HTTPS.
	Protocol string `json:"protocol"`

	// Description for the Pool.
	Description string `json:"description"`

	// A list of listeners objects IDs.
	Listeners []structs.ResourceRef `json:"listeners"`

	// A list of member objects IDs.
	Members []structs.ResourceRef `json:"members"`

	// The ID of associated health monitor.
	MonitorID string `json:"healthmonitor_id"`

	// The administrative state of the Pool, which is up (true) or down (false).
	AdminStateUp bool `json:"admin_state_up"`

	// Pool name. Does not have to be unique.
	Name string `json:"name"`

	ProjectID string `json:"project_id"`

	// The unique ID for the Pool.
	ID string `json:"id"`

	// A list of load balancer objects IDs.
	Loadbalancers []structs.ResourceRef `json:"loadbalancers"`

	// Indicates whether connections in the same session will be processed by the
	// same Pool member or not.
	Persistence *SessionPersistence `json:"session_persistence"`

	IpVersion string `json:"ip_version"`

	SlowStart *SlowStart `json:"slow_start"`

	// Deletion protection for the pool.
	DeletionProtectionEnable bool `json:"member_deletion_protection_enable"`

	// ID of the VPC where the backend server group works.
	VpcId string `json:"vpc_id"`

	// Type of the backend server group.
	Type string `json:"type"`
}

// PoolPage is the page returned by a pager when traversing over a
// collection of pools.
type PoolPage struct {
	pagination.PageWithInfo
}

// IsEmpty checks whether a PoolPage struct is empty.
func (r PoolPage) IsEmpty() (bool, error) {
	is, err := ExtractPools(r)
	return len(is) == 0, err
}

// ExtractPools accepts a Page struct, specifically a PoolPage struct,
// and extracts the elements into a slice of Pool structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractPools(r pagination.Page) ([]Pool, error) {
	var s []Pool

	err := extract.IntoSlicePtr(bytes.NewReader((r.(PoolPage)).Body), &s, "pools")
	if err != nil {
		return nil, err
	}
	return s, nil
}

type commonResult struct {
	golangsdk.Result
}

// Extract is a function that accepts a result and extracts a pool.
func (r commonResult) Extract() (*Pool, error) {
	s := new(Pool)
	err := r.ExtractIntoStructPtr(s, "pool")
	if err != nil {
		return nil, err
	}
	return s, nil
}

// CreateResult represents the result of a Create operation. Call its Extract
// method to interpret the result as a Pool.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a Get operation. Call its Extract
// method to interpret the result as a Pool.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an Update operation. Call its Extract
// method to interpret the result as a Pool.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a Delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	Err error
}
