package vpc_endpoint

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateWhitelistOpts struct {
	// List of Tenant IDs to whitelist
	Permissions []string `json:"vpcPermissions" required:"true"`
}

// Modifying the VPC Endpoint Service Whitelist
func UpdateWhitelist(client *golangsdk.ServiceClient, clusterID string, opts UpdateWhitelistOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// POST /v1.0/{project_id}/clusters/{cluster_id}/vpcepservice/permissions
	url := client.ServiceURL("clusters", clusterID, "vpcepservice", "permissions")

	_, err = client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}
