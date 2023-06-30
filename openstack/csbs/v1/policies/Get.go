package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// Get will get a single backup policy with specific ID. call the Extract method on the GetResult.
func Get(client *golangsdk.ServiceClient, policyId string) (*BackupPolicy, error) {
	// GET https://{endpoint}/v1/{project_id}/policies/{policy_id}
	raw, err := client.Get(client.ServiceURL("policies", policyId), nil, openstack.StdRequestOpts())

	if err != nil {
		return nil, err
	}

	var res BackupPolicy
	err = extract.IntoStructPtr(raw.Body, &res, "policy")
	return &res, err
}
