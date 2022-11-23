package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateRdsOpts struct {
	// Specifies the DB instance name.
	// DB instances of the same type can have same names under the same tenant.
	// The value must be 4 to 64 characters in length and start with a letter.
	// It is case-sensitive and can contain only letters, digits, hyphens (-), and underscores (_).
	Name string `json:"name"  required:"true"`
	// Specifies the database information.
	Datastore *Datastore `json:"datastore" required:"true"`
	// Specifies the HA configuration parameters, which are used when creating primary/standby DB instances.
	Ha *Ha `json:"ha,omitempty"`
	// Specifies the parameter template ID.
	ConfigurationId string `json:"configuration_id,omitempty"`
	// Specifies the database port information.
	// The MySQL database port ranges from 1024 to 65535 (excluding 12017 and 33071, which are occupied by the RDS system and cannot be used).
	// The PostgreSQL database port ranges from 2100 to 9500.
	// The Microsoft SQL Server database port is 1433 or ranges from 2100 to 9500 (excluding 5355 and 5985).
	// If this parameter is not set, the default value is as follows:
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
	BackupStrategy *BackupStrategy `json:"backup_strategy,omitempty"`
	// Specifies the key ID for disk encryption. The default value is empty.
	DiskEncryptionId string `json:"disk_encryption_id,omitempty"`
	// Specifies the specification code. The value cannot be empty.
	FlavorRef string `json:"flavor_ref" required:"true"`
	// Specifies the volume information.
	Volume *Volume `json:"volume" required:"true"`
	// Specifies the region ID. The value cannot be empty.
	Region string `json:"region" required:"true"`
	// Specifies the AZ ID. If the DB instance is not a single instance, you need to specify an AZ for each node of the instance and separate the AZs with commas (,). For details, see the example.
	// The value cannot be empty. For details about how to obtain this parameter value, see Regions and Endpoints.
	AvailabilityZone string `json:"availability_zone" required:"true"`
	// Specifies the VPC ID. To obtain this parameter value, use either of the following methods:
	// Method 1: Log in to VPC console and view the VPC ID in the VPC details.
	// Method 2: See the "Querying VPCs" section in the Virtual Private Cloud API Reference.
	VpcId string `json:"vpc_id" required:"true"`
	// Specifies the network ID. To obtain this parameter value, use either of the following methods:
	// Method 1: Log in to VPC console and click the target subnet on the Subnets page. You can view the network ID on the displayed page.
	// Method 2: See the "Querying Subnets" section under "APIs" or the "Querying Networks" section under "OpenStack Neutron APIs" in Virtual Private Cloud API Reference.
	SubnetId string `json:"subnet_id" required:"true"`
	// Specifies the private IP address of a DB instance. You can use the following methods to obtain the private IP address:
	// Method 1: Log in to VPC console and click the target subnet on the Subnets page. You can view the subnet CIDR block on the displayed page.
	// Method 2: See the "Querying Subnets" section under "APIs" in the Virtual Private Cloud API Reference.
	DataVip string `json:"data_vip,omitempty"`
	// Specifies the security group which the RDS DB instance belongs to. To obtain this parameter value, use either of the following methods:
	// Method 1: Log in to VPC console. Choose Access Control > Security Groups in the navigation pane on the left. On the displayed page, click the target security group. You can view the security group ID on the displayed page.
	// Method 2: See the "Querying Security Groups" section in the Virtual Private Cloud API Reference.
	SecurityGroupId string `json:"security_group_id" required:"true"`
	// Specifies the billing information, which is pay-per-use. By default, pay-per-use is used.
	ChargeInfo *ChargeInfo `json:"charge_info,omitempty"`
	// This parameter applies only to Microsoft SQL Server DB instances.
	Collation string `json:"collation,omitempty"`
}

type Datastore struct {
	// Specifies the DB engine. Value:
	// MySQL
	// PostgreSQL
	// SQLServer
	Type string `json:"type,omitempty"`
	// Specifies the database version.
	// MySQL databases support 5.6, 5.7, and 8.0. Example value: 5.7
	// PostgreSQL databases support 9.5, 9.6, 10, 11, 12, 13 and 14. Example value: 9.6
	// Microsoft SQL Server databases only support 2014 SE, 2016 SE, 2016 EE, 2017 SE, 2017 EE, 2019 SE and 2109 EE. Example value: 2014_SE
	// For details about supported database versions, see section Querying Version Information About a DB Engine.
	Version string `json:"version,omitempty"`
	// Indicates the complete version number. This parameter is returned only when the DB engine is PostgreSQL.
	CompleteVersion string `json:"complete_version,omitempty"`
}

