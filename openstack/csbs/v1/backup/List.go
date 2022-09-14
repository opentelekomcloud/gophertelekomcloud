package backup

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// List returns collection of
// backups. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(c *golangsdk.ServiceClient, opts ListOpts) ([]Backup, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	u := c.ServiceURL("checkpoint_items") + q.String()
	pages, err := pagination.NewPager(c, u, func(r pagination.PageResult) pagination.Page {
		return СsbsBackupPage{pagination.LinkedPageBase{PageResult: r}}
	}).AllPages()

	if err != nil {
		return nil, err
	}

	allBackups, err := ExtractBackups(pages)
	if err != nil {
		return nil, err
	}

	if opts.ID != "" {
		return FilterBackupsById(allBackups, opts.ID)
	}

	return allBackups, nil

}
