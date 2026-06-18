package quota

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Quota type. Multiple quota types can be queried.
	QuotaType []string `q:"quota_type"`
	// Number of records returned on each page. Value range: 1 to 2000.
	Limit int `q:"limit"`
	// ID of the last record on the previous page.
	Marker string `q:"marker"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListQuotaResp, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints(client.DomainID, "gcn", "quotas").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListQuotaResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListQuotaResp struct {
	// Request ID.
	RequestId string `json:"request_id"`
	// Quota list.
	Quotas []CentralNetworkQuota `json:"quotas"`
}

// CentralNetworkQuota describes a single quota entry.
type CentralNetworkQuota struct {
	// Quota identifier.
	QuotaKey string `json:"quota_key"`
	// Quota limit.
	QuotaLimit int `json:"quota_limit"`
	// Used quotas.
	Used int `json:"used"`
	// Unit of the quota value.
	Unit string `json:"unit"`
}
