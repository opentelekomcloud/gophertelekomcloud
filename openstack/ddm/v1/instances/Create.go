package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {

	// Instance Information. Value is of the type CreateInstanceDetail
	Instance CreateInstanceDetail `json:"instance" required:"true"`
}

type CreateInstanceDetail struct {
	// Instance name. The value can contain 4 to 64 characters, including letters, digits and hyphens (-).
	// Must start with a letter
	Name string `json:"name" required:"true"`
	// Specifies the class ID
	FlavorId string `json:"flavor_id" required:"true"`
	// Specifies the number of nodes
	NodeNum int `json:"node_num" required:"true"`
	// Specifies the Engine ID
	EngineId string `json:"engine_id" required:"true"`
	// Specifies the AZ codes. The value cannot be empty.
	AvailableZones []string `json:"available_zones" required:"true"`
	// Specifies the VPC ID
	VpcId string `json:"vpc_id" required:"true"`
	// Specifies the Security group ID
	SecurityGroupId string `json:"security_group_id" required:"true"`
	// Specifies the Subnet ID
	SubnetId string `json:"subnet_id" required:"true"`
	// Specifies the Parameter group ID
	ParamGroupId string `json:"param_group_id,omitempty"`
	// Time Zone. Default value is UTC. The value can be UTC, UTC-12:00, UTC-11:00, ..., UTC+12:00
	TimeZone string `json:"time_zone,omitempty"`
	// Username of the administrator. The username can contain 1 to 32 characters. It must start with a letter.
	// It can only contain letters, digits, and underscores (_)
	AdminUserName string `json:"admin_user_name,omitempty"`
	// Password of the administrator. The password can include 8 to 32 characters. It must be a combination of
	// uppercase letters, lowercase letters, digits, and the following special characters:
	// ~ ! @ # % ^ * - _ = + ?
	AdminUserPassword string `json:"admin_user_password,omitempty"`
}

// This function is used to create a DDM instance
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Instance, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/instances
	raw, err := client.Post(client.ServiceURL("instances"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Instance
	return &res, extract.Into(raw.Body, &res)
}

type Instance struct {
	// Indicates the DDM instance ID.
	Id string `json:"id"`
	// ID of the order
	OrderId string `json:"order_id"`
	// ID of the job for creating an instance.
	JobId string `json:"job_id"`
}
