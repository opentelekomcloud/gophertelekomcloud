package vaults

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type BillingUpdate struct {
	Size int `json:"size,omitempty"`
}

type VaultBindRules struct {
	// Filters automatically associated resources by tag.
	Tags []tags.ResourceTag `json:"tags,omitempty"`
}

type UpdateOpts struct {
	Billing    *BillingUpdate  `json:"billing,omitempty"`
	Name       string          `json:"name,omitempty"`
	AutoBind   *bool           `json:"auto_bind,omitempty"`
	BindRules  *VaultBindRules `json:"bind_rules,omitempty"`
	AutoExpand *bool           `json:"auto_expand,omitempty"`
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*Vault, error) {
	b, err := build.RequestBody(opts, "vault")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("vaults", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Vault
	return &res, extract.IntoStructPtr(raw.Body, &res, "vault")
}
