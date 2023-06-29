package monitors

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToMonitorListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the Monitor attributes you want to see returned. SortKey allows you to
// sort by a particular Monitor attribute. SortDir sets the direction, and is
// either `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
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
	// Specifies the number of records on each page.
	//
	// Minimum: 0
	//
	// Maximum: 2000
	//
	// Default: 2000
	Limit *int `q:"limit"`
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
	PageReverse *bool `q:"page_reverse"`
	// Specifies the health check ID.
	//
	// Multiple IDs can be queried in the format of id=xxx&id=xxx.
	Id []string `q:"id"`
	// Specifies the port used for the health check.
	//
	// Multiple ports can be queried in the format of monitor_port=xxx&monitor_port=xxx.
	MonitorPort []string `q:"monitor_port"`
	// Specifies the domain name to which HTTP requests are sent during the health check.
	//
	// The value can contain only digits, letters, hyphens (-), and periods (.) and must start with a digit or letter.
	//
	// Multiple domain names can be queried in the format of domain_name=xxx&domain_name=xxx.
	DomainName []string `q:"domain_name"`
	// Specifies the health check name.
	//
	// Multiple names can be queried in the format of name=xxx&name=xxx.
	Name []string `q:"name"`
	// Specifies the interval between health checks, in seconds. The value ranges from 1 to 50.
	//
	// Multiple intervals can be queried in the format of delay=xxx&delay=xxx.
	Delay []string `q:"delay"`
	// Specifies the number of consecutive health checks when the health check result of a backend server changes from OFFLINE to ONLINE.
	//
	// Multiple values can be queried in the format of max_retries=xxx&max_retries=xxx.
	MaxRetries []string `q:"max_retries"`
	// Specifies the administrative status of the health check.
	//
	// The value can be true (health check is enabled) or false (health check is disabled).
	AdminStateUp *bool `q:"admin_state_up"`
	// Specifies the number of consecutive health checks when the health check result of a backend server changes from ONLINE to OFFLINE. The value ranges from 1 to 10.
	//
	// Multiple values can be queried in the format of max_retries_down=xxx&max_retries_down=xxx.
	MaxRetriesDown []string `q:"max_retries_down"`
	// Specifies the maximum time required for waiting for a response from the health check, in seconds.
	Timeout *int `q:"timeout"`
	// Specifies the health check protocol. The value can be TCP, UDP_CONNECT, HTTP, or HTTPS.
	//
	// Multiple protocols can be queried in the format of type=xxx&type=xxx.
	Type []string `q:"type"`
	// Specifies the expected HTTP status code. This parameter will take effect only when type is set to HTTP or HTTPS.
	//
	// The value options are as follows:
	//
	// A specific value, for example, 200
	//
	// A list of values that are separated with commas (,), for example, 200, 202
	//
	// A value range, for example, 200-204
	//
	// The default value is 200. Multiple status codes can be queried in the format of expected_codes=xxx&expected_codes=xxx.
	ExpectedCodes []string `q:"expected_codes"`
	// Specifies the HTTP request path for the health check. The value must start with a slash (/), and the default value is /. This parameter is available only when type is set to HTTP.
	//
	// Multiple paths can be queried in the format of url_path=xxx&url_path=xxx.
	UrlPath []string `q:"url_path"`
	// Specifies the HTTP method.
	//
	// The value can be GET, HEAD, POST, PUT, DELETE, TRACE, OPTIONS, CONNECT, or PATCH.
	//
	// Multiple methods can be queried in the format of http_method=xxx&http_method=xxx.
	//
	// This parameter is unsupported. Please do not use it.
	HttpMethod []string `q:"http_method"`
	// Specifies the enterprise project ID.
	//
	// If this parameter is not passed, resources in the default enterprise project are queried, and authentication is performed based on the default enterprise project.
	//
	// If this parameter is passed, its value can be the ID of an existing enterprise project (resources in the specific enterprise project are required) or all_granted_eps (resources in all enterprise projects are queried).
	//
	// Multiple IDs can be queried in the format of enterprise_project_id=xxx&enterprise_project_id=xxx.
	//
	// This parameter is unsupported. Please do not use it.
	EnterpriseProjectId []string `q:"enterprise_project_id"`
}

// List returns a Pager which allows you to iterate over a collection of
// health monitors. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those health monitors that are owned by the
// tenant who submits the request, unless an admin user submits the request.
func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	return pagination.NewPager(client, client.ServiceURL("healthmonitors")+query.String(), func(r pagination.PageResult) pagination.Page {
		return MonitorPage{PageWithInfo: pagination.NewPageWithInfo(r)}
	})
}

// MonitorPage is the page returned by a pager when traversing over a
// collection of health monitors.
type MonitorPage struct {
	pagination.PageWithInfo
}

// IsEmpty checks whether a MonitorPage struct is empty.
func (r MonitorPage) IsEmpty() (bool, error) {
	is, err := ExtractMonitors(r)
	return len(is) == 0, err
}

// ExtractMonitors accepts a Page struct, specifically a MonitorPage struct,
// and extracts the elements into a slice of Monitor structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractMonitors(r pagination.Page) ([]Monitor, error) {
	var res []Monitor
	err := extract.IntoSlicePtr(r.(MonitorPage).BodyReader(), &res, "healthmonitors")
	return res, err
}
