package ipgroups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// Specifies the number of records on each page.
	//
	// Minimum: 0
	//
	// Maximum: 2000
	//
	// Default: 2000
	Limit int `q:"limit"`
	// Specifies the ID of the last record on the previous page.
	//
	// Note:
	//
	// This parameter must be used together with limit.
	//
	// If this parameter is not specified, the first page will be queried.
	//
	// This parameter cannot be left blank or set to an invalid ID.
	Marker string `q:"marker"`
	// Specifies whether to use reverse query. Values:
	//
	// true: Query the previous page.
	//
	// false (default): Query the next page.
	//
	// Note:
	//
	// This parameter must be used together with limit.
	//
	// If page_reverse is set to true and you want to query the previous page, set the value of marker to the value of previous_marker.
	PageReverse bool `q:"page_reverse"`
	// Specifies the ID of the IP address group.
	ID []string `q:"id"`
	// Specifies the name of the IP address group.
	Name []string `q:"name"`
	// Provides supplementary information about the IP address group.
	Description []string `q:"description"`
	// Lists the IP addresses in the IP address group. Multiple IP addresses are separated with commas.
	IpList []string `q:"ip_list"`
}

// List is used to obtain the parameter ipGroup list
func List(client *golangsdk.ServiceClient, opts ListOpts) pagination.Pager {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	// GET https://{Endpoint}/v3/{project_id}/backups
	return pagination.NewPager(client, client.ServiceURL("ipgroups")+q.String(), func(r pagination.PageResult) pagination.Page {
		return IpGroupPage{PageWithInfo: pagination.NewPageWithInfo(r)}
	})
}

type IpGroupPage struct {
	pagination.PageWithInfo
}

func (r IpGroupPage) IsEmpty() (bool, error) {
	is, err := ExtractIpGroups(r)
	return len(is) == 0, err
}

func ExtractIpGroups(r pagination.Page) ([]IpGroup, error) {
	var res []IpGroup
	err := extract.IntoSlicePtr(r.(IpGroupPage).BodyReader(), &res, "ipgroups")
	return res, err
}
