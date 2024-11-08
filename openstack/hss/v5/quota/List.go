package quota

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// Offset from which the query starts. If the value is less than 0, it is automatically converted to 0.
	Offset *int `q:"offset"`
	// Number of items displayed on each page.
	// A value less than or equal to 0 will be automatically converted to 10,
	// and a value greater than 200 will be automatically converted to 200.
	Limit int `q:"limit"`
	// HSS edition. Its values and their meaning are as follows:
	// hss.version.null: none
	// hss.version.enterprise: enterprise edition
	// hss.version.premium: premium edition
	// hss.version.container.enterprise: container edition
	Version string `q:"version"`
	// Type. Its value can be:
	// host_resource
	// container_resource
	Category string `q:"category"`
	// Quota status. It can be:
	// QUOTA_STATUS_NORMAL
	// QUOTA_STATUS_EXPIRED
	// QUOTA_STATUS_FREEZE
	QuotaStatus string `q:"quota_status"`
	// Usage status. It can be:
	// USED_STATUS_IDLE
	// USED_STATUS_USED
	UsedStatus string `q:"used_status"`
	// Server name
	HostName string `q:"host_name"`
	// Specifies the resource ID of the HSS quota.
	ResourceId string `q:"resource_id"`
	// on_demand: pay-per-use
	ChargingMode string `q:"charging_mode"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]QuotaResp, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("billing", "quotas-detail").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET /v5/{project_id}/billing/quotas-detail
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return QuotaPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractQuotas(pages)
}

type QuotaPage struct {
	pagination.NewSinglePageBase
}

func ExtractQuotas(r pagination.NewPage) ([]QuotaResp, error) {
	var s struct {
		Quotas []QuotaResp `json:"data_list"`
	}
	err := extract.Into(bytes.NewReader((r.(QuotaPage)).Body), &s)
	return s.Quotas, err
}

type QuotaResp struct {
	// Resource ID of an HSS quota
	ResourceId string `json:"resource_id"`
	// Resource specification code
	Version string `json:"version"`
	// Quota status
	QuotaStatus string `json:"quota_status"`
	// Usage status.
	UsedStatus string `json:"used_status"`
	// Host ID
	HostId string `json:"host_id"`
	// Server name
	HostName string `json:"host_name"`
	// on_demand: pay-per-use
	ChargingMode string `json:"charging_mode"`
	// Tags
	Tags []tags.ResourceTag `json:"tags"`
	// Expiration time. The value -1 indicates that the resource will not expire.
	ExpireTime int64 `json:"expire_time"`
	// Whether quotas are shared. Its value can be:
	SharedQuota string `json:"shared_quota"`
}
