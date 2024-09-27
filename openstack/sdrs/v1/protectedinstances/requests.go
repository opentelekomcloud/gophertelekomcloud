package protectedinstances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

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

	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return InstancePage{pagination.SinglePageBase(r)}
	})
}
