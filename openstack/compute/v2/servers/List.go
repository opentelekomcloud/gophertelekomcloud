package servers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the server attributes you want to see returned. Marker and Limit are used for pagination.
type ListOpts struct {
	// ChangesSince is a time/date stamp for when the server last changed status.
	ChangesSince string `q:"changes-since"`
	// Image is the name of the image in URL format.
	Image string `q:"image"`
	// Flavor is the name of the flavor in URL format.
	Flavor string `q:"flavor"`
	// Name of the server as a string; can be queried with regular expressions.
	// Realize that ?name=bob returns both bob and bobb. If you need to match bob
	// only, you can use a regular expression matching the syntax of the
	// underlying database server implemented for Compute.
	Name string `q:"name"`
	// Status is the value of the status of the server so that you can filter on "ACTIVE" for example.
	Status string `q:"status"`
	// Host is the name of the host as a string.
	Host string `q:"host"`
	// Marker is a UUID of the server at which you want to set a marker.
	Marker string `q:"marker"`
	// Limit is an integer value for the limit of values to return.
	Limit int `q:"limit"`
	// AllTenants is a bool to show all tenants.
	AllTenants bool `q:"all_tenants"`
	// TenantID lists servers for a particular tenant. Setting "AllTenants = true" is required.
	TenantID string `q:"tenant_id"`
}

// List makes a request against the API to list servers accessible to you.
func List(client *golangsdk.ServiceClient, opts ListOpts) pagination.Pager {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	return pagination.NewPager(client, client.ServiceURL("servers", "detail")+query.String(),
		func(r pagination.PageResult) pagination.Page {
			return ServerPage{pagination.LinkedPageBase{PageResult: r}}
		})
}
