package query

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CountOpts struct {
	// Specifies the region ID.
	RegionId string `q:"region_id"`
	// Specifies the project ID.
	ProjectId string `q:"project_id"`
	// Specifies the resource type
	Type string `q:"type"`
	// Specifies the resource ID.
	Id string `q:"id"`
	// Specifies the resource name.
	Name string `q:"name"`
	// Specifies tags. The format is key or key=value.
	Tags []string `q:"tags"`
}

func GetCount(client *golangsdk.ServiceClient, domainId string, opts CountOpts) (*int, error) {
	// GET /v1/resource-manager/domains/{domain_id}/all-resources/count
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("resource-manager", "domains", domainId, "all-resources", "count").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res struct {
		Count int `json:"total_count"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.Count, err
}
