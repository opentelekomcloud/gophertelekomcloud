package backups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/instances"
)

type RestoreToNewOpts struct {
	// Specifies the DB instance name.
	// DB instances of the same type can have same names under the same tenant.
	// The value must be 4 to 64 characters in length and start with a letter. It is case-insensitive and can contain only letters, digits, hyphens (-), and underscores (_).
	Name string `json:"name" required:"true"`
	// Specifies the HA configuration parameters, which are used when creating primary/standby DB instances.
	Ha *instances.Ha `json:"ha,omitempty"`
	// Specifies the parameter template ID.
	ConfigurationId string `json:"configuration_id,omitempty"`
	// Specifies the database port information.
	//
	// The MySQL database port ranges from 1024 to 65535 (excluding 12017 and 33071, which are occupied by the RDS system and cannot be used).
	// The PostgreSQL database port ranges from 2100 to 9500.
	// The Microsoft SQL Server database port is 1433 or ranges from 2100 to 9500 (excluding 5355 and 5985).
	// If this parameter is not set, the default value is as follows:
	//
	// For MySQL, the default value is 3306.
	// For PostgreSQL, the default value is 5432.
	// For Microsoft SQL Server, the default value is 1433.
	Port string `json:"port,omitempty"`
	// Specifies the database password.
	// Valid value:
	// The value cannot be empty and should contain 8 to 32 characters, including uppercase and lowercase letters, digits, and the following special characters: ~!@#%^*-_=+?
	// You are advised to enter a strong password to improve security, preventing security risks such as brute force cracking.
	// If provided password will be considered by system as weak, you will receive an error and you should provide stronger password.
	Password string `json:"password" required:"true"`
	// Specifies the advanced backup policy.
	BackupStrategy *instances.BackupStrategy `json:"backup_strategy,omitempty"`
	// Specifies the key ID for disk encryption. The default value is empty.
	DiskEncryptionId string `json:"disk_encryption_id,omitempty"`
	// Specifies the specification code. The value cannot be empty.
	FlavorRef string `json:"flavor_ref" required:"true"`
	// Specifies the volume information.
	Volume *instances.Volume `json:"volume" required:"true"`
	// Specifies the AZ ID. If the DB instance is not a single instance, you need to specify an AZ for each node of the instance and separate the AZs with commas (,). For details, see the example.
	// The value cannot be empty.
	AvailabilityZone string `json:"availability_zone" required:"true"`
	// Specifies the VPC ID. To obtain this parameter value, use either of the following methods:
	//
	// Method 1: Log in to VPC console and view the VPC ID in the VPC details.
	// Method 2: See the "Querying VPCs" section in the Virtual Private Cloud API Reference.
	VpcId string `json:"vpc_id" required:"true"`
	// Specifies the network ID. To obtain this parameter value, use either of the following methods:
	//
	// Method 1: Log in to VPC console and click the target subnet on the Subnets page. You can view the network ID on the displayed page.
	// Method 2: See the "Querying Subnets" section under "APIs" or the "Querying Networks" section under "OpenStack Neutron APIs" in Virtual Private Cloud API Reference.
	SubnetId string `json:"subnet_id" required:"true"`
	// Specifies the floating IP address of a DB instance. To obtain this parameter value, use either of the following methods:
	//
	// Method 1: Log in to VPC console and click the target subnet on the Subnets page. You can view the subnet CIDR block on the displayed page.
	// Method 2: See the "Querying Subnets" section under "APIs" in the Virtual Private Cloud API Reference.
	DataVip string `json:"data_vip,omitempty"`
	// Specifies the security group which the RDS DB instance belongs to. To obtain this parameter value, use either of the following methods:
	//
	// Method 1: Log in to VPC console. Choose Access Control > Security Groups in the navigation pane on the left. On the displayed page, click the target security group. You can view the security group ID on the displayed page.
	// Method 2: See the "Querying Security Groups" section in the Virtual Private Cloud API Reference.
	SecurityGroupId string `json:"security_group_id" required:"true"`
	// Specifies the restoration information.
	RestorePoint RestorePoint `json:"restore_point" required:"true"`
	// This parameter applies only to Microsoft SQL Server DB instances.
	Collation string `json:"collation,omitempty"`
}

type RestoreType string

const (
	TypeBackup    RestoreType = "backup"
	TypeTimestamp RestoreType = "timestamp"
)

type RestorePoint struct {
	// Specifies the DB instance ID.
	InstanceID string `json:"instance_id" required:"true"`
	// Specifies the restoration mode. Enumerated values include:
	//
	// backup: indicates restoration from backup files. In this mode, backup_id is mandatory when type is not mandatory.
	// timestamp: indicates point-in-time restoration. In this mode, restore_time is mandatory when type is mandatory.
	Type RestoreType `json:"type" required:"true"`
	// Specifies the ID of the backup used to restore data. This parameter must be specified when the backup file is used for restoration.
	//
	// NOTICE:
	// When type is not mandatory, backup_id is mandatory.
	BackupID string `json:"backup_id,omitempty"`
	// Specifies the time point of data restoration in the UNIX timestamp. The unit is millisecond and the time zone is UTC.
	//
	// NOTICE:
	// When type is mandatory, restore_time is mandatory.
	RestoreTime int `json:"restore_time,omitempty"`
}

func RestoreToNew(c *golangsdk.ServiceClient, opts RestoreToNewOpts) (*instances.CreateRds, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/v3/{project_id}/instances
	raw, err := c.Post(c.ServiceURL("instances"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	if err != nil {
		return nil, err
	}

	var res instances.CreateRds
	err = extract.Into(raw.Body, &res)
	return &res, err
}
