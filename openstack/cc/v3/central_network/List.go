package central_network

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Number of records returned on each page. Value range: 1 to 2000.
	Limit int `q:"limit"`
	// ID of the last record on the previous page.
	Marker string `q:"marker"`
	// Keyword for sorting.
	SortKey string `q:"sort_key"`
	// Sorting order. Value options: asc, desc.
	SortDir string `q:"sort_dir"`
	// Filter by resource IDs.
	ID []string `q:"id"`
	// Filter by resource names.
	Name []string `q:"name"`
	// Filter by status.
	State []string `q:"state"`
	// Filter by enterprise project IDs.
	EnterpriseProjectId []string `q:"enterprise_project_id"`
	// Filter by enterprise router IDs.
	EnterpriseRouterId []string `q:"enterprise_router_id"`
	// Filter by attachment IDs.
	AttachmentInstanceId []string `q:"attachment_instance_id"`
	// Filter by global connection bandwidth IDs.
	GlobalConnectionBandwidthId []string `q:"global_connection_bandwidth_id"`
	// Filter by connection IDs.
	ConnectionId []string `q:"connection_id"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListCentralNetworkResp, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints(client.DomainID, "gcn", "central-networks").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListCentralNetworkResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListCentralNetworkResp struct {
	// Request ID.
	RequestId string `json:"request_id"`
	// Pagination query information.
	PageInfo PageInfo `json:"page_info"`
	// Central network list.
	CentralNetworks []CentralNetwork `json:"central_networks"`
}

// PageInfo is the pagination information returned by list operations.
type PageInfo struct {
	// Marker of the next page.
	NextMarker string `json:"next_marker"`
	// Marker of the previous page.
	PreviousMarker string `json:"previous_marker"`
	// Number of records in the current page.
	CurrentCount int `json:"current_count"`
}
