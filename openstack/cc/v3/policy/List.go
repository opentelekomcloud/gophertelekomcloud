package policy

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
	// Filter by status. Value options: AVAILABLE, CANCELING, APPLYING, FAILED, DELETED.
	State []string `q:"state"`
	// Filter by whether the policy is applied.
	IsApplied *bool `q:"is_applied"`
	// Filter by version.
	Version []int `q:"version"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListPolicyResp, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints(client.DomainID, "gcn", "central-network", opts.CentralNetworkId, "policies").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListPolicyResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListPolicyResp struct {
	// Request ID.
	RequestId string `json:"request_id"`
	// Pagination query information.
	PageInfo PageInfo `json:"page_info"`
	// Central network policy list.
	CentralNetworkPolicies []CentralNetworkPolicy `json:"central_network_policies"`
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
