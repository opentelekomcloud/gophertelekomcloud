package shares

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// The availability zone of the SFS Turbo file system
	AvailabilityZone string `json:"availability_zone" required:"true"`
	// Share description
	Description string `json:"description,omitempty"`
	// The metadata information
	Metadata Metadata `json:"metadata,omitempty"`
	// Defines the SFS Turbo file system name
	Name string `json:"name" required:"true"`
	// The security group ID
	SecurityGroupID string `json:"security_group_id" required:"true"`
	// Defines the SFS Turbo file system protocol to use, the valid value is NFS.
	ShareProto string `json:"share_proto,omitempty"`
	// ShareType defines the file system type. the valid values are STANDARD and PERFORMANCE.
	ShareType string `json:"share_type" required:"true"`
	// Size in GB, range from 500 to 32768.
	Size int `json:"size" required:"true"`
	// The subnet ID
	SubnetID string `json:"subnet_id" required:"true"`
	// The backup ID
	BackupID string `json:"backup_id,omitempty"`
	// The VPC ID
	VpcID string `json:"vpc_id" required:"true"`
}

// Metadata specifies the metadata information
type Metadata struct {
	CryptKeyID string `json:"crypt_key_id,omitempty"`
	ExpandType string `json:"expand_type,omitempty"`
	HpcBW      string `json:"hpc_bw,omitempty"`
}

// Create will create a new SFS Turbo file system based on the values in CreateOpts. To extract
// the Share object from the response, call the Extract method on the
// CreateResult.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*TurboResponse, error) {
	b, err := build.RequestBody(opts, "share")
	if err != nil {
		return nil, err
	}
	// POST /v1/{project_id}/sfs-turbo/shares
	raw, err := client.Post(client.ServiceURL("sfs-turbo", "shares"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 202},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res TurboResponse
	return &res, extract.Into(raw.Body, &res)
}
