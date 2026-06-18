package global_connection_bandwidth

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Number of records returned on each page. Value range: 1 to 2000.
	Limit int `q:"limit"`
	// ID of the last record on the previous page.
	Marker string `q:"marker"`
	// Filter by resource IDs.
	ID []string `q:"id"`
	// Filter by resource names.
	Name []string `q:"name"`
	// Filter by enterprise project IDs.
	EnterpriseProjectId []string `q:"enterprise_project_id"`
	// Filter by instance IDs.
	InstanceId []string `q:"instance_id"`
	// Filter by instance types.
	InstanceType []string `q:"instance_type"`
	// Filter by service binding types.
	BindingService []string `q:"binding_service"`
	// Filter by bandwidth types.
	Type []string `q:"type"`
	// Filter by status.
	AdminState []string `q:"admin_state"`
	// Filter by billing options.
	ChargeMode []string `q:"charge_mode"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListGlobalConnectionBandwidthResp, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints(client.DomainID, "gcb", "gcbandwidths").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListGlobalConnectionBandwidthResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListGlobalConnectionBandwidthResp struct {
	// Request ID.
	RequestId string `json:"request_id"`
	// Pagination query information.
	PageInfo PageInfo `json:"page_info"`
	// Global connection bandwidth list.
	GlobalConnectionBandwidths []GlobalConnectionBandwidth `json:"globalconnection_bandwidths"`
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
