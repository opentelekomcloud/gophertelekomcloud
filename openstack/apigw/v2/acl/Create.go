package acl

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	GatewayID string `json:"-"`
	// Access control policy name. It can contain 3 to 64 characters, starting with a letter.
	// Only letters, digits, and underscores (_) are allowed.
	Name string `json:"acl_name" required:"true"`
	// Type.
	// PERMIT (whitelist)
	// DENY (blacklist)
	Type string `json:"acl_type" required:"true"`
	// One or more objects from which the access will be controlled. Separate multiple objects with commas.
	// If entity_type is set to IP, enter up to 100 IP addresses.
	// If entity_type is set to DOMAIN, enter account names.
	// Each account name can contain up to 64 ASCII characters except commas (,).
	// Do not use only digits. The total length cannot exceed 1024 characters.
	// If entity_type is set to DOMAIN_ID, enter account IDs.
	Value string `json:"acl_value" required:"true"`
	// Object type.
	// IP: IP address.
	// DOMAIN: Account name.
	// DOMAIN_ID: Account ID.
	EntityType string `json:"entity_type" required:"true"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*AclResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("apigw", "instances", opts.GatewayID, "acls"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res AclResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type AclResp struct {
	// Name
	Name string `json:"acl_name"`
	// Type.
	// PERMIT (whitelist)
	// DENY (blacklist)
	Type string `json:"acl_type"`
	// Access control objects.
	Value string `json:"acl_value"`
	// Object type.
	// IP
	// DOMAIN
	// DOMAIN_ID
	EntityType string `json:"entity_type"`
	// ID.
	ID string `json:"id"`
	// Update time.
	UpdateTime string `json:"update_time"`
	// Number of bounded APIs.
	BindNum int `json:"bind_num"`
}
