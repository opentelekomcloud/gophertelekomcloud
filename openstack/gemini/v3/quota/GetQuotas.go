package quota

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetQuotas(client *golangsdk.ServiceClient) (*QuotasDetailResponse, error) {
	raw, err := client.Get(client.ServiceURL("quotas"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res QuotasDetailResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type QuotasDetailResponse struct {
	Quotas ShowResourcesListResponseBody `json:"quotas"`
}

type ShowResourcesListResponseBody struct {
	Resources []ShowResourcesDetailResponseBody `json:"resources"`
}

type ShowResourcesDetailResponseBody struct {
	Type  string `json:"type"`
	Quota int    `json:"quota"`
	Used  int    `json:"used"`
}
