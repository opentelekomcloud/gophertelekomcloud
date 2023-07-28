package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"net/url"
)

type ListInstanceOpts struct {
	// Specifies the instance ID, which can be obtained by calling the API for querying instances and details.
	Id string `q:"id"`
	// Specifies the DB instance name.
	//
	// If you use asterisk (*) at the beginning of the name, fuzzy search results are returned. Otherwise, the exact results are returned.
	//
	// NOTE:
	// The asterisk (*) is a reserved character in the system and cannot be used alone.
	Name string `q:"name"`
	// Specifies the instance type.
	//
	// Sharding indicates the cluster instance.
	// ReplicaSet indicate the replica set instance.
	// Single indicates the single node instance.
	Mode string `q:"mode"`
	// Specifies the database type. The value is DDS-Community.
	DataStoreType string `q:"datastore_type"`
	// Specifies the VPC ID.
	VpcId string `q:"vpc_id"`
	// Specifies the network ID of the subnet.
	SubnetId string `q:"subnet_id"`
	//
	// Specifies the index position. The query starts from the next instance creation time indexed by this parameter under a specified project. If offset is set to N, the resource query starts from the N+1 piece of data.
	//
	// The value must be greater than or equal to 0. If this parameter is not transferred, offset is set to 0 by default, indicating that the query starts from the latest created DB instance.
	Offset int `q:"offset"`
	// Specifies the maximum allowed number of DB instances.
	//
	// The value ranges from 1 to 100. If this parameter is not transferred, the first 100 DB instances are queried by default.
	Limit int `q:"limit"`
	// Query based on the instance tag key and value.
	//
	// {key} indicates the tag key, and {value} indicates the tag value. A maximum of 20 key-value pairs are supported. The key cannot be empty or duplicate, but the value can be empty.
	//
	// To query instances with multiple tag keys and values, separate key-value pairs with commas (,).
	Tags string `q:"tags"`
}

