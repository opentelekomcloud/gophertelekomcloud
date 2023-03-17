package ipgroups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// CreateOpts is the common options' struct used in this package's Create
// operation.
type CreateOpts struct {
	// Specifies the IP address group name.
	Name string `json:"name,omitempty"`

	// Provides supplementary information about the IP address group.
	Description string `json:"description,omitempty"`

	// Specifies the project ID of the IP address group.
	ProjectId string `json:"project_id,omitempty"`

	// Specifies the IP addresses or CIDR blocks in the IP address group. [] indicates any IP address.
	IpList []IpGroupOption `json:"ip_list,omitempty"`
}

type IpGroupOption struct {
	// Specifies the IP addresses in the IP address group.
	Ip string `json:"ip" required:"true"`

	// Provides remarks about the IP address group.
	Description string `json:"description"`
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToIpGroupsCreateMap() (map[string]interface{}, error)
}

// ToIpGroupsCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToIpGroupsCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "ipgroup")
}

// Create is an operation which provisions a new IP address group based on the
// configuration defined in the CreateOpts struct. Once the request is
// validated and progress has started on the provisioning process, a
// CreateResult will be returned.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToIpGroupsCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(rootURL(client), b, &r.Body, nil)
	return
}

// Get retrieves a particular IP address group based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(resourceURL(client, id), &r.Body, nil)
	return
}

// Delete will permanently delete a particular LoadBalancer based on its
// unique ID.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(resourceURL(client, id), nil)
	return
}

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToIpGroupListQuery() (string, error)
}

type ListOpts struct {
	Limit       int    `q:"limit"`
	Marker      string `q:"marker"`
	PageReverse bool   `q:"page_reverse"`

	ID          []string `q:"id"`
	Name        []string `q:"name"`
	Description []string `q:"description"`
	IpList      []string `q:"ip_list"`
}

// ToIpGroupListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToIpGroupListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(client)
	if opts != nil {
		query, err := opts.ToIpGroupListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return IpGroupPage{PageWithInfo: pagination.NewPageWithInfo(r)}
	})
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToIpGroupUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents options for updating a IpGroup.
type UpdateOpts struct {
	// Specifies the IP address group name.
	Name string `json:"name,omitempty"`
	// Provides supplementary information about the IP address group.
	Description string `json:"description,omitempty"`
	// Lists the IP addresses in the IP address group.
	IpList []IpGroupOption `json:"ip_list,omitempty"`
}

// ToIpGroupUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToIpGroupUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "listener")
}

// Update is an operation which modifies the attributes of the specified
// IpGroup.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToIpGroupUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(resourceURL(client, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return
}

// BatchOptsBuilder allows extensions to add additional parameters to the
// BatchAction request.
type BatchOptsBuilder interface {
	ToIpGroupsBatchMap() (map[string]interface{}, error)
}

type IpList struct {
	// Specifies the IP addresses in the IP address group.
	Ip string `json:"ip" required:"true"`
}

// BatchDeleteOpts contains all the values needed to perform BatchDelete on the IP address group.
type BatchDeleteOpts struct {
	// Specifies IP addresses that will be deleted from an IP address group in batches.
	IpGroup []IpList `json:"ipgroup"`
}

// ToIpGroupsBatchMap builds a BatchAction request body from BatchOpts.
func (opts BatchDeleteOpts) ToIpGroupsBatchMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// ipListUpdate is used to create or update the ip list of specific ip group.
func ipListUpdate(client *golangsdk.ServiceClient, ipGroupId string, opts UpdateOptsBuilder) (r BatchResult) {
	b, err := opts.ToIpGroupUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(ipListUpdateResourceURL(client, ipGroupId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// ipListDelete is used to create or update the ip list of specific ip group.
func ipListDelete(client *golangsdk.ServiceClient, ipGroupId string, opts BatchOptsBuilder) (r BatchResult) {
	b, err := opts.ToIpGroupsBatchMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(ipListDeleteResourceURL(client, ipGroupId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
