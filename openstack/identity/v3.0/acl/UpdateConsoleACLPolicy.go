package acl

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ACLPolicy struct {
	DomainId             string                 `json:"-"`
	AllowAddressNetmasks []AllowAddressNetmasks `json:"allow_address_netmasks,omitempty"`
	AllowIPRanges        []AllowIPRanges        `json:"allow_ip_ranges,omitempty"`
}

type AllowAddressNetmasks struct {
	AddressNetmask string `json:"address_netmask" required:"true"`
	Description    string `json:"description,omitempty"`
}

type AllowIPRanges struct {
	IPRange     string `json:"ip_range" required:"true"`
	Description string `json:"description,omitempty"`
}

func ConsoleACLPolicyUpdate(client *golangsdk.ServiceClient, opts ACLPolicy) (r *ACLPolicy, err error) {
	b, err := build.RequestBody(opts, "console_acl_policy")
	if err != nil {
		return nil, err
	}

	// PUT /v3.0/OS-SECURITYPOLICY/domains/{domain_id}/console-acl-policy
	raw, err := client.Put(client.ServiceURL("OS-SECURITYPOLICY", "domains", opts.DomainId, "console-acl-policy"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ACLPolicy

	return &res, extract.IntoStructPtr(raw.Body, &res, "console_acl_policy")
}
