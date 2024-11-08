package advanced

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateQueryOpts struct {
	DomainId    string `json:"-"`
	QueryId     string `json:"-"`
	Name        string `json:"name" required:"true"`
	Description string `json:"description,omitempty"`
	Expression  string `json:"expression" required:"true"`
}

func UpdateQuery(client *golangsdk.ServiceClient, opts UpdateQueryOpts) (*Query, error) {
	// PUT /v1/resource-manager/domains/{domain_id}/stored-queries/{query_id}
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("resource-manager", "domains", opts.DomainId, "stored-queries", opts.QueryId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Query

	err = extract.Into(raw.Body, &res)
	return &res, err
}
