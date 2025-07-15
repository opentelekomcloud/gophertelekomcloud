package quota

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type quotaOpts struct {
	// Specifies the resource type.
	// endpoint_service: indicates the VPC endpoint service.
	// endpoint: indicates the VPC endpoint.
	Type string `q:"type"`
}

func Get(client *golangsdk.ServiceClient, resourceType string) (*Quota, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("quotas").
		WithQueryParams(&quotaOpts{Type: resourceType}).Build()
	if err != nil {
		return nil, err
	}
	// GET /v1/{project_id}/quotas?type={resource_type}
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Quota
	return &res, extract.Into(raw.Body, &res)
}

type Resources struct {
	// Specifies the resource type. You can query the quota of resources of a specified type by configuring this parameter.
	// endpoint_service: indicates the VPC endpoint service.
	// endpoint: indicates the VPC endpoint.
	Type string `json:"type"`
	// Specifies the number of created resources.
	Used int `json:"used"`
	// Specifies the maximum quota of resources.
	Quota int `json:"quota"`
}

type QuotasResp struct {
	// Lists the resources.
	Resources []Resources `json:"resources"`
}

type Quota struct {
	// Specifies quota details.
	Quotas *QuotasResp `json:"quotas"`
}
