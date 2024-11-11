package keys

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

func DisableKeyRotation(client *golangsdk.ServiceClient, opts KeyRotationOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("kms", "disable-key-rotation"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
