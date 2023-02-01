package others

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ShowImageQuota(client *golangsdk.ServiceClient) ([]QuotaInfo, error) {
	// GET /v1/cloudimages/quota
	raw, err := client.Get(client.ServiceURL("cloudimages", "quota"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res struct {
		Resources []QuotaInfo `json:"resources"`
	}
	err = extract.IntoStructPtr(raw.Body, &res, "quotas")
	return res.Resources, err
}

type QuotaInfo struct {
	// Specifies the type of the resource to be queried.
	Type string `json:"type"`
	// Specifies the used quota.
	Used int `json:"used"`
	// Specifies the total quota.
	Quota int `json:"quota"`
	// Specifies the minimum quota.
	Min int `json:"min"`
	// Specifies the maximum quota.
	Max int `json:"max"`
}
