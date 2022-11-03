package volumetypes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

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

// ExtractVolumeTypes extracts and returns Volumes. It is used while iterating over a volumetypes.List call.
func ExtractVolumeTypes(r pagination.Page) ([]VolumeType, error) {
	var s []VolumeType
	err := ExtractVolumeTypesInto(r, &s)
	return s, err
}

type commonResult struct {
	golangsdk.Result
}

// Extract will get the Volume Type object out of the commonResult object.
func (r commonResult) Extract() (*VolumeType, error) {
	var s VolumeType
	err := r.ExtractInto(&s)
	return &s, err
}

// ExtractInto converts our response data into a volume type struct
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "volume_type")
}

// GetResult contains the response body and error from a Get request.
type GetResult struct {
	commonResult
}

// CreateResult contains the response body and error from a Create request.
type CreateResult struct {
	commonResult
}

// DeleteResult contains the response body and error from a Delete request.
type DeleteResult struct {
	golangsdk.ErrResult
}

// UpdateResult contains the response body and error from an Update request.
type UpdateResult struct {
	commonResult
}

// extraSpecsResult contains the result of a call for (potentially) multiple
// key-value pairs. Call its Extract method to interpret it as a
// map[string]interface.
type extraSpecsResult struct {
	golangsdk.Result
}

// ListExtraSpecsResult contains the result of a Get operation. Call its Extract
// method to interpret it as a map[string]interface.
type ListExtraSpecsResult struct {
	extraSpecsResult
}

// CreateExtraSpecsResult contains the result of a Create operation. Call its
// Extract method to interpret it as a map[string]interface.
type CreateExtraSpecsResult struct {
	extraSpecsResult
}

// Extract interprets any extraSpecsResult as ExtraSpecs, if possible.
func (r extraSpecsResult) Extract() (map[string]string, error) {
	var s struct {
		ExtraSpecs map[string]string `json:"extra_specs"`
	}
	err := r.ExtractInto(&s)
	return s.ExtraSpecs, err
}

// extraSpecResult contains the result of a call for individual a single
// key-value pair.
type extraSpecResult struct {
	golangsdk.Result
}

// GetExtraSpecResult contains the result of a Get operation. Call its Extract
// method to interpret it as a map[string]interface.
type GetExtraSpecResult struct {
	extraSpecResult
}

// UpdateExtraSpecResult contains the result of an Update operation. Call its
// Extract method to interpret it as a map[string]interface.
type UpdateExtraSpecResult struct {
	extraSpecResult
}

// DeleteExtraSpecResult contains the result of a Delete operation. Call its
// ExtractErr method to determine if the call succeeded or failed.
type DeleteExtraSpecResult struct {
	golangsdk.ErrResult
}

// Extract interprets any extraSpecResult as an ExtraSpec, if possible.
func (r extraSpecResult) Extract() (map[string]string, error) {
	var s map[string]string
	err := r.ExtractInto(&s)
	return s, err
}

// VolumeTypeAccess represents an ACL of project access to a specific Volume Type.
type VolumeTypeAccess struct {
	// VolumeTypeID is the unique ID of the volume type.
	VolumeTypeID string `json:"volume_type_id"`

	// ProjectID is the unique ID of the project.
	ProjectID string `json:"project_id"`
}

// AccessPage contains a single page of all VolumeTypeAccess entries for a volume type.
type AccessPage struct {
	pagination.SinglePageBase
}

// IsEmpty indicates whether an AccessPage is empty.
func (page AccessPage) IsEmpty() (bool, error) {
	v, err := ExtractAccesses(page)
	return len(v) == 0, err
}

// ExtractAccesses interprets a page of results as a slice of VolumeTypeAccess.
func ExtractAccesses(r pagination.Page) ([]VolumeTypeAccess, error) {
	var s struct {
		VolumeTypeAccesses []VolumeTypeAccess `json:"volume_type_access"`
	}
	err := (r.(AccessPage)).ExtractInto(&s)
	return s.VolumeTypeAccesses, err
}

// AddAccessResult is the response from a AddAccess request. Call its
// ExtractErr method to determine if the request succeeded or failed.
type AddAccessResult struct {
	golangsdk.ErrResult
}

// RemoveAccessResult is the response from a RemoveAccess request. Call its
// ExtractErr method to determine if the request succeeded or failed.
type RemoveAccessResult struct {
	golangsdk.ErrResult
}
