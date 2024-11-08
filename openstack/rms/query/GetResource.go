package query

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetResource(client *golangsdk.ServiceClient, domainId, service, resourceType, id string) (*Resource, error) {
	// GET /v1/resource-manager/domains/{domain_id}/provider/{provider}/type/{type}/resources/{resource_id}
	raw, err := client.Get(client.ServiceURL(
		"resource-manager", "domains", domainId,
		"provider", service, "type", resourceType, "resources", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Resource
	err = extract.Into(raw.Body, &res)
	return &res, err
}
