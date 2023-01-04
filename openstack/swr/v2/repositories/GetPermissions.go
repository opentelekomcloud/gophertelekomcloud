package repositories

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/swr/v2/organizations"
)

func GetPermissions(client *golangsdk.ServiceClient, organization, repository string) (*organizations.Permissions, error) {
	// GET /v2/manage/namespaces/{namespace}/repos/{repository}/access
	raw, err := client.Get(client.ServiceURL("manage", "namespaces", organization, "repos", repository, "access"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res organizations.Permissions
	err = extract.Into(raw.Body, &res)
	return &res, err
}
