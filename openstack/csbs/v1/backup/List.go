package backup

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the attributes you want to see returned. Marker and Limit are used for pagination.
type ListOpts struct {
	Status       string `q:"status"`
	Limit        string `q:"limit"`
	Marker       string `q:"marker"`
	Sort         string `q:"sort"`
	AllTenants   string `q:"all_tenants"`
	Name         string `q:"name"`
	ResourceId   string `q:"resource_id"`
	ResourceName string `q:"resource_name"`
	PolicyId     string `q:"policy_id"`
	VmIp         string `q:"ip"`
	CheckpointId string `q:"checkpoint_id"`
	ID           string
	ResourceType string `q:"resource_type"`
}

// List returns collection of
// backups. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(c *golangsdk.ServiceClient, opts ListOpts) ([]Backup, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}

	pages, err := pagination.NewPager(c, c.ServiceURL("checkpoint_items")+q.String(),
		func(r pagination.PageResult) pagination.Page {
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
		return filterBackupsById(allBackups, opts.ID)
	}

	return allBackups, nil
}

// СsbsBackupPage is the page returned by a pager when traversing over a
// collection of backups.
type СsbsBackupPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of backups has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r СsbsBackupPage) NextPageURL() (string, error) {
	var res []golangsdk.Link
	err := extract.IntoSlicePtr(r.BodyReader(), &res, "checkpoint_items_links")
	if err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(res)
}

// IsEmpty checks whether a СsbsBackupPage struct is empty.
func (r СsbsBackupPage) IsEmpty() (bool, error) {
	is, err := ExtractBackups(r)
	return len(is) == 0, err
}

// ExtractBackups accepts a Page struct, specifically a СsbsBackupPage struct,
// and extracts the elements into a slice of Backup structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractBackups(r pagination.Page) ([]Backup, error) {
	var res []Backup
	err := extract.IntoSlicePtr(r.(СsbsBackupPage).BodyReader(), &res, "checkpoint_items")
	if err != nil {
		return nil, err
	}
	return res, nil
}

func filterBackupsById(backups []Backup, filterId string) ([]Backup, error) {
	var refinedBackups []Backup

	for _, backup := range backups {

		if filterId == backup.Id {
			refinedBackups = append(refinedBackups, backup)
		}
	}

	return refinedBackups, nil
}
