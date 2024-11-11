package keys

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type KeyRotationOpts struct {
	// ID of a CMK
	KeyID string `json:"key_id" required:"true"`
	// Rotation interval of a CMK
	Interval int `json:"rotation_interval"`
	// 36-byte serial number of a request message
	Sequence string `json:"sequence,omitempty"`
}

func EnableKeyRotation(client *golangsdk.ServiceClient, opts KeyRotationOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("kms", "enable-key-rotation"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
