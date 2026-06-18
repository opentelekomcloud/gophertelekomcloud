package global_connection_bandwidth

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListSitesOpts struct {
	// Number of records returned on each page. Value range: 1 to 2000.
	Limit int `q:"limit"`
	// ID of the last record on the previous page.
	Marker string `q:"marker"`
	// Filter by resource IDs.
	ID []string `q:"id"`
	// Site name in English.
	NameEn string `q:"name_en"`
	// Site name in Chinese.
	NameCn string `q:"name_cn"`
	// Site code.
	SiteCode string `q:"site_code"`
	// Site type. Value options: Area, SubArea, Region.
	SiteType string `q:"site_type"`
}

// ListSites queries the site list.
func ListSites(client *golangsdk.ServiceClient, opts ListSitesOpts) (*ListSitesResp, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints(client.DomainID, "gcb", "sites").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListSitesResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListSitesResp struct {
	// Request ID.
	RequestId string `json:"request_id"`
	// Total number of records.
	TotalCount int `json:"total_count"`
	// Pagination query information.
	PageInfo PageInfo `json:"page_info"`
	// Site list.
	SiteInfos []Site `json:"site_infos"`
}

// Site describes a global connection bandwidth site.
type Site struct {
	// Instance ID.
	ID string `json:"id"`
	// Resource description.
	Description string `json:"description"`
	// Time when the resource was created.
	CreatedAt string `json:"created_at"`
	// Time when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// Region ID.
	RegionId string `json:"region_id"`
	// English site name.
	NameEn string `json:"name_en"`
	// Chinese site name.
	NameCn string `json:"name_cn"`
	// Site code.
	SiteCode string `json:"site_code"`
	// Site type. Value options: Area, SubArea, Region.
	SiteType string `json:"site_type"`
	// Comma-separated supported services.
	ServiceList string `json:"service_list"`
	// Center or edge site designation.
	PublicBorderGroup string `json:"public_border_group"`
	// Groups the site belongs to.
	GroupList []SiteGroupReference `json:"group_list"`
}

// SiteGroupReference describes a site group reference.
type SiteGroupReference struct {
	// Group instance ID.
	ID string `json:"id"`
	// Group description.
	Description string `json:"description"`
	// English group name.
	NameEn string `json:"name_en"`
	// Chinese group name.
	NameCn string `json:"name_cn"`
}
