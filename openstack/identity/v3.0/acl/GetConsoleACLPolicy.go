package acl

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ConsoleACLPolicyGet(client *golangsdk.ServiceClient, id string) (r *ACLPolicy, err error) {
	// GET  /v3.0/OS-SECURITYPOLICY/domains/{domain_id}/console-acl-policy
	raw, err := client.Get(client.ServiceURL("OS-SECURITYPOLICY", "domains", id, "console-acl-policy"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ACLPolicy
	err = extract.IntoStructPtr(raw.Body, &res, "console_acl_policy")
	return &res, err
}
