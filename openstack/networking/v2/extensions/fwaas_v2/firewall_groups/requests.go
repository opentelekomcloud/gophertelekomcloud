package firewall_groups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
	// "fmt"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToFirewallGroupListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the firewall attributes you want to see returned. SortKey allows you to sort
// by a particular firewall attribute. SortDir sets the direction, and is either
// `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	TenantID        string `q:"tenant_id"`
	Name            string `q:"name"`
	Description     string `q:"description"`
	AdminStateUp    bool   `q:"admin_state_up"`
	Shared          bool   `q:"public"`
	IngressPolicyID string `q:"ingress_firewall_policy_id"`
	EgressPolicyID  string `q:"egress_firewall_policy_id"`
	ID              string `q:"id"`
	Limit           int    `q:"limit"`
	Marker          string `q:"marker"`
	SortKey         string `q:"sort_key"`
	SortDir         string `q:"sort_dir"`
}

// ToFirewallListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToFirewallGroupListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// firewall_groups. It accepts a ListOpts struct, which allows you to filter
// and sort the returned collection for greater efficiency.
//
// Default policy settings return only those firewall_groups that are owned by the
// tenant who submits the request, unless an admin user submits the request.
func List(c *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(c)
	if opts != nil {
		query, err := opts.ToFirewallGroupListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.Pager{
		Client:     c,
		InitialURL: url,
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return FirewallGroupPage{pagination.LinkedPageBase{PageResult: r}}
		},
	}
}

// CreateOptsBuilder is the interface options structs have to satisfy in order
// to be used in the main Create operation in this package. Since many
// extensions decorate or modify the common logic, it is useful for them to
// satisfy a basic interface in order for them to be used.
type CreateOptsBuilder interface {
	ToFirewallGroupCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains all the values needed to create a new firewall_group.
type CreateOpts struct {
	IngressPolicyID string `json:"ingress_firewall_policy_id,omitempty"`
	EgressPolicyID  string `json:"egress_firewall_policy_id,omitempty"`
	// Only required if the caller has an admin role and wants to create a firewall
	// for another tenant.
	TenantID     string `json:"tenant_id,omitempty"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	AdminStateUp *bool  `json:"admin_state_up,omitempty"`
	Shared       *bool  `json:"public,omitempty"`
}

// ToFirewallGroupCreateMap casts a CreateOpts struct to a map.
func (opts CreateOpts) ToFirewallGroupCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "firewall_group")
}

// Create accepts a CreateOpts struct and uses the values to create a new firewall group
func Create(c *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToFirewallGroupCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	// fmt.Printf("Creating %+v.\n", r)
	_, r.Err = c.Post(rootURL(c), b, &r.Body, nil)
	// fmt.Printf("Created %+v.\n", r)
	return
}

// Get retrieves a particular firewall based on its unique ID.
func Get(c *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(resourceURL(c, id), &r.Body, nil)
	return
}

// UpdateOptsBuilder is the interface options structs have to satisfy in order
// to be used in the main Update operation in this package. Since many
// extensions decorate or modify the common logic, it is useful for them to
// satisfy a basic interface in order for them to be used.
type UpdateOptsBuilder interface {
	ToFirewallGroupUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contains the values used when updating a firewall.
type UpdateOpts struct {
	IngressPolicyID string `json:"ingress_firewall_policy_id,omitempty"`
	EgressPolicyID  string `json:"egress_firewall_policy_id,omitempty"`
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	AdminStateUp    *bool  `json:"admin_state_up,omitempty"`
	Shared          *bool  `json:"public,omitempty"`
}

// ToFirewallGroupUpdateMap casts a CreateOpts struct to a map.
func (opts UpdateOpts) ToFirewallGroupUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "firewall_group")
}

// Update allows firewall_groups to be updated.
func Update(c *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToFirewallGroupUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(resourceURL(c, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete will permanently delete a particular firewall based on its unique ID.
func Delete(c *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(resourceURL(c, id), nil)
	return
}
