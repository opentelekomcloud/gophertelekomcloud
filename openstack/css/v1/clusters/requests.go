package clusters

import (
	"fmt"
	"log"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type CreateOptsBuilder interface {
	ToClusterCreateMap() (map[string]interface{}, error)
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

type InstanceSpec struct {
	// Flavor - instance flavor name.
	Flavor string `json:"flavorRef" required:"true"`
	// Volume - information about the volume.
	Volume *Volume `json:"volume" required:"true"`
	// Nics - subnet information.
	Nics             *Nics  `json:"nics" required:"true"`
	AvailabilityZone string `json:"availability_zone,omitempty"`
}

type Datastore struct {
	// Version - engine version.
	// The default value is 7.6.2.
	Version string `json:"version" required:"true"`
	// Type - Engine type.
	// The default value is `elasticsearch`.
	Type string `json:"type,omitempty"`
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

func (opts CreateOpts) ToClusterCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "cluster")
}

func WaitForClusterOperationSucces(client *golangsdk.ServiceClient, id string, timeout int) error {
	return golangsdk.WaitFor(timeout, func() (bool, error) {
		cluster, err := Get(client, id).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.BaseError); ok {
				return true, err
			}
			log.Printf("Error waiting for CSS cluster: %s", err) // ignore connection-related errors
			return false, nil
		}

		switch s := cluster.Status; s {
		case "100":
			time.Sleep(30 * time.Second) // make a bigger wait if it's not ready
			return false, nil
		case "200":
			return true, nil
		case "303":
			return true, fmt.Errorf("cluster operartion failed: %+v", cluster.FailedReasons)
		default:
			return true, fmt.Errorf("invalid status: %s", s)
		}
	})
}

type ClusterExtendOptsBuilder interface {
	ToExtendClusterMap() (map[string]interface{}, error)
}

// ClusterExtendCommonOpts is used to extend cluster with only `common` nodes.
// Clusters with master, client, or cold data nodes cannot use this.
type ClusterExtendCommonOpts struct {
	// ModifySize - number of instances to be added.
	ModifySize int `json:"modifySize"`
}

func (opts ClusterExtendCommonOpts) ToExtendClusterMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "grow")
}

// ClusterExtendSpecialOpts is used to extend cluster with special nodes.
// If a cluster has master, client, or cold data nodes, this should be used
type ClusterExtendSpecialOpts struct {
	// Type of the instance to be scaled out.
	// Select at least one from `ess`, `ess-cold`, `ess-master`, and `ess-client`.
	// You can only add instances, rather than increase storage capacity, on nodes of the `ess-master` and `ess-client` types.
	Type string `json:"type"`
	// NodeSize - number of instances to be scaled out.
	// The total number of existing instances and newly added instances in a cluster cannot exceed 32.
	NodeSize int `json:"nodesize"`
	// DiskSize - Storage capacity of the instance to be expanded.
	// The total storage capacity of existing instances and newly added instances in a cluster cannot exceed the maximum instance storage capacity allowed when a cluster is being created.
	// In addition, you can expand the instance storage capacity for a cluster for up to six times.
	// Unit: GB
	DiskSize int `json:"disksize"`
}

func (opts ClusterExtendSpecialOpts) ToExtendClusterMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "grow")
}

// ExtendCluster - extends cluster capacity.
// ClusterExtendCommonOpts should be used as options to extend cluster with only `common` nodes.
// ClusterExtendSpecialOpts should be used if extended cluster has master, client, or cold data nodes.
func ExtendCluster(client *golangsdk.ServiceClient, clusterID string, opts ClusterExtendOptsBuilder) (r ExtendResult) {
	url := ""
	switch opts.(type) {
	case ClusterExtendCommonOpts:
		url = client.ServiceURL("clusters", clusterID, "extend")
	case ClusterExtendSpecialOpts:
		url = client.ServiceURL("clusters", clusterID, "role_extend")
	default:
		r.Err = fmt.Errorf("invalid options type provided: %T", opts)
		return
	}

	b, err := opts.ToExtendClusterMap()
	if err != nil {
		return nil, err
	}
	raw, err = client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func WaitForClusterToExtend(client *golangsdk.ServiceClient, id string, timeout int) error {
	return golangsdk.WaitFor(timeout, func() (bool, error) {
		cluster, err := Get(client, id).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.BaseError); ok {
				return true, err
			}
			log.Printf("Error waiting for CSS cluster to extend: %s", err) // ignore connection-related errors
			return false, nil
		}
		// No active action
		if len(cluster.Actions) == 0 {
			return true, nil
		}
		if cluster.Actions[0] == "GROWING" {
			time.Sleep(30 * time.Second) // make a bigger wait if it's not ready
			return false, nil
		}
		return false, fmt.Errorf("unexpected cluster actions: %v; progress: %v", cluster.Actions, cluster.ActionProgress)
	})
}
