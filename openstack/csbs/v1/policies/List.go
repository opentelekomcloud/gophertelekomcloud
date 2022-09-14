package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// List returns a Pager which allows you to iterate over a collection of
// backup policies. It accepts a ListOpts struct, which allows you to
// filter the returned collection for greater efficiency.
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]BackupPolicy, error) {
	query, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	url := client.ServiceURL("policies") + query.String()
	pages, err := pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return BackupPolicyPage{pagination.LinkedPageBase{PageResult: r}}
	}).AllPages()
	if err != nil {
		return nil, err
	}
	policies, err := ExtractBackupPolicies(pages)
	if err != nil {
		return nil, err
	}

	return FilterPolicies(policies, opts)
}
