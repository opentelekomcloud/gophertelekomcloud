package advanced

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateQueryOpts struct {
	DomainId    string `json:"-"`
	Name        string `json:"name" required:"true"`
	Description string `json:"description,omitempty"`
	Expression  string `json:"expression" required:"true"`
}

func CreateQuery(client *golangsdk.ServiceClient, opts CreateQueryOpts) (*Query, error) {
	// POST /v1/resource-manager/domains/{domain_id}/stored-queries
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("resource-manager", "domains", opts.DomainId, "stored-queries"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Query

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Query struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Expression  string `json:"expression"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
}