func List(client *golangsdk.ServiceClient, opts ListInstanceOpts) (*ListResponse, error) {
	var opts2 interface{} = opts
	q, err := build.QueryString(opts2)
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/v3/{project_id}/instances
	raw, err := client.Get(client.ServiceURL("instances")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListResponse struct {
	Instances  []InstanceResponse `json:"instances"`
	TotalCount int                `json:"total_count"`
}

type InstanceResponse struct {
	// Indicates the DB instance ID.
	Id string `json:"id"`
	// Indicates the DB instance name.
	Name string `json:"name"`
	// Instance remarks
	Remark string `json:"remark"`
	// Indicates the DB instance status.
	//
	// Valid value:
	//
	// normal: indicates that the instance is running properly.
	// abnormal: indicates that the instance is abnormal.
	// creating: indicates that the instance is being created.
	// data_disk_full: The storage space is full.
	// createfail: indicates that the instance failed to be created.
	// enlargefail: indicates that nodes failed to be added to the instance.
	// NOTE:
	// Actions that are being executed on an instance, for example, rebooting, which are essentially different from the instance status. For details, see the actions field in this table.
	Status string `json:"status"`
	// Indicates the database port number. The port range is 2100 to 9500.
	Port int `json:"port,string"`
	// Indicates the instance type, which is the same as the request parameter.
	Mode string `json:"mode"`
	// Indicates the region where the DB instance is deployed.
	Region string `json:"region"`
	// Indicates the database information.
	DataStore DataStore `json:"datastore"`
	// Indicates the storage engine. The value is wiredTiger.
	Engine string `json:"engine"`
	// Indicates the DB instance creation time.
	Created string `json:"created"`
	// Indicates the time when a DB instance is updated.
	Updated string `json:"updated"`
	// Indicates the default username. The value is rwuser.
	DbUserName string `json:"db_user_name"`
	// Indicates that SSL is enabled or not.
	//
	// 1: indicate that SSL is enabled.
	// 0: indicate that SSL is disabled.
	Ssl int `json:"ssl"`
	// Indicates the VPC ID.
	VpcId string `json:"vpc_id"`
	// Indicates the network ID of the subnet.
	SubnetId string `json:"subnet_id"`
	// Indicates the security group ID.
	SecurityGroupId string `json:"security_group_id"`
	// Indicates the backup policy.
	BackupStrategy BackupStrategy `json:"backup_strategy"`
	// Indicates the maintenance time window.
	MaintenanceWindow string `json:"maintenance_window"`
	// Indicates group information.
	Groups []Group `json:"groups"`
	// Indicates the disk encryption key ID. This parameter is returned only when the instance disk is encrypted.
	DiskEncryptionId string `json:"disk_encryption_id"`
	// Indicates the time zone.
	TimeZone string `json:"time_zone"`
	// Action that is being executed on an instance.
	//
	// Valid value:
	//
	// RESTARTING: The instance is being restarted.
	// RESTORE: restoring.
	// RESIZE_FLAVOR: The specifications are being changed.
	// RESTORE_TO_NEW_INSTANCE: The instance is being restored.
	// MODIFY_VPC_PEER: Cross-subnet access is being configured.
	// CREATE: creating
	// FROZEN: The account is frozen.
	// RESIZE_VOLUME: The storage is being scaled up.
	// RESTORE_CHECK: The restoration is being checked.
	// RESTORE_FAILED_HANGUP: The restoration failed.
	// CLOSE_AUDIT_LOG: Disabling audit log.
	// OPEN_AUDIT_LOG: Enabling audit log.
	// CREATE_IP_SHARD: The shard IP address is being enabled.
	// CREATE_IP_CONFIG: The config IP address is being enabled.
	// GROWING: The node is being scaled up.
	// SET_CONFIGURATION: Parameters are being modified.
	// RESTORE_TABLE: The database is being backed up.
	// MODIFY_SECURITYGROUP: A security group is being changed.
	// BIND_EIP: The EIP is being changed.
	// UNBIND_EIP: The EIP is being unbound.
	// SWITCH_SSL: The SSL is being switched.
	// SWITCH_PRIMARY: A primary/standby switchover is being performed.
	// CHANGE_DBUSER_PASSWORD: The password is being changed.
	// MODIFY_PORT: The port is being changed.
	// MODIFY_IP: The private IP address is being changed.
	// DELETE_INSTANCE: The instance is being deleted.
	// REBOOT: The system is restarting.
	// BACKUP: The backup is in progress.
	// MIGRATE_AZ: The AZ is being changed.
	// RESTORING: The backup is in progress.
	// PWD_RESETING: The password is being reset.
	// UPGRADE_DATABASE: The patch is being upgraded.
	// DATA_MIGRATION: Data is being migrated.
	// SHARD_GROWING: The shard is being scaled out.
	// APPLY_CONFIGURATION: A parameter group is being changed.
	// RESET_PASSWORD: The password is being reset.
	// GROWING_REVERT: Nodes are being deleted.
	// SHARD_GROWING_REVERT: Shards are being deleted.
	// LOG_PLAINTEXT_SWITCH: The slow query log configuration is being modified.
	// CREATE_DATABASE_USER: The database user is being created.
	// CREATE_DATABASE_ROLE: The database role is being created.
	// MODIFY_NAME: The name is being changed.
	// MODIFY_PRIVATE_DNS: The private zone is being modified.
	// MODIFY_OP_LOG_SIZE: The oplog size is being changed.
	// ADD_READONLY_NODES: Read replicas are being scaled up.
	Actions []string `json:"actions"`
	// The value is set to "0".
	PayMode string `json:"pay_mode"`
	// Tag list
	Tags []tags.ResourceTag `json:"tags"`
}

type Group struct {
	// Indicates the node type.
	//
	// Valid value:
	//
	// shard
	// config
	// mongos
	// replica
	// single
	Type string `json:"type"`
	//
	Id string `json:"id"`
	//
	Name string `json:"name"`
	//
	Status string `json:"status"`
	// Indicates the volume information.
	Volume Volume `json:"volume"`
	// Indicates node information.
	Nodes []Nodes `json:"nodes"`
}

type Volume struct {
	// Indicates the disk size. Unit: GB
	Size string `json:"size"`
	// Indicates the disk usage. Unit: GB
	Used string `json:"used"`
}

type Nodes struct {
	// Indicates the node ID.
	Id string `json:"id"`
	// Indicates the node name.
	Name string `json:"name"`
	// Indicates the node status.
	//
	// Valid value:
	//
	// normal: The instance is running properly.
	// abnormal: The instance is abnormal.
	// backup: The instance is being backed up.
	// frozen: The instance has been frozen.
	// unfrozen: The instance is being unfrozen.
	// restore_table: Database- and table-level backup and restoration are being performed for the DB instance.
	// reboot: The instance is being restarted.
	// upgrade_database: The instance version is being upgraded.
	// resize_flavor: The instance class is being changed.
	// resize_volume: The instance storage is being scaled up.
	// restore: The instance is being restored.
	// bind_eip: An EIP is being bound to the instance.
	Status string `json:"status"`
	// Indicates the node role.
	//
	// Valid value:
	// master: This value is returned for the mongos node.
	// Primary: This value is returned for the primary shard and config nodes, the primary node of a replica set, and a single node.
	// Secondary: This value is returned for the secondary shard and config nodes, and the secondary node of a replica set.
	// Hidden: This value is returned for the hidden shard and config nodes, and the hidden node of a replica set.
	// unknown. This value is returned when the node is abnormal.
	Role string `json:"role"`
	// Indicates the private IP address of a node. By default, this parameter is valid only for mongos nodes, replica set instances, and single node instances. The value exists only after ECSs are created successfully. Otherwise, the value is "".
	//
	// CAUTION:
	// After the shard or config IP address is enabled, private IP addresses are assigned to the primary and secondary shard or config nodes of the cluster instance.
	PrivateIP string `json:"private_ip"`
	// Indicates the EIP that has been bound. This parameter is valid only for mongos nodes of cluster instances, primary nodes and secondary nodes of replica set instances, and single node instances.
	PublicIP string `json:"public_ip"`
	// Indicates the resource specification code.
	SpecCode string `json:"spec_code"`
	// Indicates the AZ.
	AvailabilityZone string `json:"availability_zone"`
}
