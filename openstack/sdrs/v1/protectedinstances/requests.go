package protectedinstances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToInstanceCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains all the values needed to create a new instance.
type CreateOpts struct {
	// Group ID
	GroupID string `json:"server_group_id" required:"true"`
	// Server ID
	ServerID string `json:"server_id" required:"true"`
	// Instance Name
	Name string `json:"name" required:"true"`
	// Instance Description
	Description string `json:"description,omitempty"`
	// Cluster ID
	ClusterID string `json:"cluster_id,omitempty"`
	// Subnet ID
	SubnetID string `json:"primary_subnet_id,omitempty"`
	// IP Address
	IpAddress string `json:"primary_ip_address,omitempty"`
	// Flavor ID of the DR site server
	Flavor string `json:"flavorRef,omitempty"`
	// Tags list
	Tags []Tags `json:"tags,omitempty"`
}

type Tags struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// ToInstanceCreateMap builds a create request body from CreateOpts.
func (opts CreateOpts) ToInstanceCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "protected_instance")
}

// Create will create a new Instance based on the values in CreateOpts.
func Create(c *golangsdk.ServiceClient, opts CreateOptsBuilder) (r JobResult) {
	b, err := opts.ToInstanceCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(rootURL(c), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToInstanceUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contains all the values needed to update an Instance.
type UpdateOpts struct {
	// Instance name
	Name string `json:"name" required:"true"`
}

// ToInstanceUpdateMap builds a update request body from UpdateOpts.
func (opts UpdateOpts) ToInstanceUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "protected_instance")
}

// Update accepts a UpdateOpts struct and uses the values to update an Instance.The response code from api is 200
func Update(c *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToInstanceUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(resourceURL(c, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Get retrieves a particular Instance based on its unique ID.
func Get(c *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(resourceURL(c, id), &r.Body, openstack.StdRequestOpts())
	return
}

// DeleteOptsBuilder allows extensions to add additional parameters to the
// Delete request.
type DeleteOptsBuilder interface {
	ToInstanceDeleteMap() (map[string]interface{}, error)
}

// DeleteOpts contains all the values needed to delete an Instance.
type DeleteOpts struct {
	// Delete Target Server
	DeleteTargetServer *bool `json:"delete_target_server,omitempty"`
	// Delete Target Eip
	DeleteTargetEip *bool `json:"delete_target_eip,omitempty"`
}

// ToInstanceDeleteMap builds a update request body from DeleteOpts.
func (opts DeleteOpts) ToInstanceDeleteMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Delete will permanently delete a particular Instance based on its unique ID.
func Delete(c *golangsdk.ServiceClient, id string, opts DeleteOptsBuilder) (r JobResult) {
	b, err := opts.ToInstanceDeleteMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.DeleteWithBodyResp(resourceURL(c, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type ListOptsBuilder interface {
	ToInstanceListQuery() (string, error)
}

type ListOpts struct {
	ServerGroupID        string   `q:"server_group_id"`
	ServerGroupIDs       []string `q:"server_group_ids"`
	ProtectedInstanceIDs []string `q:"protected_instance_ids"`
	Limit                int      `q:"limit"`
	Offset               int      `q:"offset"`
	Status               string   `q:"status"`
	Name                 string   `q:"name"`
	QueryType            string   `q:"query_type"`
	AvailabilityZone     string   `q:"availability_zone"`
}

func (opts ListOpts) ToInstanceListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

func List(c *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(c)
	if opts != nil {
		q, err := opts.ToInstanceListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += q
	}

	return pagination.Pager{
		Client:     c,
		InitialURL: url,
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return InstancePage{pagination.SinglePageBase(r)}
		},
	}
}
