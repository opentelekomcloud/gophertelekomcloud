package security

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func UpdatePasswordPolicy(client *golangsdk.ServiceClient, id string, opts PasswordPolicy) (*PasswordPolicy, error) {
	b, err := build.RequestBody(opts, "password_policy")
	if err != nil {
		return nil, err
	}

	// PUT /v3.0/OS-SECURITYPOLICY/domains/{domain_id}/password-policy
	raw, err := client.Put(client.ServiceURL("OS-SECURITYPOLICY", "domains", id, "password-policy"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res PasswordPolicy
	return &res, extract.IntoStructPtr(raw.Body, &res, "password_policy")
}
