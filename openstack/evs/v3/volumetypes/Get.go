package volumetypes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get retrieves the Volume Type with the provided ID.
func Get(client *golangsdk.ServiceClient, id string) (*VolumeType, error) {
	raw, err := client.Get(client.ServiceURL("types", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res VolumeType
	err = extract.Into(raw.Body, &res)
	return &res, err
}

// VolumeType contains all the information associated with an OpenStack Volume Type.
type VolumeType struct {
	// Unique identifier for the volume type.
	ID string `json:"id"`
	// Human-readable display name for the volume type.
	Name string `json:"name"`
	// Human-readable description for the volume type.
	Description string `json:"description"`
	// Arbitrary key-value pairs defined by the user.
	ExtraSpecs map[string]string `json:"extra_specs"`
	// Whether the volume type is publicly visible.
	IsPublic bool `json:"is_public"`
	// Qos Spec ID
	QosSpecID string `json:"qos_specs_id"`
	// Volume Type access public attribute
	PublicAccess bool `json:"os-volume-type-access:is_public"`
}

type ExtraSpecs struct {
	// Reserved field
	VolumeBackendName string `json:"volume_backend_name"`
	// Reserved field
	AvailabilityZone string `json:"availability-zone"`
	// Reserved field
	HWAZ string `json:"HW:availability_zone"`
	// Specifies the AZs that support the current disk type.
	RESKEYAZ string `json:"RESKEY:availability_zones"`
}
