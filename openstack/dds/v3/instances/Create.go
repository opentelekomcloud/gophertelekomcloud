package instances

import (
	"net/http"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type CreateOpts struct {
	// Specifies the DB instance name. Instance name, which can be the same as an existing name.
	// The instance name must contain 4 to 64 characters and must start with a letter. It is case sensitive and can contain letters, digits, hyphens (-), and underscores (_). It cannot contain other special characters.
	Name string `json:"name" required:"true"`
	// Specifies the database information.
	DataStore DataStore `json:"datastore" required:"true"`
	// Specifies the region ID.
	// The value cannot be empty.
	Region string `json:"region" required:"true"`
	// Specifies the AZ ID. You can select multiple AZs to create a cross-AZ cluster based on az_status returned by the API described in Querying Database Specifications.
	AvailabilityZone string `json:"availability_zone" required:"true"`
	// Specifies the VPC ID.
	VpcId string `json:"vpc_id" required:"true"`
	// Specifies the network ID of the subnet.
	SubnetId string `json:"subnet_id" required:"true"`
	// Specifies the security group ID.
	SecurityGroupId string `json:"security_group_id" required:"true"`
	// Database access port
	// Value range: 2100-9500, 27017, 27018, and 27019.
	// If this parameter is not transferred, the port of the created DB instance is 8635 by default.
	Port string `json:"port,omitempty"`
	// Specifies the database password.
	// The value must be 8 to 32 characters in length and contain uppercase letters (A to Z), lowercase letters (a to z), digits (0 to 9), and special characters, such as ~!@#%^*-_=+?
	// Enter a strong password to improve security, preventing security risks such as brute force cracking.
	Password string `json:"password" required:"true"`
	// Specifies the key ID used for disk encryption. The string must comply with UUID regular expression rules.
	// If this parameter is not transferred, disk encryption is not performed.
	DiskEncryptionId string `json:"disk_encryption_id,omitempty"`
	// Specifies the instance type. Cluster, replica set, and single node instances are supported.
	// Valid value:
	// Sharding
	// ReplicaSet
	// Single
	Mode string `json:"mode" required:"true"`
	// Specifies the instance specifications.
	Flavor []Flavor `json:"flavor" required:"true"`
	// Specifies the advanced backup policy.
	BackupStrategy BackupStrategy `json:"backup_strategy" required:"true"`
	// Specifies whether to enable or disable SSL.
	// Valid value:
	// The value 0 indicates that SSL is disabled by default.
	// The value 1 indicates that SSL is enabled by default.
	// If this parameter is not transferred, SSL is enabled by default.
	Ssl string `json:"ssl_option,omitempty"`
	// Tag list
	// A maximum of 20 tags can be added for each instance.
	Tags []tags.ResourceTag `json:"tags,omitempty"`
}

type DataStore struct {
	// Specifies the database type. The value is DDS-Community.
	Type string `json:"type" required:"true"`
	// Specifies the database version. Versions 4.2, 4.0, and 3.4 are supported. The value can be 4.2, 4.0, or 3.4.
	Version string `json:"version" required:"true"`
	// Specifies the storage engine. DDS supports the WiredTiger and RocksDB storage engines.
	// If the database version is 4.2 and the storage engine is RocksDB, the value is rocksDB.
	// If the database version is 4.0 or 3.4 and the storage engine is WiredTiger, the value is wiredTiger.
	StorageEngine string `json:"storage_engine" required:"true"`
}

type Flavor struct {
	// Specifies the node type.
	// Valid value:
	// For a cluster instance, the value can be mongos, shard, or config.
	// For a replica set instance, the value is replica.
	// For a single node instance, the value is single.
	Type string `json:"type" required:"true"`
	// Specifies node quantity.
	// Valid value:
	// mongos: The value ranges from 2 to 32.
	// mongos: The value ranges from 2 to 32.
	// config: The value is 1.
	// replica: The number of nodes can be 3, 5, or 7.
	// single: The value is 1.
	Num int `json:"num" required:"true"`
	// Specifies the disk type.
	// Valid value: ULTRAHIGH, which indicates the type SSD.
	// This parameter is valid for the shard and config nodes of a cluster instance, replica set instances, and single node instances. This parameter is invalid for mongos nodes. Therefore, you do not need to specify the storage space for mongos nodes.
	Storage string `json:"storage,omitempty"`
	// Specifies the disk size.
	// This parameter is mandatory for all nodes except mongos. This parameter is invalid for the mongos nodes.
	// The value must be a multiple of 10. The unit is GB.
	// For a cluster instance, the storage space of a shard node can be 10 to 2000 GB, and the config storage space is 20 GB. This parameter is invalid for mongos nodes. Therefore, you do not need to specify the storage space for mongos nodes.
	// For a replica set instance, the value ranges from 10 to 2000.
	// For a single node instance, the value ranges from 10 to 1000.
	Size int `json:"size,omitempty"`
	// Specifies the resource specification code. For details about how to obtain the value, see the response values of spec_code in Querying Database Specifications.
	// In a cluster instance, multiple specifications need to be specified. All specifications must be of the same series, that is, general-purpose (s6), enhanced (c3), or enhanced II (c6).
	SpecCode string `json:"spec_code" required:"true"`
}

type BackupStrategy struct {
	// Specifies the backup time window. Automated backups will be triggered during the backup time window.
	// The value cannot be empty. It must be a valid value in the "hh:mm-HH:MM" format. The current time is in the UTC format.
	// The HH value must be 1 greater than the hh value.
	// The values of mm and MM must be the same and must be set to 00.
	// If this parameter is not transferred, the default backup time window is set to 00:00-01:00.
	// Example value:
	//23:00-00:00
	StartTime string `json:"start_time" required:"true"`
	// Specifies the number of days to retain the generated backup files.
	// The value range is from 0 to 732.
	// If this parameter is set to 0, the automated backup policy is not set.
	// If this parameter is not transferred, the automated backup policy is enabled by default. Backup files are stored for seven days by default.
	KeepDays *int `json:"keep_days,omitempty"`
	// Specifies the backup cycle configuration. Data will be automatically backed up on the selected days every week.
	// Value range: The value is a number separated by DBS case commas (,). The number indicates the week.
	// The restrictions on the backup retention period are as follows:
	// This parameter is not transferred if its value is set to 0.
	// If you set the retention period to 1 to 6 days, data is automatically backed up each day of the week.
	// Set the parameter value to 1,2,3,4,5,6,7.
	// If you set the retention period to 7 to 732 days, select at least one day of the week for the backup cycle.
	// Example value: 1,2,3,4
	Period string `json:"period,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Instance, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/v3/{project_id}/instances
	raw, err := client.Post(client.ServiceURL("instances"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	return extra(err, raw)
}

func extra(err error, raw *http.Response) (*Instance, error) {
	if err != nil {
		return nil, err
	}

	var res Instance
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Instance struct {
	// Indicates the DB instance ID.
	Id string `json:"id"`
	// Same as the request parameter.
	Name string `json:"name"`
	// Indicates the database information, which is the same as the request parameter.
	DataStore DataStore `json:"datastore"`
	// Indicates the creation time in the following format: yyyy-mm-dd hh:mm:ss.
	CreatedAt string `json:"created"`
	// Indicates the DB instance status. The value is creating.
	Status string `json:"status"`
	// Indicates the region ID, which is the same as the request parameter.
	Region string `json:"region"`
	// Indicates the AZ ID, which is the same as the request parameter.
	AvailabilityZone string `json:"availability_zone"`
	// Indicates the VPC ID, which is the same as the request parameter.
	VpcId string `json:"vpc_id"`
	// Indicates the network ID of the subnet, which is the same as the request parameter.
	SubnetId string `json:"subnet_id"`
	// Indicates the security group ID, which is the same as the request parameter.
	SecurityGroupId string `json:"security_group_id"`
	// Indicates the database port.
	Port int `json:"port"`
	// Indicates the ID of the disk encryption key, which is the same as the request parameter.
	DiskEncryptionId string `json:"disk_encryption_id"`
	// Indicates the instance type, which is the same as the request parameter.
	Mode string `json:"mode"`
	// Indicates the instance specification, which is the same as the request parameter.
	Flavor []FlavorOpt `json:"flavor"`
	// Indicates the advanced backup policy, which is the same as the request parameter.
	BackupStrategy BackupStrategyOpt `json:"backup_strategy"`
	// Indicates whether to enable SSL, which functions the same as the request parameter.
	Ssl string `json:"ssl_option"`
	// Indicates the ID of the workflow for creating a DB instance.
	JobId string `json:"job_id"`
	// Tag list, which is the same as the request parameter.
	Tags []tags.ResourceTag `json:"tags"`
}

type FlavorOpt struct {
	// Specifies the node type.
	// Valid value:
	// For a cluster instance, the value can be mongos, shard, or config.
	// For a replica set instance, the value is replica.
	// For a single node instance, the value is single.
	Type string `json:"type" required:"true"`
	// Specifies node quantity.
	// Valid value:
	// mongos: The value ranges from 2 to 32.
	// mongos: The value ranges from 2 to 32.
	// config: The value is 1.
	// replica: The number of nodes can be 3, 5, or 7.
	// single: The value is 1.
	Num string `json:"num" required:"true"`
	// Specifies the disk type.
	// Valid value: ULTRAHIGH, which indicates the type SSD.
	// This parameter is valid for the shard and config nodes of a cluster instance, replica set instances, and single node instances. This parameter is invalid for mongos nodes. Therefore, you do not need to specify the storage space for mongos nodes.
	Storage string `json:"storage,omitempty"`
	// Specifies the disk size.
	// This parameter is mandatory for all nodes except mongos. This parameter is invalid for the mongos nodes.
	// The value must be a multiple of 10. The unit is GB.
	// For a cluster instance, the storage space of a shard node can be 10 to 2000 GB, and the config storage space is 20 GB. This parameter is invalid for mongos nodes. Therefore, you do not need to specify the storage space for mongos nodes.
	// For a replica set instance, the value ranges from 10 to 2000.
	// For a single node instance, the value ranges from 10 to 1000.
	Size string `json:"size,omitempty"`
	// Specifies the resource specification code. For details about how to obtain the value, see the response values of spec_code in Querying Database Specifications.
	// In a cluster instance, multiple specifications need to be specified. All specifications must be of the same series, that is, general-purpose (s6), enhanced (c3), or enhanced II (c6).
	SpecCode string `json:"spec_code" required:"true"`
}

type BackupStrategyOpt struct {
	StartTime string `json:"start_time" required:"true"`
	KeepDays  string `json:"keep_days,omitempty"`
	Period    string `json:"period,omitempty"`
}
