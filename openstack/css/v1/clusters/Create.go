package clusters

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type CreateOpts struct {
	// Instance - instance specification
	Instance *InstanceSpec `json:"instance" required:"true"`
	// Datastore - type of the data search engine
	Datastore *Datastore `json:"datastore,omitempty"`
	// Name - cluster name.
	// It contains 4 to 32 characters. Only letters, digits, hyphens (-), and underscores (_) are allowed.
	// The value must start with a letter.
	Name string `json:"name" required:"true"`
	// InstanceNum - number of clusters. The value range is 1 to 32.
	InstanceNum int `json:"instanceNum" required:"true"`
	// BackupStrategy - configuration of automatic snapshot creation.
	// This function is enabled by default.
	BackupStrategy *BackupStrategy `json:"backupStrategy,omitempty"`
	// DiskEncryption - disk encryption configuration
	DiskEncryption *DiskEncryption `json:"diskEncryption" required:"true"`
	// HttpsEnabled - whether communication is not encrypted on the cluster.
	HttpsEnabled string `json:"httpsEnable,omitempty"`
	// AuthorityEnabled - whether to enable authentication.
	// Available values include `true` and `false`. Authentication is disabled by default.
	// When authentication is enabled, `HttpsEnabled` must be set to `true`.
	AuthorityEnabled bool `json:"authorityEnable,omitempty"`
	// AdminPassword - password of the cluster user `admin` in security mode.
	// This parameter is mandatory only when `AuthorityEnabled` is set to `true`.
	AdminPassword string `json:"adminPwd,omitempty"`
	// Tags - tags of a cluster.
	Tags []tags.ResourceTag `json:"tags,omitempty"`
}

type InstanceSpec struct {
	// Flavor - instance flavor name.
	Flavor string `json:"flavorRef" required:"true"`
	// Volume - information about the volume.
	Volume *Volume `json:"volume" required:"true"`
	// Nics - subnet information.
	Nics             *Nics  `json:"nics" required:"true"`
	AvailabilityZone string `json:"availability_zone,omitempty"`
}

type Volume struct {
	// Type of the volume.
	// One of:
	//   - `COMMON`: Common I/O
	//   - `HIGH`: High I/O
	//   - `ULTRAHIGH`: Ultra-high I/O
	Type string `json:"volume_type" required:"true"`
	// Size of the volume, which must be a multiple of 4 and 10.
	// Unit: GB.
	Size int `json:"size" required:"true"`
}

type Nics struct {
	// VpcID - VPC ID which is used for configuring cluster network.
	VpcID string `json:"vpcId" required:"true"`
	// SubnetID - Subnet ID.
	// All instances in a cluster must have the same subnet.
	SubnetID string `json:"netId" required:"true"`
	// SecurityGroupID - Security group ID.
	// All instances in a cluster must have the same security grou.
	SecurityGroupID string `json:"securityGroupId" required:"true"`
}

type BackupStrategy struct {
	// Period - time when a snapshot is created every day.
	// Snapshots can only be created on the hour.
	// The time format is the time followed by the time zone, specifically, `HH:mm z`.
	// In the format, `HH:mm` refers to the hour time and `z` refers to the time zone, for example,
	// `00:00 GMT+08:00` and `01:00 GMT+08:00`.
	Period string `json:"period" required:"true"`
	// Prefix - prefix of the name of the snapshot that is automatically created.
	Prefix string `json:"prefix" required:"true"`
	// KeepDay - number of days for which automatically created snapshots are reserved.
	// Value range: `1` to `90`
	KeepDay int `json:"keepday" required:"true"`
}

type DiskEncryption struct {
	// SystemEncrypted - value `1` indicates encryption is performed, and value `0` indicates encryption is not performed.
	// Boolean sucks.
	Encrypted string `json:"systemEncrypted" required:"true"`
	// Key ID.
	//  - The Default Master Keys cannot be used to create grants.
	//    Specifically, you cannot use Default Master Keys whose aliases end with `/default` in KMS to create clusters.
	//  - After a cluster is created, do not delete the key used by the cluster. Otherwise, the cluster will become unavailable.
	CmkID string `json:"systemCmkid"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*CreatedCluster, error) {
	b, err := build.RequestBodyMap(opts, "cluster")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("clusters"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res CreatedCluster
	err = extract.IntoStructPtr(raw.Body, &res, "cluster")
	return &res, err
}

type CreatedCluster struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
