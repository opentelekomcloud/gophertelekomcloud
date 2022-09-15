package snapshots

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToSnapshotCreateMap() (map[string]interface{}, error)
}

// PolicyCreateOpts contains options for creating a snapshot policy.
// This object is passed to the snapshots.PolicyCreate function.
type PolicyCreateOpts struct {
	Prefix     string `json:"prefix" required:"true"`
	Period     string `json:"period" required:"true"`
	KeepDay    int    `json:"keepday" required:"true"`
	Enable     string `json:"enable" required:"true"`
	DeleteAuto string `json:"deleteAuto,omitempty"`
}

// ToSnapshotCreateMap assembles a request body based on the contents of a
// PolicyCreateOpts.
func (opts PolicyCreateOpts) ToSnapshotCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// CreateOpts contains options for creating a snapshot.
// This object is passed to the snapshots.Create function.
type CreateOpts struct {
	Name        string `json:"name" required:"true"`
	Description string `json:"description,omitempty"`
	Indices     string `json:"indices,omitempty"`
}

// ToSnapshotCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToSnapshotCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

type UpdateConfigurationOptsBuilder interface {
	ToUpdateConfigurationMap() (map[string]interface{}, error)
}

type UpdateConfigurationOpts struct {
	// OBS bucket used for index data backup.
	// If there is snapshot data in an OBS bucket, only the OBS bucket is used and cannot be changed.
	Bucket string `json:"bucket" required:"true"`
	// IAM agency used to access OBS.
	Agency string `json:"agency" required:"true"`
	// Key ID used for snapshot encryption.
	SnapshotCmkID string `json:"snapshotCmkId,omitempty"`
}

func (opts UpdateConfigurationOpts) ToUpdateConfigurationMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}
