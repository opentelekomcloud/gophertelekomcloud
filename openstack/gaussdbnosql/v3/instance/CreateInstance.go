package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateInstanceOpts struct {
	// Instance name, which can be the same as an existing instance name.
	// Value range:
	// The value must be 4 to 64 characters in length and start with a letter (from A to Z or from a to z).
	// It is case-sensitive and can contain only letters, digits (from 0 to 9), hyphens (-), and underscores (_).
	Name string `json:"name"`
	// Database information.
	Datastore Datastore `json:"datastore"`
	// Region ID
	// The value cannot be empty.
	Region string `json:"region"`
	// AZ ID
	// For details about the value, see az_status returned in 5.3 Querying All Instance Specifications.
	// Separate multiple AZs by commas (,).
	AvailabilityZone string `json:"availability_zone"`
	// VPC ID. To obtain this parameter value, use either of the following methods:
	// Method 1: Log in to VPC console and view the VPC ID on the VPC details page.
	// Method 2: See the "Querying VPCs" section in the Virtual Private Cloud API Reference.
	VpcId string `json:"vpc_id"`
	// Network ID of the subnet. To obtain this parameter value, use either of the following methods:
	// Method 1: Log in to VPC console and click the target subnet on the Subnets page. You can view the subnet ID on the displayed page.
	// Method 2: See the "Querying Subnets" section in the Virtual Private Cloud API Reference.
	SubnetId string `json:"subnet_id"`
	// Security group ID. To obtain the security group ID, perform either of the following methods:
	// Method 1: Log in to VPC console. Choose Access Control > Security Groups in the navigation pane on the left.
	// On the displayed page, click the target security group. You can view the security group ID on the displayed page.
	// Method 2: See the "Querying Security Groups" section in the Virtual Private Cloud API Reference.
	SecurityGroupId string `json:"security_group_id"`
	// Database password
	// The value must be 8 to 32 characters in length and contain uppercase letters (A to Z),
	// lowercase letters (a to z), digits (0 to 9), and special characters, such as ~!@#%^*-_=+?
	// You are advised to enter a strong password to improve security, preventing security risks such as brute force cracking.
	Password string `json:"password"`
	// Instance type
	// GaussDB(for Cassandra) supports the cluster type. The value is "Cluster".
	Mode string `json:"mode"`
	// Instance specifications
	Flavor []InstanceFlavor `json:"flavor"`
	// Parameter template ID
	ConfigurationId string `json:"configuration_id,omitempty"`
	// Advanced backup policy.
	BackupStrategy BackupStrategy `json:"backup_strategy,omitempty"`
	// Specifies whether to enable or disable SSL.
	// Valid value:
	// The value "0" indicates that SSL is disabled by default.
	// The value "1" indicates that SSL is enabled by default.
	// If this parameter is not transferred, SSL is disabled by default.
	SslOption string `json:"ssl_option,omitempty"`
}

func CreateInstance(client *golangsdk.ServiceClient, opts CreateInstanceOpts) (*CreateInstanceResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/v3/{project_id}/instances
	raw, err := client.Post(client.ServiceURL("instances"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res CreateInstanceResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type CreateInstanceResponse struct {
	// Instance ID
	Id string `json:"id"`
	// Database information, which is the same as the request parameter.
	Datastore Datastore `json:"datastore"`
	// Instance name, which is the same as the request parameter.
	Name string `json:"name"`
	// Creation time, which is in the yyyy-mm-dd hh:mm:ss format.
	Created string `json:"created"`
	// Instance status. The value is creating.
	Status string `json:"status"`
	// Region ID, which is the same as the request parameter.
	Region string `json:"region"`
	// AZ ID, which is the same as the request parameter.
	AvailabilityZone string `json:"availability_zone"`
	// VPC ID, which is the same as the request parameter.
	VpcId string `json:"vpc_id"`
	// Network ID of the subnet, which is the same as the request parameter.
	SubnetId string `json:"subnet_id"`
	// Security group ID, which is the same as the request parameter.
	SecurityGroupId string `json:"security_group_id"`
	// Instance type, which is the same as the request parameter.
	Mode string `json:"mode"`
	// Instance specification, which is the same as the request parameter.
	Flavor []InstanceFlavor `json:"flavor"`
	// Advanced backup policy, which is the same as the request parameter.
	BackupStrategy BackupStrategy `json:"backup_strategy"`
	// Indicates whether to enable SSL, which functions the same as the request parameter.
	SslOption string `json:"ssl_option"`
	// ID of the job for creating an instance.
	JobId string `json:"job_id"`
}

type Datastore struct {
	// Database type
	// GaussDB(for Cassandra) instances are supported.
	// The value "cassandra" indicates that a GaussDB(for Cassandra) DB instance is created.
	Type string `json:"type"`
	// Database version
	// The value "3.11" indicates that the GaussDB(for Cassandra) 3.11 is supported.
	Version string `json:"version"`
	// Storage engine
	// The value "rocksDB" indicates that the GaussDB(for Cassandra) instance supports the RocksDB storage engine.
	StorageEngine string `json:"storage_engine"`
}

type InstanceFlavor struct {
	// Node quantity
	// The number of GaussDB(for Cassandra) instance nodes ranges from 3 to 200.
	Num string `json:"num"`
	// Disk type
	// Valid value: ULTRAHIGH, which indicates the SSD disk.
	Storage string `json:"storage"`
	// Storage space. The value must be an integer, in GB.
	Size string `json:"size"`
	// Resource specification code
	SpecCode string `json:"spec_code"`
}

type BackupStrategy struct {
	// Backup time window Automated backups will be triggered during the backup time window.
	// The value cannot be empty. It must be a valid value in the "hh:mm-HH:MM" format.
	// The current time is in the UTC format.
	StartTime string `json:"start_time"`
	// Number of days to retain the generated backup files.
	// Value range: 0-35.
	KeepDays string `json:"keep_days,omitempty"`
}
