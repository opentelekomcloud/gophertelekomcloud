package quota

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListQuotasOpts struct {
	Offset                *string `q:"offset"`
	Limit                 *string `q:"limit"`
	EnterpriseProjectName *string `q:"enterprise_project_name"`
}

func ListQuotas(client *golangsdk.ServiceClient, opts ListQuotasOpts) (*QuotasResponse, error) {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("quotas")+query.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res QuotasResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type QuotasResponse struct {
	QuotaList  []Quota `json:"quota_list"`
	TotalCount int     `json:"total_count"`
}

type Quota struct {
	EnterpriseProjectId       string `json:"enterprise_project_id"`
	EnterpriseProjectName     string `json:"enterprise_project_name"`
	InstanceQuota             int    `json:"instance_quota"`
	VcpusQuota                int    `json:"vcpus_quota"`
	RamQuota                  int    `json:"ram_quota"`
	AvailabilityInstanceQuota int    `json:"availability_instance_quota"`
	AvailabilityVcpusQuota    int    `json:"availability_vcpus_quota"`
	AvailabilityRamQuota      int    `json:"availability_ram_quota"`
}
