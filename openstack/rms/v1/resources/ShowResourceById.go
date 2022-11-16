package resources

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ShowResourceByIdOpts struct {
	// Specifies the cloud service name.
	// Maximum length: 20
	Provider string
	// Specifies the resource type.
	// Maximum length: 20
	Type string
	// Specifies the resource ID.
	// Maximum length: 256
	ResourceId string
}

func ShowResourceById(client *golangsdk.ServiceClient, opts ShowResourceByIdOpts) (*ResourceEntity, error) {
	// GET /v1/resource-manager/domains/{domain_id}/provider/{provider}/type/{type}/resources/{resource_id}
	raw, err := client.Get(client.ServiceURL("provider", opts.Provider, "type", opts.Type, "resources", opts.ResourceId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ResourceEntity
	err = extract.Into(raw.Body, &res)
	return &res, err
}
