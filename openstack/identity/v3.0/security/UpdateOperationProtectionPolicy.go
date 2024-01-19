package security

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func UpdateOperationProtectionPolicy(client *golangsdk.ServiceClient, id string, opts ProtectionPolicy) (*ProtectionPolicy, error) {
	b, err := build.RequestBody(opts, "protect_policy")
	if err != nil {
		return nil, err
	}

	// PUT /v3.0/OS-SECURITYPOLICY/domains/{domain_id}/protect-policy
	raw, err := client.Put(client.ServiceURL("OS-SECURITYPOLICY", "domains", id, "protect-policy"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ProtectionPolicy
	return &res, extract.IntoStructPtr(raw.Body, &res, "protect_policy")
}
