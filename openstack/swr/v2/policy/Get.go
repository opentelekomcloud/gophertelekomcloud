package policy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to query details about an image retention policy.
func Get(client *golangsdk.ServiceClient, organization, repository, policy string) (*ImageRetentionPolicy, error) {
	// GET /v2/manage/namespaces/{namespace}/repos/{repository}/retentions/{retention_id}
	raw, err := client.Get(client.ServiceURL("manage", "namespaces", organization, "repos", repository, "retentions", policy), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ImageRetentionPolicy
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ImageRetentionPolicy struct {
	// Retention policy matching rule. The value is or.
	Algorithm string `json:"algorithm"`
	// ID.
	ID int `json:"id"`
	// Image retention policy.
	Rules []Rule `json:"rules"`
	// Reserved field.
	Scope string `json:"scope"`
}
