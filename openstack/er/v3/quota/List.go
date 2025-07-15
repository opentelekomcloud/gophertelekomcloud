package quota

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// You can query the quotas of the following resources:
	// er_instance: total and used quotas of enterprise routers
	// vpc_attachment: total and used quotas of VPC attachments
	// route_table: total and used quotas of route tables
	// static_route: total and used quotas of static routes
	// vpc_er: total and used quotas of enterprise routers that a VPC can be attached to
	// flow_log: total and used quotas of flow logs that can be created for each attachment
	// dc_attachment: total and used quotas of Direct Connect gateway attachments
	// vpn_attachment: total and used quotas of VPN gateway attachments
	// connect_attachment: total and used quotas of Connect gateway attachments. This type of attachments is not supported now.
	// peering_attachment: total and used quotas of peering connection attachments. This type of attachments is not supported now.
	// can_attachment: total and used quotas of intelligent access gateway attachments. This type of attachments is not supported now.
	Type []string `q:"type"`
	// Enterprise router ID
	InstanceID string `q:"erId"`
	// Route table ID
	RouteTableID string `q:"routeTableId"`
	// VPC ID
	VpcID string `q:"vpcId"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]QuotaResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("enterprise-router", "quotas").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []QuotaResponse
	err = extract.IntoSlicePtr(raw.Body, &res, "quotas")
	return res, err
}

type QuotaResponse struct {
	// Quota type
	Type string `json:"quota_key"`
	// Available quota. The value -1 indicates that there is no quota limit.
	AvailableQuota int64 `json:"quota_limit"`
	// Used quota
	UsedQuota int64 `json:"used"`
	// Measurement unit of used quotas
	Unit string `json:"unit"`
}
