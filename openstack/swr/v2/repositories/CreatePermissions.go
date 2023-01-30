package repositories

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/swr/v2/organizations"
)

func CreatePermissions(client *golangsdk.ServiceClient, organization, repository string, opts []organizations.Auth) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// POST /v2/manage/namespaces/{namespace}/repos/{repository}/access
	_, err = client.Post(client.ServiceURL("manage", "namespaces", organization, "repos", repository, "access"), b, nil, nil)
	return
}
