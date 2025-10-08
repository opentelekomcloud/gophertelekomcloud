package quota

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateQuotasOpts struct {
	QuotaList []SetQuota `json:"quota_list" required:"true"`
}

func UpdateQuotas(client *golangsdk.ServiceClient, opts UpdateQuotasOpts) (*SetQuotasResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("quotas"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res SetQuotasResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
