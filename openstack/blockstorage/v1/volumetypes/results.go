package volumetypes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// VolumeType contains all information associated with an OpenStack Volume Type.
type VolumeType struct {
	ExtraSpecs map[string]interface{} `json:"extra_specs"` // user-defined metadata
	ID         string                 `json:"id"`          // unique identifier
	Name       string                 `json:"name"`        // display name
}

// CreateResult contains the response body and error from a Create request.
type CreateResult struct {
	commonResult
}

// GetResult contains the response body and error from a Get request.
type GetResult struct {
	commonResult
}

// DeleteResult contains the response error from a Delete request.
type DeleteResult struct {
	golangsdk.ErrResult
}

// VolumeTypePage is a pagination.Pager that is returned from a call to the List function.
type VolumeTypePage struct {
	pagination.SinglePageBase
}

// IsEmpty returns true if a VolumeTypePage contains no Volume Types.
func (r VolumeTypePage) IsEmpty() (bool, error) {
	volumeTypes, err := ExtractVolumeTypes(r)
	return len(volumeTypes) == 0, err
}

// ExtractVolumeTypes extracts and returns Volume Types.
func ExtractVolumeTypes(r pagination.Page) ([]VolumeType, error) {
	var res struct {
		VolumeTypes []VolumeType `json:"volume_types"`
	}
	err := (r.(VolumeTypePage)).ExtractInto(&res)
	return res.VolumeTypes, err
}

type commonResult struct {
	golangsdk.Result
}

// Extract will get the Volume Type object out of the commonResult object.
func (r commonResult) Extract() (*VolumeType, error) {
	var res struct {
		VolumeType *VolumeType `json:"volume_type"`
	}
	err = extract.Into(raw.Body, &res)
	return res.VolumeType, err
}
