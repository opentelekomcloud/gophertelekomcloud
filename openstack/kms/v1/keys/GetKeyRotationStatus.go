package keys

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type RotationOpts struct {
	// ID of a CMK
	KeyID string `json:"key_id" required:"true"`
	// Rotation interval of a CMK
	Interval int `json:"rotation_interval"`
	// 36-byte serial number of a request message
	Sequence string `json:"sequence,omitempty"`
}

func GetKeyRotationStatus(client *golangsdk.ServiceClient, opts RotationOpts) (*KeyRotationResult, error) {
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

	var res KeyRotationResult

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type KeyRotationResult struct {
	// Key rotation status. The default value is false, indicating that key rotation is disabled.
	Enabled bool `json:"key_rotation_enabled"`
	// Rotation interval. The value is an integer in the range 30 to 365.
	Interval int `json:"rotation_interval"`
	// Last key rotation time. The timestamp indicates the total microseconds past the start of the epoch date (January 1, 1970).
	LastRotationTime string `json:"last_rotation_time"`
	// Number of key rotations.
	NumberOfRotations int `json:"number_of_rotations"`
}
