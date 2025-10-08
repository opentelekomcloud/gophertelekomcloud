package quota

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListProjectQuotasOpts struct {
	Type *string `q:"type"`
}

func ListProjectQuotas(client *golangsdk.ServiceClient, opts ListProjectQuotasOpts) (*ProjectQuotasResponse, error) {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("project-quotas")+query.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ProjectQuotasResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ProjectQuotasResponse struct {
	Quotas ProjectQuotas `json:"quotas"`
}

type ProjectQuotas struct {
	Resources []Resource `json:"resources"`
}

type Resource struct {
	Type  string `json:"type"`
	Used  int    `json:"used"`
	Quota int    `json:"quota"`
}
