package quota

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// List all the quotas
func List(client *golangsdk.ServiceClient) ([]Quota, error) {
	// GET /v1.0/tms/quotas
	raw, err := client.Get(client.ServiceURL("tms", "quotas"), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res []Quota
	err = extract.IntoSlicePtr(raw.Body, &res, "quotas")
	return res, err
}

type Quota struct {
	// Specifies the quota key.
	QuotaKey string `json:"quota_key"`
	// Specifies the quota value.
	QuotaLimit int `json:"quota_limit"`
	// Specifies the used quota.
	Used int `json:"used"`
	// Specifies the unit.
	Unit string `json:"unit"`
}
