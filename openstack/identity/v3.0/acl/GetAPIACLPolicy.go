package acl

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func APIACLPolicyGet(client *golangsdk.ServiceClient, id string) (r *ACLPolicy, err error) {
	// GET  /v3.0/OS-SECURITYPOLICY/domains/{domain_id}/api-acl-policy
	raw, err := client.Get(client.ServiceURL("OS-SECURITYPOLICY", "domains", id, "api-acl-policy"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ACLPolicy
	err = extract.IntoStructPtr(raw.Body, &res, "api_acl_policy")
	return &res, err
}
