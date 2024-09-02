package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func QueryInstanceDetails(client *golangsdk.ServiceClient, instanceId string) (*QueryInstanceDetailsResponse, error) {

	// GET /v1/{project_id}/instances/{instance_id}
	raw, err := client.Get(client.ServiceURL("instances", instanceId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res QueryInstanceDetailsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type QueryInstanceDetailsResponse struct {
	// DDM instance ID
	Id string `json:"id"`
	// DDM instance status. For details about this parameter value, see Status Description at
	// https://docs.otc.t-systems.com/distributed-database-middleware/api-ref/appendix/status_description.html
	Status string `json:"status"`
	// DDM instance name
	Name string `json:"name"`
	// Time when the DDM instance is created
	Created string `json:"created"`
	// Time when the DDM instance is last updated
	Updated string `json:"updated"`
	// Name of the AZ where the DDM instance is located
	AvailableZone string `json:"available_zone"`
	// VPC ID
	VpcId string `json:"vpc_id"`
	// Subnet ID
	SubnetId string `json:"subnet_id"`
	// Security group ID
	SecurityGroupId string `json:"security_group_id"`
	// Number of nodes
	NodeCount int `json:"node_count"`
	// Address for accessing the DDM instance
	AccessIp string `json:"access_ip"`
	// Port for accessing the DDM instance
	AccessPort string `json:"access_port"`
	// Node status. For details, see Status Description.
	NodeStatus string `json:"node_status"`
	// Number of CPUs
	CoreCount string `json:"core_count"`
	// Memory size in GB
	RamCapacity string `json:"ram_capacity"`
	// Response message. This parameter is not returned if no abnormality occurs.
	ErrorMsg string `json:"error_msg,omitempty"`
	// Project ID
	ProjectId string `json:"project_id"`
	// The function has not been supported, and this field is reserved.
	OrderId string `json:"order_id"`
	// Engine version (core instance version)
	EngineVersion string `json:"engine_version"`
	// Node information
	Nodes []GetDetailfNodesInfo `json:"nodes"`
	// Username of the administrator. The username: Can include 1 to 32 characters. Must start with a letter. Can contain only letters, digits, and underscores (_).
	AdminUserName string `json:"admin_user_name"`
}

type GetDetailfNodesInfo struct {
	// Status of the DDM instance node. For details, see Status Description at
	// https://docs.otc.t-systems.com/distributed-database-middleware/api-ref/appendix/status_description.html
	Status string `json:"status"`
	// Port of the DDM instance node
	Port string `json:"port"`
	// IP address of the DDM instance node
	IP string `json:"ip"`
}
