package volumeactions

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// ImageMetadataOptsBuilder allows extensions to add additional parameters to the
// ImageMetadataRequest request.
type ImageMetadataOptsBuilder interface {
	ToImageMetadataMap() (map[string]interface{}, error)
}

// ImageMetadataOpts contains options for setting image metadata to a volume.
type ImageMetadataOpts struct {
	// The image metadata to add to the volume as a set of metadata key and value pairs.
	Metadata map[string]string `json:"metadata"`
}

// ToImageMetadataMap assembles a request body based on the contents of a
// ImageMetadataOpts.
func (opts ImageMetadataOpts) ToImageMetadataMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "os-set_image_metadata")
}

// SetImageMetadata will set image metadata on a volume based on the values in ImageMetadataOptsBuilder.
func SetImageMetadata(client *golangsdk.ServiceClient, id string, opts ImageMetadataOptsBuilder) (r SetImageMetadataResult) {
	b, err := opts.ToImageMetadataMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(client.ServiceURL("volumes", id, "action"), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	_, r.Header, r.Err = golangsdk.ParseResponse(resp, err)
	return
}

// MigrationPolicy type represents a migration_policy when changing types.
type MigrationPolicy string

// Supported attributes for MigrationPolicy attribute for changeType operations.
const (
	MigrationPolicyNever    MigrationPolicy = "never"
	MigrationPolicyOnDemand MigrationPolicy = "on-demand"
)

// ChangeTypeOptsBuilder allows extensions to add additional parameters to the
// ChangeType request.
type ChangeTypeOptsBuilder interface {
	ToVolumeChangeTypeMap() (map[string]interface{}, error)
}

// ChangeTypeOpts contains options for changing the type of an existing Volume.
// This object is passed to the volumes.ChangeType function.
type ChangeTypeOpts struct {
	// NewType is the name of the new volume type of the volume.
	NewType string `json:"new_type" required:"true"`

	// MigrationPolicy specifies if the volume should be migrated when it is
	// re-typed. Possible values are "on-demand" or "never". If not specified,
	// the default is "never".
	MigrationPolicy MigrationPolicy `json:"migration_policy,omitempty"`
}

// ToVolumeChangeTypeMap assembles a request body based on the contents of an
// ChangeTypeOpts.
func (opts ChangeTypeOpts) ToVolumeChangeTypeMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "os-retype")
}

// ChangeType will change the volume type of the volume based on the provided information.
// This operation does not return a response body.
func ChangeType(client *golangsdk.ServiceClient, id string, opts ChangeTypeOptsBuilder) (r ChangeTypeResult) {
	b, err := opts.ToVolumeChangeTypeMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(client.ServiceURL("volumes", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	_, r.Header, r.Err = golangsdk.ParseResponse(resp, err)
	return
}

// ReImageOpts contains options for Re-image a volume.
type ReImageOpts struct {
	// New image id
	ImageID string `json:"image_id"`
	// set true to re-image volumes in reserved state
	ReImageReserved bool `json:"reimage_reserved"`
}

// ToReImageMap assembles a request body based on the contents of a ReImageOpts.
func (opts ReImageOpts) ToReImageMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "os-reimage")
}

// ReImage will re-image a volume based on the values in ReImageOpts
func ReImage(client *golangsdk.ServiceClient, id string, opts ReImageOpts) (r ReImageResult) {
	b, err := opts.ToReImageMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(client.ServiceURL("volumes", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	_, r.Header, r.Err = golangsdk.ParseResponse(resp, err)
	return
}
