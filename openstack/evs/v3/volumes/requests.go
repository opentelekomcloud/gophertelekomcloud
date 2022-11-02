package volumes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// Get retrieves the Volume with the provided ID. To extract the Volume object
// from the response, call the Extract method on the GetResult.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	resp, err := client.Get(client.ServiceURL("volumes", id), &r.Body, nil)
	_, r.Header, r.Err = golangsdk.ParseResponse(resp, err)
	return
}

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToVolumeListQuery() (string, error)
}

// ListOpts holds options for listing Volumes. It is passed to the volumes.List
// function.
type ListOpts struct {
	// AllTenants will retrieve volumes of all tenants/projects.
	AllTenants bool `q:"all_tenants"`

	// Metadata will filter results based on specified metadata.
	Metadata map[string]string `q:"metadata"`

	// Name will filter by the specified volume name.
	Name string `q:"name"`

	// Status will filter by the specified status.
	Status string `q:"status"`

	// TenantID will filter by a specific tenant/project ID.
	// Setting AllTenants is required for this.
	TenantID string `q:"project_id"`

	// Comma-separated list of sort keys and optional sort directions in the
	// form of <key>[:<direction>].
	Sort string `q:"sort"`

	// Requests a page size of items.
	Limit int `q:"limit"`

	// Used in conjunction with limit to return a slice of items.
	Offset int `q:"offset"`

	// The ID of the last-seen item.
	Marker string `q:"marker"`
}

// ToVolumeListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToVolumeListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

// List returns Volumes optionally limited by the conditions provided in ListOpts.
func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := client.ServiceURL("volumes", "detail")
	if opts != nil {
		query, err := opts.ToVolumeListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return VolumePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToVolumeUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contain options for updating an existing Volume. This object is passed
// to the volumes.Update function. For more information about the parameters, see
// the Volume object.
type UpdateOpts struct {
	Name        *string           `json:"name,omitempty"`
	Description *string           `json:"description,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

// ToVolumeUpdateMap assembles a request body based on the contents of an
// UpdateOpts.
func (opts UpdateOpts) ToVolumeUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "volume")
}

// Update will update the Volume with provided information. To extract the updated
// Volume from the response, call the Extract method on the UpdateResult.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToVolumeUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Put(client.ServiceURL("volumes", id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	_, r.Header, r.Err = golangsdk.ParseResponse(resp, err)
	return
}
