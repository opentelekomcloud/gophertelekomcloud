package volumetransfers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToTransferListQuery() (string, error)
}

// ListOpts holds options for listing Transfers. It is passed to the transfers.List
// function.
type ListOpts struct {
	// AllTenants will retrieve transfers of all tenants/projects.
	AllTenants bool `q:"all_tenants"`

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

// ToTransferListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToTransferListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

// List returns Transfers optionally limited by the conditions provided in ListOpts.
func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := client.ServiceURL("os-volume-transfer", "detail")
	if opts != nil {
		query, err := opts.ToTransferListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return TransferPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves the Transfer with the provided ID. To extract the Transfer object
// from the response, call the Extract method on the GetResult.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	resp, err := client.Get(client.ServiceURL("os-volume-transfer", id), &r.Body, nil)
	_, r.Header, r.Err = golangsdk.ParseResponse(resp, err)
	return
}
