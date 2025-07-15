package acl

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func APIACLPolicyUpdate(client *golangsdk.ServiceClient, opts ACLPolicy) (r *ACLPolicy, err error) {
	b, err := build.RequestBody(opts, "api_acl_policy")
	if err != nil {
		return nil, err
	}

	// PUT /v3.0/OS-SECURITYPOLICY/domains/{domain_id}/api-acl-policy
	raw, err := client.Put(client.ServiceURL("OS-SECURITYPOLICY", "domains", opts.DomainId, "api-acl-policy"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ACLPolicy

	return &res, extract.IntoStructPtr(raw.Body, &res, "api_acl_policy")
}
