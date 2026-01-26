package policy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to query image retention policies.
func List(client *golangsdk.ServiceClient, organization, repository string) ([]ImageRetentionPolicy, error) {
	// GET /v2/manage/namespaces/{namespace}/repos/{repository}/retentions}
	raw, err := client.Get(client.ServiceURL("manage", "namespaces", organization, "repos", repository, "retentions"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []ImageRetentionPolicy
	err = extract.Into(raw.Body, &res)
	return res, err
}
