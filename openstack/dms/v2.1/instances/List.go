package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	InstanceId     string `q:"instance_id"`
	Name           string `q:"name"`
	Engine         string `q:"engine"`
	Status         string `q:"status"`
	IncludeFailure string `q:"include_failure"`
	ExactMatchName string `q:"exact_match_name"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints(ResourcePath).WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListResponse struct {
	Instances  []Instance `json:"instances"`
	TotalCount int        `json:"instance_num"`
}
