package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateInstanceOpts struct {
	// Billing mode
	ChargeInfo *ChargeInfo `json:"charge_info,omitempty"`
	// Region ID The value cannot be empty. To obtain this value, see Regions and Endpoints.
	Region string `json:"region" required:"true"`
	// Instance name Instances of the same type can have same names under the same tenant.
	// The name consists of 4 to 64 characters and starts with a letter.
	// It is case-sensitive and can contain only letters, digits, hyphens (-), and underscores (_).
	Name string `json:"name" required:"true"`
	// Database information
	Datastore Datastore `json:"datastore" required:"true"`
	// Instance type, which is case-insensitive. Currently, its value can only be Cluster.
	Mode string `json:"mode" required:"true"`
	// Specification code. For details, see Querying Database Specifications.
	FlavorRef string `json:"flavor_ref" required:"true"`
	// VPC ID. To obtain this value, use either of the following methods:
	// Method 1: Log in to the VPC console and view the VPC ID on the VPC details page.
	// Method 2: See section "Querying VPCs" in the Virtual Private Cloud API Reference.
	VpcId string `json:"vpc_id" required:"true"`
	// Network ID of the subnet. To obtain this value, use either of the following methods:
	// Method 1: Log in to the VPC console and click the target subnet on the Subnets page to view the network ID on the displayed page.
	// Method 2: See section "Querying Subnets" in the Virtual Private Cloud API Reference.
	SubnetId string `json:"subnet_id" required:"true"`
	// Security group ID If the network ACL is enabled for the subnet used by the created instance, this parameter is optional.
	// If the network ACL is not enabled, this parameter is mandatory.
	// Method 1: Log in to VPC console. Choose Access Control > Security Groups in the navigation pane on the left.
	// On the displayed page, click the target security group. You can view the security group ID on the displayed page.
	// Method 2: See section "Querying Security Groups" in the Virtual Private Cloud API Reference.
	SecurityGroupId string `json:"security_group_id,omitempty"`
	// Parameter template ID
	ConfigurationId string `json:"configuration_id,omitempty"`
	// Database password. Value range: The password consists of 8 to 32 characters and contains at least three types of the following:
	// uppercase letters, lowercase letters, digits, and special characters (~!@#%^*-_=+?).
	// Enter a strong password to improve security, preventing security risks such as brute force cracking.
	// If you enter a weak password, the system automatically determines that the password is invalid.
	Password string `json:"password" required:"true"`
	// Automated backup policy
	BackupStrategy *BackupStrategy `json:"backup_strategy,omitempty"`
	// UTC time zone. l If this parameter is not specified, UTC is used by default.
	// If this parameter is specified, the value ranges from UTC-12:00 to UTC+12:00 at the full hour.
	// For example, the parameter can be UTC+08:00 rather than UTC+08:30.
	TimeZone string `json:"time_zone,omitempty"`
	// AZ type. The value can be Single or multi.
	AvailabilityZoneMode string `json:"availability_zone_mode"`
	// Primary AZ
	MasterAvailabilityZone string `json:"master_availability_zone,omitempty"`
	// Number of created read replicas. The value ranges from 1 to 9. An instance contains up to 15 read replicas.
	SlaveCount *int `json:"slave_count"`
	// Volume information.
	// Missing in documentation
	Volume *MysqlVolume `json:"volume,omitempty"`
	// Tag list. Instances are created based on tag keys and values.
	// {key} indicates the tag key. It must be unique and cannot be empty.
	// {value} indicates the tag value, which can be empty. To create instances with multiple tag keys and values,
	// separate key-value pairs with commas (,). Up to 10 key-value pairs can be added.
	Tags []MysqlTags `json:"tags,omitempty"`
	// Dedicated resource pool ID. This parameter can be displayed only after the dedicated resource pool is enabled.
	DedicatedResourceId string `json:"dedicated_resource_id,omitempty"`
}

type MysqlVolume struct {
	// Storage space. The default value is 40 in GB.
	// The value ranges from 40 GB to 128,000 GB and must be a multiple of 10.
	Size string `json:"size"`
}

type MysqlTags struct {
	// Tag key. The value can contain up to 36 unicode characters. The value cannot be an empty string, a space, or left blank.
	// Only uppercase letters, lowercase letters, digits, hyphens (-), and underscores (_) are allowed.
	Key string `json:"key"`
	// Tag value. It contains up to 43 Unicode characters. The value can be an empty string.
	// Only uppercase letters, lowercase letters, digits, periods (.), hyphens (-), and underscores (_) are allowed.
	Value string `json:"value"`
}

func CreateInstance(client *golangsdk.ServiceClient, opts CreateInstanceOpts) (*CreateInstanceResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/mysql/v3/{project_id}/instances
	raw, err := client.Post(client.ServiceURL("instances"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res CreateInstanceResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type CreateInstanceResponse struct {
	// Instance information
	Instance InstResponse `json:"instance"`
	// Instance creation task ID This parameter is returned only for the creation of pay-per-use instances.
	JobId string `json:"job_id"`
}

type InstResponse struct {
	// Instance ID
	Id string `json:"id"`
	// Instance name. Instances of the same type can have same names under the same tenant.
	// The name consists of 4 to 64 characters and starts with a letter.
	// It is case-insensitive and contains only letters, digits, hyphens (-), and underscores (_).
	Name string `json:"name"`
	// Instance status Value: creating.
	Status string `json:"status"`
	// Database information
	Datastore Datastore `json:"datastore"`
	// Instance type. Currently, only the cluster type is supported.
	Mode string `json:"mode"`
	// Parameter template ID.
	ConfigurationId string `json:"configuration_id"`
	// Database port, which is the same as the request parameter.
	Port string `json:"port"`
	// Automated backup policy
	BackupStrategy BackupStrategy `json:"backup_strategy"`
	// Region ID, which is the same as the request parameter.
	Region string `json:"region"`
	// AZ mode, which is the same as the request parameter.
	AvailabilityZoneMode string `json:"availability_zone_mode"`
	// Primary AZ ID.
	MasterAvailabilityZone string `json:"master_availability_zone"`
	// VPC ID, which is the same as the request parameter.
	VpcId string `json:"vpc_id"`
	// Security group ID, which is the same as the request parameter.
	SecurityGroupId string `json:"security_group_id"`
	// Subnet ID, which is the same as the request parameter.
	SubnetId string `json:"subnet_id"`
	// Specification code, which is the same as the request parameter.
	FlavorRef string `json:"flavor_ref"`
	// Billing mode, which is yearly/monthly or pay-per-use (default setting).
	ChargeInfo ChargeInfo `json:"charge_info"`
}

type Datastore struct {
	// DB engine. Currently, only gaussdb-mysql is supported.
	Type string `json:"type" required:"true"`
	// DB engine version For details about supported database versions, see Querying Version Information About a DB Engine.
	Version string `json:"version" required:"true"`
}

type BackupStrategy struct {
	// Automated backup start time. The automated backup will be triggered within one hour after the time specified by this parameter.
	// The value cannot be empty. It must be a valid value in the "hh:mm-HH:MM" format. The current time is in the UTC format.
	// The HH value must be 1 greater than the hh value.
	// The values of mm and MM must be the same and must be set to 00.
	// Example value: 21:00-22:00
	StartTime string `json:"start_time"`
	// Automated backup retention days. The value ranges from 1 to 732.
	KeepDays string `json:"keep_days,omitempty"`
}

type ChargeInfo struct {
	// Billing mode. Value: postPaid
	ChargeMode string `json:"charge_mode"`
	// Order ID
	OrderId string `json:"order_id"`
}
