package organizations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

func UpdatePermissions(client *golangsdk.ServiceClient, organization string, opts []Auth) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// PATCH /v2/manage/namespaces/{namespace}/access
	_, err = client.Patch(client.ServiceURL("manage", "namespaces", organization, "access"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	return
}
