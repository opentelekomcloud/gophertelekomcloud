package keys

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type KeyRotationStatus struct {
	// Key rotation status. The default value is false, indicating that key rotation is disabled.
	Enabled bool `json:"key_rotation_enabled"`
	// Rotation interval. The value is an integer in the range 30 to 365.
	Interval int `json:"rotation_interval"`
	// Last key rotation time
	LastRotationTime string `json:"last_rotation_time"`
	// Number of key rotations
	NumberOfRotations int `json:"number_of_rotations"`
}

func GetKeyRotationStatus(client *golangsdk.ServiceClient, opts KeyRotationOpts) (*KeyRotationStatus, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("kms", "get-key-rotation-status"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res KeyRotationStatus
	err = extract.Into(raw.Body, &res)
	return &res, err
}
