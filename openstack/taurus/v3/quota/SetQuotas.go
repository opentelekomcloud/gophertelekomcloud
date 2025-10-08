package quota

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type SetQuotasOpts struct {
	QuotaList []SetQuota `json:"quota_list" required:"true"`
}

type SetQuota struct {
	EnterpriseProjectId   string `json:"enterprise_project_id" required:"true"`
	EnterpriseProjectName string `json:"enterprise_project_name" required:"true"`
	InstanceQuota         int    `json:"instance_quota" required:"true"`
	VcpusQuota            int    `json:"vcpus_quota" required:"true"`
	RamQuota              int    `json:"ram_quota" required:"true"`
}

func SetQuotas(client *golangsdk.ServiceClient, opts SetQuotasOpts) (*SetQuotasResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("quotas"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res SetQuotasResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type SetQuotasResponse struct {
	QuotaList []SetQuota `json:"quota_list"`
}
