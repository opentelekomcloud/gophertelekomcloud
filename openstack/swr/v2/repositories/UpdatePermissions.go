package repositories

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/swr/v2/organizations"
)

func UpdatePermissions(client *golangsdk.ServiceClient, organization, repository string, opts []organizations.Auth) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// PATCH /v2/manage/namespaces/{namespace}/repos/{repository}/access
	_, err = client.Patch(client.ServiceURL("manage", "namespaces", organization, "repos", repository, "access"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	return
}
