package vaults

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type DissociateResourcesOpts struct {
	ResourceIDs []string `json:"resource_ids"`
}

func DissociateResources(client *golangsdk.ServiceClient, vaultID string, opts DissociateResourcesOpts) ([]string, error) {
	reqBody, err := build.RequestBodyMap(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("vaults", vaultID, "removeresources"), reqBody, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res []string
	return res, extract.IntoSlicePtr(raw.Body, &res, "remove_resource_ids")
}
