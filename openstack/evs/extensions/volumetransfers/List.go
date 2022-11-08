package volumetransfers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOpts holds options for listing Transfers. It is passed to the transfers.List function.
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
func List(client *golangsdk.ServiceClient, opts ListOpts) pagination.Pager {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	// GET /v3/{project_id}/os-volume-transfer
	return pagination.NewPager(client, client.ServiceURL("os-volume-transfer")+q.String(), func(r pagination.PageResult) pagination.Page {
		return TransferPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// TransferPage is a pagination.pager that is returned from a call to the List function.
type TransferPage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if a ListResult contains no Transfers.
func (r TransferPage) IsEmpty() (bool, error) {
	transfers, err := ExtractTransfers(r)
	return len(transfers) == 0, err
}

func (r TransferPage) NextPageURL() (string, error) {
	var res []golangsdk.Link
	err := extract.IntoSlicePtr(r.BodyReader(), &res, "transfers_links")
	if err != nil {
		return "", err
	}

	return golangsdk.ExtractNextURL(res)
}

// ExtractTransfers extracts and returns Transfers. It is used while iterating over a transfers.List call.
func ExtractTransfers(r pagination.Page) ([]Transfer, error) {
	var res []Transfer
	err := extract.IntoSlicePtr(r.(TransferPage).Result.BodyReader(), &res, "transfers")
	return res, err
}
