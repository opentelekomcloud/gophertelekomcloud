package capability

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Capability. Multiple capabilities can be queried.
	Capability []string `q:"capability"`
	// Number of records returned on each page. Value range: 1 to 2000.
	Limit int `q:"limit"`
	// ID of the last record on the previous page.
	Marker string `q:"marker"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListCapabilityResp, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints(client.DomainID, "gcn", "capabilities").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListCapabilityResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListCapabilityResp struct {
	// Request ID.
	RequestId string `json:"request_id"`
	// Pagination query information.
	PageInfo PageInfo `json:"page_info"`
	// Capability list.
	Capabilities []CapabilityEntry `json:"capabilities"`
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

// CapabilityEntry describes a central network capability.
type CapabilityEntry struct {
	// Instance ID.
	ID string `json:"id"`
	// Account ID.
	DomainId string `json:"domain_id"`
	// Capability.
	Capability string `json:"capability"`
	// Capability specifications.
	Specifications CapabilitySpecifications `json:"specifications"`
}

// CapabilitySpecifications describes the specifications of a capability.
type CapabilitySpecifications struct {
	// Whether the capability is supported.
	IsSupport bool `json:"is_support"`
	// Supported bandwidth size range.
	SizeRange ConnectionBandwidthSizeRange `json:"size_range"`
	// Supported billing options.
	ChargeMode []string `json:"charge_mode"`
	// Free lines.
	FreeLines []ConnectionBandwidthFreeLine `json:"free_lines"`
	// Supported regions.
	SupportRegions []string `json:"support_regions"`
	// Regions that support IPv6.
	SupportIpv6Regions []string `json:"support_ipv6_regions"`
	// Regions that support DSCP.
	SupportDscpRegions []string `json:"support_dscp_regions"`
	// Regions that support STS5.
	SupportSts5Regions []string `json:"support_sts5_regions"`
	// Supported sites.
	SupportSites []SiteSpecifications `json:"support_sites"`
	// Regions that support freezing.
	SupportFreezeRegions []string `json:"support_freeze_regions"`
}

// ConnectionBandwidthSizeRange describes the bandwidth size range.
type ConnectionBandwidthSizeRange struct {
	// Minimum bandwidth in Mbit/s.
	Min int64 `json:"min"`
	// Maximum bandwidth in Mbit/s.
	Max int64 `json:"max"`
}

// ConnectionBandwidthFreeLine describes a free line.
type ConnectionBandwidthFreeLine struct {
	// Local site code.
	LocalSiteCode string `json:"local_site_code"`
	// Remote site code.
	RemoteSiteCode string `json:"remote_site_code"`
}

// SiteSpecifications describes a supported site.
type SiteSpecifications struct {
	// Region ID.
	RegionId string `json:"region_id"`
	// Site code.
	SiteCode string `json:"site_code"`
}
