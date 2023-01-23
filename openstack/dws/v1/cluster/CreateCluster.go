package cluster

import (
	"fmt"
	"net/http"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateClusterOpts struct {
	// Node type
	NodeType string `json:"node_type" required:"true"`
	// Number of cluster nodes. For a cluster, the value ranges from 3 to 256. For a hybrid data warehouse (standalone), the value is 1.
	NumberOfNode int `json:"number_of_node" required:"true"`
	// Subnet ID, which is used for configuring cluster network.
	SubnetId string `json:"subnet_id" required:"true"`
	// ID of a security group, which is used for configuring cluster network.
	SecurityGroupId string `json:"security_group_id" required:"true"`
	// VPC ID, which is used for configuring cluster network.
	VpcId string `json:"vpc_id" required:"true"`
	// AZ of a cluster.
	AvailabilityZone string `json:"availability_zone,omitempty"`
	// Service port of a cluster. The value ranges from 8000 to 30000. The default value is 8000.
	Port int `json:"port,omitempty"`
	// Cluster name, which must be unique. The cluster name must contain 4 to 64 characters, which must start with a letter.
	// Only letters, digits, hyphens (-), and underscores (_) are allowed.
	Name string `json:"name" required:"true"`
	// Administrator username for logging in to a GaussDB(DWS) cluster. The username must:
	// Consist of lowercase letters, digits, or underscores.
	// Start with a lowercase letter or an underscore.
	// Contain 1 to 63 characters.
	// Cannot be a keyword of the GaussDB(DWS) database.
	UserName string `json:"user_name" required:"true"`
	// Administrator password for logging in to a GaussDB(DWS) cluster
	UserPwd string `json:"user_pwd" required:"true"`
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
	PublicBindType string `json:"public_bind_type" required:"true"`
	// EIP ID
	EipId string `json:"eip_id,omitempty"`
}

// CreateCluster is an asynchronous API. It takes 10 to 15 minutes to create a cluster.
func CreateCluster(client *golangsdk.ServiceClient, opts CreateClusterOpts) (string, error) {
	b, err := build.RequestBody(opts, "cluster")
	if err != nil {
		return "", err
	}

	// POST /v1.0/{project_id}/clusters
	raw, err := client.Post(client.ServiceURL("clusters"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return ExtractClusterId(err, raw)
}

func ExtractClusterId(err error, raw *http.Response) (string, error) {
	if err != nil {
		return "", err
	}

	var res struct {
		// Cluster ID
		Id string `json:"id"`
	}
	err = extract.IntoStructPtr(raw.Body, &res, "cluster")
	return res.Id, err
}

func WaitForCreate(c *golangsdk.ServiceClient, id string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		current, err := ListClusterDetails(c, id)
		if err != nil {
			return false, err
		}

		if current.Status == "AVAILABLE" {
			return true, nil
		}

		if current.Status == "CREATION FAILED" {
			return false, fmt.Errorf("cluster creation failed: " + current.FailedReasons.ErrorMsg)
		}

		time.Sleep(10 * time.Second)

		return false, nil
	})
}
