package connection

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Central network ID.
	CentralNetworkId string `json:"-" required:"true"`
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
	// Filter by global connection bandwidth IDs.
	GlobalConnectionBandwidthId []string `q:"global_connection_bandwidth_id"`
	// Filter by bandwidth type. Value options: BandwidthPackage, TestBandwidth.
	BandwidthType string `q:"bandwidth_type"`
	// Filter by connection type. Value options: ER-ER, ER-GDGW, ER-ER_ROUTE_TABLE.
	ConnectionType string `q:"connection_type"`
	// Whether the connection is cross-region.
	IsCrossRegion *bool `q:"is_cross_region"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListConnectionResp, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints(client.DomainID, "gcn", "central-network", opts.CentralNetworkId, "connections").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListConnectionResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListConnectionResp struct {
	// Request ID.
	RequestId string `json:"request_id"`
	// Pagination query information.
	PageInfo PageInfo `json:"page_info"`
	// Central network connection list.
	CentralNetworkConnections []CentralNetworkConnection `json:"central_network_connections"`
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

// CentralNetworkConnection is a connection on a central network.
type CentralNetworkConnection struct {
	// Instance ID.
	ID string `json:"id"`
	// Instance name.
	Name string `json:"name"`
	// Resource description.
	Description string `json:"description"`
	// Account ID.
	DomainId string `json:"domain_id"`
	// Enterprise project ID.
	EnterpriseProjectId string `json:"enterprise_project_id"`
	// Central network ID.
	CentralNetworkId string `json:"central_network_id"`
	// Central network plane ID.
	CentralNetworkPlaneId string `json:"central_network_plane_id"`
	// Global connection bandwidth ID.
	GlobalConnectionBandwidthId string `json:"global_connection_bandwidth_id"`
	// Bandwidth type. Value options: BandwidthPackage, TestBandwidth.
	BandwidthType string `json:"bandwidth_type"`
	// Bandwidth size in Mbit/s.
	BandwidthSize int64 `json:"bandwidth_size"`
	// Connection status.
	State string `json:"state"`
	// Whether the connection is frozen.
	IsFrozen bool `json:"is_frozen"`
	// Connection type. Value options: ER-ER, ER-GDGW, ER-ER_ROUTE_TABLE.
	ConnectionType string `json:"connection_type"`
	// Two connection endpoints.
	ConnectionPointPair []ConnectionPoint `json:"connection_point_pair"`
	// Time when the resource was created.
	CreatedAt string `json:"created_at"`
	// Time when the resource was updated.
	UpdatedAt string `json:"updated_at"`
}

// ConnectionPoint describes a single connection endpoint.
type ConnectionPoint struct {
	// Instance ID.
	ID string `json:"id"`
	// Project ID.
	ProjectId string `json:"project_id"`
	// Region ID.
	RegionId string `json:"region_id"`
	// Geographic site code.
	SiteCode string `json:"site_code"`
	// Connection endpoint ID.
	InstanceId string `json:"instance_id"`
	// Parent resource ID.
	ParentInstanceId string `json:"parent_instance_id"`
	// Resource type. Value options: ER, GDGW, ER_ROUTE_TABLE.
	Type string `json:"type"`
}