type Ha struct {
	// Specifies the primary/standby or cluster instance type. The value is Ha (case-insensitive).
	Mode string `json:"mode,omitempty"`
	// Specifies the replication mode for the standby DB instance.
	// Value:
	// For MySQL, the value is async or semisync.
	// For PostgreSQL, the value is async or sync.
	// For Microsoft SQL Server, the value is sync.
	// NOTE
	// async indicates the asynchronous replication mode.
	// semisync indicates the semi-synchronous replication mode.
	// sync indicates the synchronous replication mode.
	ReplicationMode string `json:"replication_mode,omitempty"`
}

type BackupStrategy struct {
	// Specifies the backup time window. Automated backups will be triggered during the backup time window.
	// The value cannot be empty. It must be a valid value in the "hh:mm-HH:MM" format. The current time is in the UTC format.
	// The HH value must be 1 greater than the hh value.
	// The values of mm and MM must be the same and must be set to any of the following: 00, 15, 30, or 45.
	// Example value:
	// 08:15-09:15
	// 23:00-00:00
	StartTime string `json:"start_time" required:"true"`
	// Specifies the retention days for specific backup files.
	// The value range is from 0 to 732. If this parameter is not specified or set to 0, the automated backup policy is disabled. To extend the retention period, contact customer service. Automated backups can be retained for up to 2562 days.
	// NOTICE
	// Primary/standby DB instances and Cluster DB instances of Microsoft SQL Server do not support disabling the automated backup policy.
	KeepDays int `json:"keep_days,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateRdsOpts) (*CreateRds, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/v3/{project_id}/instances
	raw, err := client.Post(client.ServiceURL("instances"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return extra(err, raw)
}

type CreateRds struct {
	// Indicates the DB instance information.
	Instance Instance `json:"instance"`
	// Indicates the ID of the DB instance creation task.
	JobId string `json:"job_id"`

	OrderId string `json:"order_id"`
}

type Instance struct {
	// Indicates the DB instance ID.
	// NOTE
	// The v3 DB instance ID is incompatible with the v1 DB instance ID.
	Id string `json:"id"`
	// Indicates the DB instance name. Indicates the DB instance name. DB instances of the same type can have same names under the same tenant.
	// The value must be 4 to 64 characters in length and start with a letter. It is case-insensitive and can contain only letters, digits, hyphens (-), and underscores (_).
	Name string `json:"name"`
	// Indicates the DB instance status. For example, BUILD indicates that the DB instance is being created.
	Status string `json:"status"`
	// Indicates the database information.
	Datastore Datastore `json:"datastore"`
	// Indicates the HA configuration parameters. This parameter is returned only when primary/standby DB instances are created
	Ha Ha `json:"ha"`
	// Indicates the parameter template ID. This parameter is returned only when a custom parameter template is used during DB instance creation.
	ConfigurationId string `json:"configuration_id"`
	// Indicates the database port, which is the same as the request parameter.
	Port string `json:"port"`
	// Indicates the automated backup policy.
	BackupStrategy BackupStrategy `json:"backup_strategy"`
	// Indicates the key ID for disk encryption. By default, this parameter is empty and is returned only when it is specified during the DB instance creation.
	DiskEncryptionId string `json:"disk_encryption_id"`
	// Indicates the specification code. The value cannot be empty.
	FlavorRef string `json:"flavor_ref"`
	// Indicates the volume information.
	Volume Volume `json:"volume"`
	// Indicates the region ID.
	Region string `json:"region"`
	// Indicates the AZ ID.
	AvailabilityZone string `json:"availability_zone"`
	// Indicates the VPC ID. To obtain this parameter value, use either of the following methods:
	// Method 1: Log in to VPC console and view the VPC ID in the VPC details.
	// Method 2: See the "Querying VPCs" section in the Virtual Private Cloud API Reference.
	VpcId string `json:"vpc_id"`
	// Indicates the network ID. To obtain this parameter value, use either of the following methods:
	// Method 1: Log in to VPC console and click the target subnet on the Subnets page. You can view the network ID on the displayed page.
	// Method 2: See the "Querying Subnets" section under "APIs" or the "Querying Networks" section under "OpenStack Neutron APIs" in Virtual Private Cloud API Reference.
	SubnetId string `json:"subnet_id"`
	// Indicates the security group which the RDS DB instance belongs to. To obtain this parameter value, use either of the following methods:
	// Method 1: Log in to VPC console. Choose Access Control > Security Groups in the navigation pane on the left. On the displayed page, click the target security group. You can view the security group ID on the displayed page.
	// Method 2: See the "Querying Security Groups" section in the Virtual Private Cloud API Reference.
	SecurityGroupId string `json:"security_group_id"`
	// Indicates the billing information, which is pay-per-use.
	ChargeInfo ChargeInfo `json:"charge_info"`
	// Indicates the Collation set for Microsoft SQL Server.
	Collation string `json:"collation"`
}
