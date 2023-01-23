package snapshot

import (
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dws/v1/cluster"
)

type RestoreClusterOpts struct {
	// ID of the snapshot to be restored
	SnapshotId string `json:"-"`
	// Cluster name, which must be unique. The cluster name must contain 4 to 64 characters, which must start with a letter.
	// Only letters, digits, hyphens (-), and underscores (_) are allowed.
	Name string `json:"name" required:"true"`
	// Subnet ID, which is used for configuring cluster network. The default value is the same as that of the original cluster.
	SubnetId string `json:"subnet_id,omitempty"`
	// Security group ID, which is used for configuring cluster network. The default value is the same as that of the original cluster.
	SecurityGroupId string `json:"security_group_id,omitempty"`
	// VPC ID, which is used for configuring cluster network. The default value is the same as that of the original cluster.
	VpcId string `json:"vpc_id,omitempty"`
	// AZ of a cluster. The default value is the same as that of the original cluster.
	AvailabilityZone string `json:"availability_zone,omitempty"`
	// Service port of a cluster. The value ranges from 8000 to 30000. The default value is 8000.
	Port int `json:"port,omitempty"`
	// Public IP address. If the parameter is not specified, public connection is not used by default.
	PublicIp cluster.PublicIp `json:"public_ip,omitempty"`
	// Enterprise project. The default enterprise project ID is 0.
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
}

func RestoreCluster(client *golangsdk.ServiceClient, opts RestoreClusterOpts) (string, error) {
	b, err := build.RequestBody(opts, "restore")
	if err != nil {
		return "", err
	}

	// POST /v1.0/{project_id}/snapshots/{snapshot_id}/actions
	raw, err := client.Post(client.ServiceURL("snapshots", opts.SnapshotId, "actions"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return cluster.ExtractClusterId(err, raw)
}

func WaitForRestore(c *golangsdk.ServiceClient, id string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		current, err := cluster.ListClusterDetails(c, id)
		if err != nil {
			return false, err
		}

		if current.Status == "AVAILABLE" && current.TaskStatus != "RESTORING" {
			return true, nil
		}

		time.Sleep(10 * time.Second)

		return false, nil
	})
}
