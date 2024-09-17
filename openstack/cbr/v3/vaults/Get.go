package vaults

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*Vault, error) {
	raw, err := client.Get(client.ServiceURL("vaults", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Vault
	return &res, extract.IntoStructPtr(raw.Body, &res, "vault")
}
