package endpoints

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListIPsOpts contains the query parameters of the ListIPs function.
type ListIPsOpts struct {
	// Number of resources returned on each page. Value range: 0 to 500. The default value is 500.
	Limit int `q:"limit,omitempty"`
	// The offset of pagination query. It specifies the number of rows or records to skip
	// from the beginning of the result set before retrieving the desired data.
	// Value range: 0 to 2147483647. The default value is 0.
	Offset int `q:"offset,omitempty"`
	// Resource status.
	Status string `q:"status,omitempty"`
}

// ListIPs is used to query IP addresses of an endpoint.
// Send GET /v2.1/endpoints/{endpoint_id}/ipaddresses
func ListIPs(client *golangsdk.ServiceClient, endpointID string, opts ListIPsOpts) ([]IPAddress, error) {
	if endpointID == "" {
		return nil, golangsdk.ErrMissingInput{Argument: "endpointID"}
	}

	url, err := golangsdk.NewURLBuilder().WithEndpoints("endpoints", endpointID, "ipaddresses").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(v21URL(client, url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []IPAddress
	err = extract.IntoSlicePtr(raw.Body, &res, "ipaddresses")
	return res, err
}

// IPAddress is the IpaddressesData object of an endpoint.
type IPAddress struct {
	// Resource status. The value can be: ACTIVE, PENDING_CREATE, PENDING_DELETE, ERROR.
	Status string `json:"status"`
	// IP address ID, which is a resource identifier in UUID format.
	ID string `json:"id"`
	// IP address.
	IP string `json:"ip"`
	// Creation time. Format: yyyy-MM-dd'T'HH:mm:ss.SSS
	CreateTime string `json:"create_time"`
	// Update time. Format: yyyy-MM-dd'T'HH:mm:ss.SSS
	UpdateTime string `json:"update_time"`
	// Subnet ID.
	SubnetID string `json:"subnet_id"`
	// Error message.
	ErrorInfo string `json:"error_info"`
}
