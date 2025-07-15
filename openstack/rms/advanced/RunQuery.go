package advanced

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type RunQueryOpts struct {
	DomainId   string `json:"-"`
	Expression string `json:"expression" required:"true"`
}

func RunQuery(client *golangsdk.ServiceClient, opts RunQueryOpts) (*QueryResponse, error) {
	// POST /v1/resource-manager/domains/{domain_id}/run-query
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("resource-manager", "domains", opts.DomainId, "run-query"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res QueryResponse

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type QueryResponse struct {
	QueryInfo QueryInfo     `json:"query_info"`
	Results   []interface{} `json:"results"`
}

type QueryInfo struct {
	SelectFields []string `json:"select_fields"`
}
