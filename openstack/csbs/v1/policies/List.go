package policies

import (
	"bytes"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// Exact matching based on field name
	Name string `q:"name"`
	// The value of sort is a group of properties separated by commas (,) and sorting directions.
	// The value format is <key1>[:<direction>],<key2>[:<direction>], where the value of direction is asc
	// (in ascending order) or desc (in descending order). If the parameter direction is not specified,
	// backup policies are sorted in descending order by time. The value of sort contains a maximum of 255 characters.
	Sort string `q:"sort"`
	// Number of resources displayed per page. The value must be a positive integer. The value defaults to 1000.
	Limit int `q:"limit"`
	// ID of the last record displayed on the previous page when pagination query is applied
	Marker string `q:"marker"`
	// Offset value, which is a positive integer.
	Offset int `q:"offset"`
	// Whether backup policies of all tenants can be queried
	// This parameter is only available for administrators.
	AllTenants string `q:"all_tenants"`
}

// List returns a Pager which allows you to iterate over a collection of
// backup policies. It accepts a ListOpts struct, which allows you to
// filter the returned collection for greater efficiency.
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]BackupPolicy, error) {
	var opts2 interface{} = &opts
	query, err := build.QueryString(opts2)
	if err != nil {
		return nil, err
	}

	// GET https://{endpoint}/v1/{project_id}/policies
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL("policies") + query.String(),
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return BackupPolicyPage{pagination.LinkedPageBase{PageResult: r}}
		},
	}.AllPages()
	if err != nil {
		return nil, err
	}

	policies, err := ExtractBackupPolicies(pages)
	return policies, err
}

// BackupPolicyPage is the page returned by a pager when traversing over a
// collection of backup policies.
type BackupPolicyPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of backup policies has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r BackupPolicyPage) NextPageURL() (string, error) {
	var res []golangsdk.Link
	err := extract.IntoSlicePtr(bytes.NewReader(r.Body), &res, "policies_links")
	if err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(res)
}

// IsEmpty checks whether a BackupPolicyPage struct is empty.
func (r BackupPolicyPage) IsEmpty() (bool, error) {
	is, err := ExtractBackupPolicies(r)
	return len(is) == 0, err
}

// ExtractBackupPolicies accepts a Page struct, specifically a BackupPolicyPage struct,
// and extracts the elements into a slice of Policy structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractBackupPolicies(r pagination.Page) ([]BackupPolicy, error) {
	var res []BackupPolicy
	err := extract.IntoSlicePtr(bytes.NewReader(r.(BackupPolicyPage).Body), &res, "policies")
	return res, err
}
