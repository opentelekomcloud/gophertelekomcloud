package cluster

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateClusterOpts struct {
	// Node type
	NodeType string `json:"node_type"`
	// Number of nodes in a cluster. The value ranges from 2 to 256.
	NumberOfNode int `json:"number_of_node"`
	// Subnet ID, which is used for configuring cluster network.
	SubnetId string `json:"subnet_id"`
	// ID of a security group, which is used for configuring cluster network.
	SecurityGroupId string `json:"security_group_id"`
	// VPC ID, which is used for configuring cluster network.
	VpcId string `json:"vpc_id"`
	// AZ of a cluster.
	AvailabilityZone string `json:"availability_zone,omitempty"`
	// Service port of a cluster. The value ranges from 8000 to 30000. The default value is 8000.
	Port int `json:"port,omitempty"`
	// Cluster name, which must be unique. The cluster name must contain 4 to 64 characters, which must start with a letter.
	// Only letters, digits, hyphens (-), and underscores (_) are allowed.
	Name string `json:"name"`
	// Administrator username for logging in to a GaussDB(DWS) cluster. The username must:
	// Consist of lowercase letters, digits, or underscores.
	// Start with a lowercase letter or an underscore.
	// Contain 1 to 63 characters.
	// Cannot be a keyword of the GaussDB(DWS) database.
	UserName string `json:"user_name"`
	// Administrator password for logging in to a GaussDB(DWS) cluster
	UserPwd string `json:"user_pwd"`
	// Public IP address. If the parameter is not specified, public connection is not used by default.
	PublicIp PublicIp `json:"public_ip,omitempty"`
	// Number of deployed CNs. The value ranges from 2 to the number of cluster nodes minus 1. The maximum value is 20 and the default value is 3.
	NumberOfCn int `json:"number_of_cn,omitempty"`
	// Enterprise project. The default enterprise project ID is 0.
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
}

type PublicIp struct {
	// Binding type of EIP. The value can be one of the following:
	// auto_assign
	// not_use
	// bind_existing
	PublicBindType string `json:"public_bind_type"`
	// EIP ID
	EipId string `json:"eip_id,omitempty"`
}

func CreateCluster(client *golangsdk.ServiceClient, opts CreateClusterOpts) (string, error) {
	b, err := build.RequestBody(opts, "cluster")
	if err != nil {
		return "", err
	}

	// POST /v1.0/{project_id}/clusters
	raw, err := client.Post(client.ServiceURL("clusters"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return "", err
	}

	var res Cluster
	err = extract.IntoStructPtr(raw.Body, &res, "cluster")
	return res.Id, err
}

type Cluster struct {
	// Cluster ID
	Id string `json:"id"`
}
