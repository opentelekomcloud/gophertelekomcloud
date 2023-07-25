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
	// Query based on field status is supported.
	// Value range: waiting_protect, protecting, available, waiting_restore, restoring, error, waiting_delete, deleting, and deleted
	Status string `q:"status"`
	// Number of resources displayed per page. The value must be a positive integer. The value defaults to 1000.
	Limit string `q:"limit"`
	// ID of the last record displayed on the previous page
	Marker string `q:"marker"`
	// A group of properties separated by commas (,) and sorting directions. The value format is <key1>[:<direction>],
	// <key2>[:<direction>], where the value of direction is asc (in ascending order) or desc (in descending order).
	// If the parameter direction is not specified, the default sorting direction is desc.
	// The value of sort contains a maximum of 255 characters. Enumeration values of the key are as follows:
	// created_at, updated_at, name, status, protected_at, and id.
	Sort string `q:"sort"`
	// Whether to query the backup of all tenants. Only administrators can query the backup of all tenants.
	AllTenants string `q:"all_tenants"`
	// Fuzzy search based on field name is supported.
	Name string `q:"name"`
	// Filtering based on the backup AZ is supported.
	Az string `q:"az"`
	// Filtering based on the backup object ID is supported.
	ResourceId string `q:"resource_id"`
	// Fuzzy search based on the backup object name is supported.
	ResourceName string `q:"resource_name"`
	// Filtering based on the backup start time is supported.
	// For example: 2017-04-18T01:21:52.701973
	StartTime string `q:"start_time"`
	// Filtering based on the backup end time is supported.
	// For example: 2017-04-18T01:21:52.701973
	EndTime string `q:"end_time"`
	// Supports filtering by image type, for example, backup.
	ImageType string `q:"image_type"`
	// Filtering based on policy_id is supported.
	PolicyId string `q:"policy_id"`
	// Offset value, which is a positive integer.
	Offset string `q:"offset"`
	// Filtering based on checkpoint_id is supported.
	CheckpointId string `q:"checkpoint_id"`
	// Type of the backup object. For example, OS::Nova::Server
	ResourceType string `q:"resource_type"`
	// IP address of the server.
	VmIp string `q:"ip"`
}

// List returns collection of
// backups. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(c *golangsdk.ServiceClient, opts ListOpts) ([]Backup, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}

	// GET https://{endpoint}/v1/{project_id}/checkpoint_items
	pages, err := pagination.Pager{
		Client:     c,
		InitialURL: c.ServiceURL("checkpoint_items") + q.String(),
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return СsbsBackupPage{pagination.LinkedPageBase{PageResult: r}}
		},
	}.AllPages()
	if err != nil {
		return nil, err
	}

	allBackups, err := ExtractBackups(pages)
	return allBackups, err
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
	err := extract.IntoSlicePtr(r.Body, &res, "checkpoint_items_links")
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
	err := extract.IntoSlicePtr(r.(СsbsBackupPage).Body, &res, "checkpoint_items")
	if err != nil {
		return nil, err
	}
	return res, nil
}
