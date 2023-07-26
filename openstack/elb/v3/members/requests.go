package members

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToMembersListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections
// through the API. Filtering is achieved by passing in struct field values
// that map to the Member attributes you want to see returned. SortKey allows
// you to sort by a particular Member attribute. SortDir sets the direction,
// and is either `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	Name            string `q:"name"`
	Weight          int    `q:"weight"`
	AdminStateUp    *bool  `q:"admin_state_up"`
	SubnetID        string `q:"subnet_sidr_id"`
	Address         string `q:"address"`
	ProtocolPort    int    `q:"protocol_port"`
	ID              string `q:"id"`
	OperatingStatus string `q:"operating_status"`
	Limit           int    `q:"limit"`
	Marker          string `q:"marker"`
	SortKey         string `q:"sort_key"`
	SortDir         string `q:"sort_dir"`
}

// ToMembersListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToMembersListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// members. It accepts a ListOptsBuilder, which allows you to filter and
// sort the returned collection for greater efficiency.
//
// Default policy settings return only those members that are owned by the
// tenant who submits the request, unless an admin user submits the request.
func List(client *golangsdk.ServiceClient, poolID string, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(client, poolID)
	if opts != nil {
		query, err := opts.ToMembersListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.Pager{
		Client:     client,
		InitialURL: url,
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return MemberPage{PageWithInfo: pagination.NewPageWithInfo(r)}
		},
	}
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToMemberCreateMap() (map[string]any, error)
}

// CreateOpts is the common options' struct used in this package's CreateMember
// operation.
type CreateOpts struct {
	// The IP address of the member to receive traffic from the load balancer.
	Address string `json:"address" required:"true"`

	// The port on which to listen for client traffic.
	ProtocolPort int `json:"protocol_port" required:"true"`

	// Name of the Member.
	Name string `json:"name,omitempty"`

	// ProjectID is the UUID of the project who owns the Member.
	// Only administrative users can specify a project UUID other than their own.
	ProjectID string `json:"project_id,omitempty"`

	// Specifies the weight of the backend server.
	//
	// Requests are routed to backend servers in the same backend server group based on their weights.
	//
	// If the weight is 0, the backend server will not accept new requests.
	//
	// This parameter is invalid when lb_algorithm is set to SOURCE_IP for the backend server group that contains the backend server.
	Weight *int `json:"weight,omitempty"`

	// If you omit this parameter, LBaaS uses the vip_subnet_id parameter value
	// for the subnet UUID.
	SubnetID string `json:"subnet_cidr_id,omitempty"`

	// The administrative state of the Pool. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
}

// ToMemberCreateMap builds a request body from CreateOptsBuilder.
func (opts CreateOpts) ToMemberCreateMap() (map[string]any, error) {
	return golangsdk.BuildRequestBody(opts, "member")
}

// Create will create and associate a Member with a particular Pool.
func Create(client *golangsdk.ServiceClient, poolID string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToMemberCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(rootURL(client, poolID), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	return
}

// Get retrieves a particular Pool Member based on its unique ID.
func Get(client *golangsdk.ServiceClient, poolID string, memberID string) (r GetResult) {
	_, r.Err = client.Get(resourceURL(client, poolID, memberID), &r.Body, nil)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// List request.
type UpdateOptsBuilder interface {
	ToMemberUpdateMap() (map[string]any, error)
}

// UpdateOpts is the common options' struct used in this package's Update
// operation.
type UpdateOpts struct {
	// Name of the Member.
	Name *string `json:"name,omitempty"`

	// A positive integer value that indicates the relative portion of traffic
	// that this member should receive from the pool. For example, a member with
	// a weight of 10 receives five times as much traffic as a member with a
	// weight of 2.
	Weight *int `json:"weight,omitempty"`

	// The administrative state of the Pool. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
}

// ToMemberUpdateMap builds a request body from UpdateOptsBuilder.
func (opts UpdateOpts) ToMemberUpdateMap() (map[string]any, error) {
	return golangsdk.BuildRequestBody(opts, "member")
}

// Update allows Member to be updated.
func Update(client *golangsdk.ServiceClient, poolID string, memberID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToMemberUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(resourceURL(client, poolID, memberID), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	return
}

// Delete will remove and disassociate a Member from a particular
// Pool.
func Delete(client *golangsdk.ServiceClient, poolID string, memberID string) (r DeleteResult) {
	_, r.Err = client.Delete(resourceURL(client, poolID, memberID), nil)
	return
}
