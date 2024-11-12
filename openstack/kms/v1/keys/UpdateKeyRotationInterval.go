package keys

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

func UpdateKeyRotationInterval(client *golangsdk.ServiceClient, opts RotationOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("kms", "update-key-rotation-interval"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
