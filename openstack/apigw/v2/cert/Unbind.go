package cert

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// Unbind SSL certificates from a domain
func Unbind(client *golangsdk.ServiceClient, opts BindOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// Build URL: /v2/{project_id}/apigw/instances/{instance_id}/api-groups/{group_id}/domains/{domain_id}/certificates/detach
	_, err = client.Post(client.ServiceURL("apigw", "instances", opts.InstanceID,
		"api-groups", opts.GroupID, "domains", opts.DomainID, "certificates", "detach"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 204},
	})
	return err
}
