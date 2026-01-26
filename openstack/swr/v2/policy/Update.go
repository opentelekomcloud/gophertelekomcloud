package policy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	// Retention policy matching algorithm. The value is "or".
	Algorithm string `json:"algorithm" required:"true"`
	// List of retention rules.
	Rules []Rule `json:"rules" required:"true"`
}

// This function is used to modify an image retention policy.
func Update(client *golangsdk.ServiceClient, organization, repository, policy string, opts UpdateOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// PATCH /v2/manage/namespaces/{namespace}/repos/{repository}/retentions/{retention_id}
	_, err = client.Patch(client.ServiceURL("manage", "namespaces", organization, "repos", repository, "retentions", policy), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	return
}
