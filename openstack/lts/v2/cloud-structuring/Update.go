package cloud_structuring

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

func Update(client *golangsdk.ServiceClient, opts CreateOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// PUT /v3/{project_id}/lts/struct/template
	_, err = client.Put(client.ServiceURL("lts", "struct", "template"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return err
}
