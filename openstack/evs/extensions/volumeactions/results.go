package volumeactions

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// AttachResult contains the response body and error from an Attach request.
type AttachResult struct {
	golangsdk.ErrResult
}

// BeginDetachingResult contains the response body and error from a BeginDetach
// request.
type BeginDetachingResult struct {
	golangsdk.ErrResult
}

// DetachResult contains the response body and error from a Detach request.
type DetachResult struct {
	golangsdk.ErrResult
}

// UploadImageResult contains the response body and error from an UploadImage
// request.
type UploadImageResult struct {
	golangsdk.Result
}

// SetImageMetadataResult contains the response body and error from an SetImageMetadata
// request.
type SetImageMetadataResult struct {
	golangsdk.ErrResult
}

// SetBootableResult contains the response body and error from a SetBootable
// request.
type SetBootableResult struct {
	golangsdk.ErrResult
}

// ReserveResult contains the response body and error from a Reserve request.
type ReserveResult struct {
	golangsdk.ErrResult
}

// UnreserveResult contains the response body and error from an Unreserve
// request.
type UnreserveResult struct {
	golangsdk.ErrResult
}

// TerminateConnectionResult contains the response body and error from a
// TerminateConnection request.
type TerminateConnectionResult struct {
	golangsdk.ErrResult
}

// InitializeConnectionResult contains the response body and error from an
// InitializeConnection request.
type InitializeConnectionResult struct {
	golangsdk.Result
}

// ExtendSizeResult contains the response body and error from an ExtendSize request.
type ExtendSizeResult struct {
	golangsdk.ErrResult
}

// Extract will get the connection information out of the
// InitializeConnectionResult object.
//
// This will be a generic map[string]interface{} and the results will be
// dependent on the type of connection made.
func (r InitializeConnectionResult) Extract() (map[string]interface{}, error) {
	var s struct {
		ConnectionInfo map[string]interface{} `json:"connection_info"`
	}
	err := r.ExtractInto(&s)
	return s.ConnectionInfo, err
}

// Extract will get an object with info about the uploaded image out of the
// UploadImageResult object.
func (r UploadImageResult) Extract() (VolumeImage, error) {
	var s struct {
		VolumeImage VolumeImage `json:"os-volume_upload_image"`
	}
	err := r.ExtractInto(&s)
	return s.VolumeImage, err
}

// ForceDeleteResult contains the response body and error from a ForceDelete request.
type ForceDeleteResult struct {
	golangsdk.ErrResult
}

// ChangeTypeResult contains the response body and error from an ChangeType request.
type ChangeTypeResult struct {
	golangsdk.ErrResult
}

// ReImageResult contains the response body and error from a ReImage request.
type ReImageResult struct {
	golangsdk.ErrResult
}
