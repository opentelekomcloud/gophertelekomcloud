package vaults

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BindPolicyOpts struct {
	PolicyID string `json:"policy_id"`
}

func BindPolicy(client *golangsdk.ServiceClient, vaultID string, opts BindPolicyOpts) (*PolicyBinding, error) {
	b, err := build.RequestBodyMap(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("vaults", vaultID, "associatepolicy"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res PolicyBinding
	return &res, extract.IntoStructPtr(raw.Body, &res, "associate_policy")
}

type PolicyBinding struct {
	VaultID  string `json:"vault_id"`
	PolicyID string `json:"policy_id"`
}
