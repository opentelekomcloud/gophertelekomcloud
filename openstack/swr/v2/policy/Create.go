package policy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Retention policy matching algorithm. The value is "or".
	Algorithm string `json:"algorithm" required:"true"`
	// List of retention rules.
	Rules []Rule `json:"rules" required:"true"`
}

// This function is used to create an image retention policy.
func Create(client *golangsdk.ServiceClient, organization string, repository string, opts CreateOpts) (int, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return -1, err
	}

	// POST /v2/manage/namespaces/{namespace}/repos/{repository}/retentions
	raw, err := client.Post(client.ServiceURL("manage", "namespaces", organization, "repos", repository, "retentions"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return -1, err
	}

	var res struct {
		ID int `json:"id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.ID, err
}
