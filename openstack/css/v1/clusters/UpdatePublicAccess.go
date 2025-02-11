package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UpdatePublicAccess function is used to enable public network access
func UpdatePublicAccess(client *golangsdk.ServiceClient, opts ManagePublicAccessOpts) error {
	rawOpts := Eip{
		Bandwidth: Bandwidth{
			Size: opts.Size,
		},
	}

	b, err := build.RequestBody(rawOpts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("clusters", opts.ClusterId, "public", "bandwidth"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
