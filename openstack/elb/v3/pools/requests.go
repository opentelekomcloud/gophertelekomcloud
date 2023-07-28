package pools

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToPoolListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the Pool attributes you want to see returned. SortKey allows you to
// sort by a particular Pool attribute. SortDir sets the direction, and is
// either `asc` or `desc`. Marker and Limit are used for pagination.
type ListOpts struct {
	Description     []string `q:"description"`
	HealthMonitorID []string `q:"healthmonitor_id"`
	LBMethod        []string `q:"lb_algorithm"`
	Protocol        []string `q:"protocol"`
	AdminStateUp    *bool    `q:"admin_state_up"`
	Name            []string `q:"name"`
	ID              []string `q:"id"`
	LoadbalancerID  []string `q:"loadbalancer_id"`
	Limit           int      `q:"limit"`
	Marker          string   `q:"marker"`
	SortKey         string   `q:"sort_key"`
	SortDir         string   `q:"sort_dir"`
}

// ToPoolListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToPoolListQuery() (string, error) {
	q, err := build.QueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

// List returns a Pager which allows you to iterate over a collection of
// pools. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those pools that are owned by the
// tenant who submits the request, unless an admin user submits the request.
func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(client)
	if opts != nil {
		query, err := opts.ToPoolListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.Pager{
		Client:     client,
		InitialURL: url,
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return PoolPage{PageWithInfo: pagination.NewPageWithInfo(r)}
		},
	}
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToPoolCreateMap() (map[string]interface{}, error)
}

// CreateOpts is the common options' struct used in this package's Create
// operation.
type CreateOpts struct {
	// The algorithm used to distribute load between the members of the pool.
	LBMethod string `json:"lb_algorithm" required:"true"`

	// The protocol used by the pool members, you can use either
	// ProtocolTCP, ProtocolHTTP, or ProtocolHTTPS.
	Protocol string `json:"protocol" required:"true"`

	// The Loadbalancer on which the members of the pool will be associated with.
	// Note: one of LoadbalancerID or ListenerID must be provided.
	LoadbalancerID string `json:"loadbalancer_id,omitempty"`

	// The Listener on which the members of the pool will be associated with.
	// Note: one of LoadbalancerID or ListenerID must be provided.
	ListenerID string `json:"listener_id,omitempty"`

	// ProjectID is the UUID of the project who owns the Pool.
	// Only administrative users can specify a project UUID other than their own.
	ProjectID string `json:"project_id,omitempty"`

	// Name of the pool.
	Name string `json:"name,omitempty"`

	// Human-readable description for the pool.
	Description string `json:"description,omitempty"`

	// Persistence is the session persistence of the pool.
	// Omit this field to prevent session persistence.
	Persistence *SessionPersistence `json:"session_persistence,omitempty"`

	SlowStart *SlowStart `json:"slow_start,omitempty"`

	// The administrative state of the Pool. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// Specifies whether to enable deletion protection for the pool.
	DeletionProtectionEnable *bool `json:"member_deletion_protection_enable,omitempty"`

	// Specifies the ID of the VPC where the backend server group works.
	VpcId string `json:"vpc_id,omitempty"`

	// Specifies the type of the backend server group.
	// Values:
	// instance: Any type of backend servers can be added. vpc_id is mandatory.
	// ip: Only cross-VPC backend servers can be added. vpc_id cannot be specified.
	// "": Any type of backend servers can be added.
	Type string `json:"type,omitempty"`
}

// SessionPersistence represents the session persistence feature of the load
// balancing service. It attempts to force connections or requests in the same
// session to be processed by the same member as long as it is active. Three
// types of persistence are supported:
type SessionPersistence struct {
	// The type of persistence mode.
	Type string `json:"type" required:"true"`

	// Name of cookie if persistence mode is set appropriately.
	CookieName string `json:"cookie_name,omitempty"`

	// PersistenceTimeout specifies the stickiness duration, in minutes.
	PersistenceTimeout int `json:"persistence_timeout,omitempty"`
}

type SlowStart struct {
	// Specifies whether to Enable slow start.
	Enable bool `json:"enable" required:"true"`

	// Specifies the slow start Duration, in seconds.
	Duration int `json:"duration" required:"true"`
}

// ToPoolCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToPoolCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "pool")
}

// Create accepts a CreateOpts struct and uses the values to create a new
// load balancer pool.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToPoolCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(rootURL(client), b, &r.Body, nil)
	return
}

// Get retrieves a particular pool based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(resourceURL(client, id), &r.Body, nil)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToPoolUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts is the common options' struct used in this package's Update
// operation.
type UpdateOpts struct {
	// Name of the pool.
	Name *string `json:"name,omitempty"`

	// Human-readable description for the pool.
	Description *string `json:"description,omitempty"`

	// The algorithm used to distribute load between the members of the pool. The
	// current specification supports LBMethodRoundRobin, LBMethodLeastConnections
	// and LBMethodSourceIp as valid values for this attribute.
	LBMethod string `json:"lb_algorithm,omitempty"`

	// Specifies whether to enable sticky sessions.
	Persistence *SessionPersistence `json:"session_persistence,omitempty"`

	// The administrative state of the Pool. The value can only be updated to true.
	// This parameter is unsupported. Please do not use it.
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// Specifies whether to enable slow start.
	// This parameter is unsupported. Please do not use it.
	SlowStart *SlowStart `json:"slow_start,omitempty"`

	// Specifies whether to enable deletion protection for the load balancer.
	DeletionProtectionEnable *bool `json:"member_deletion_protection_enable,omitempty"`

	// Specifies the ID of the VPC where the backend server group works.
	VpcId string `json:"vpc_id,omitempty"`

	// Specifies the type of the backend server group.
	// Values:
	// instance: Any type of backend servers can be added. vpc_id is mandatory.
	// ip: Only cross-VPC backend servers can be added. vpc_id cannot be specified.
	// "": Any type of backend servers can be added.
	Type string `json:"type,omitempty"`
}

// ToPoolUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToPoolUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "pool")
}

// Update allows pools to be updated.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToPoolUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(resourceURL(client, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete will permanently delete a particular pool based on its unique ID.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(resourceURL(client, id), nil)
	return
}
