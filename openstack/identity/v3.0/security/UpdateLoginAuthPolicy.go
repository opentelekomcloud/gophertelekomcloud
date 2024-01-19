package security

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func UpdateLoginAuthPolicy(client *golangsdk.ServiceClient, id string, opts LoginPolicy) (*LoginPolicy, error) {
	b, err := build.RequestBody(opts, "login_policy")
	if err != nil {
		return nil, err
	}

	// PUT /v3.0/OS-SECURITYPOLICY/domains/{domain_id}/login-policy
	raw, err := client.Put(client.ServiceURL("OS-SECURITYPOLICY", "domains", id, "login-policy"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res LoginPolicy
	return &res, extract.IntoStructPtr(raw.Body, &res, "login_policy")
}
