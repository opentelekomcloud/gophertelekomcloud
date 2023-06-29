package public

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type BatchCreateTaskOpts struct {
	Jobs []CreateJobOpts `json:"jobs"`
}

type CreateJobOpts struct {
	// Whether to bind an EIP. This parameter is mandatory and set to true when the network type is EIP.
	BindEip *bool `json:"bind_eip,omitempty"`
	// The migration scenario. The value can be migration (real-time migration), sync (real-time synchronization),
	// or cloudDataGuard (real-time disaster recovery).
	// Values: migration sync cloudDataGuard
	DbUseType string `json:"db_use_type" required:"true"`
	// The task name. The task name can be 4 to 50 characters in length.
	// It is case-insensitive and can contain only letters, digits, hyphens (-), and underscores (_).
	// Minimum length: 4 characters
	// Maximum length: 50 characters
	Name string `json:"name" required:"true"`
	// Task description. The task description can contain a maximum of 256 characters
	// and cannot contain the following special characters: !<>'&"\
	Description string `json:"description,omitempty"`
	// The engine type. The options are as follows: mysql (MySQL migration and MySQL synchronization),
	// mongodb (used for migration), mysql-to-taurus (used for MySQL to GaussDB(for MySQL) primary/standby synchronization),
	// cloudDataGuard-mysql (used for DR), postgresql (used for PostgreSQL synchronization).
	// Values: mysql mongodb cloudDataGuard-mysql mysql-to-taurus postgresql
	EngineType string `json:"engine_type" required:"true"`
	// Whether the destination DB instance can be read-only. This parameter is valid only when the destination
	// DB instance is a MySQL DB instance and the job_direction value is up.
	// In the DR scenario, this parameter is mandatory and set to true if the current cloud is a standby cloud.
	// If this parameter is not specified, the default value is true.
	IsTargetReadonly *bool `json:"is_target_readonly,omitempty"`
	// l The migration direction. The value can be up (to the cloud and current cloud as standby in disaster recovery),
	// down (out of cloud and current cloud as active in disaster recovery), or non-dbs (for self-built databases).
	// Values: − up − down − non-dbs
	JobDirection string `json:"job_direction" required:"true"`
	// This parameter is mandatory when db_use_type is set to cloudDataGuard.
	// If the DR type is dual-active, the value of multi_write is true. Otherwise, the value is false.
	// If db_use_type is set to other values, multi_write is optional.
	// Default value: false
	MultiWrite *bool `json:"multi_write,omitempty"`
	// Network type.
	// Value: vpn vpc eip
	// The VPC network cannot be selected in the DR scenario.
	NetType string `json:"net_type" required:"true"`
	// The number of nodes. For a MongoDB database, this parameter indicates the number of source shards.
	// This parameter is mandatory when the source database is a cluster.
	// The value ranges from 1 to 32. The default value is 2 for MySQL dual-active DR.
	NodeNum int `json:"node_num,omitempty"`
	// The flavor type. Value: high
	NodeType string `json:"node_type" required:"true"`
	// The source database information.
	SourceEndpoint Endpoint `json:"source_endpoint" required:"true"`
	// The destination database information.
	TargetEndpoint Endpoint `json:"target_endpoint" required:"true"`
	// Tag information. Up to 20 tags can be added.
	Tags []tags.ResourceTag `json:"tags,omitempty"`
	// The migration type. FULL_TRANS indicates full migration.
	// FULL_INCR_TRANS: indicates full+incremental migration.
	// INCR_TRANS: indicates incremental migration.
	// In single-active DR scenarios, only full plus incremental migration (FULL_INCR_TRANS) is supported.
	// Default value: FULL_INCR_TRANS
	// Values: FULL_TRANS FULL_INCR_TRANS INCR_TRANS
	TaskType string `json:"task_type" required:"true"`
	// ID of the subnet associated with the DRS instance.
	CustomizeSubnetId string `json:"customize_sutnet_id" required:"true"`
	// After a task is in the abnormal status for a period of time, the task is automatically stopped.
	// The unit is day. The value ranges from 14 to 100. If this parameter is not transferred, the default value is 14.
	ExpiredDays string `json:"expired_days,omitempty"`
	// AzCode of the active node. This parameter is mandatory for cross-AZ node tasks.
	MasterAz string `json:"master_az,omitempty"`
	// AzCode of the standby node. This parameter is mandatory for cross-AZ node tasks.
	SlaveAz string `json:"slave_az,omitempty"`
}

type Endpoint struct {
	// Database type.
	// Values: mysql mongodb taurus postgresql
	DbType string `json:"db_type,omitempty"`
	// azCode of the AZ where the database is located.
	AzCode string `json:"az_code,omitempty"`
	// Region where the RDS or GaussDB(for MySQL) instance is located.
	// This parameter is mandatory when an RDS or GaussDB(for MySQL) instance is used.
	// In DR scenarios, this parameter is mandatory in source_endpoint when job_direction is down
	// and is mandatory in target_endpoint when job_direction is up.
	Region string `json:"region,omitempty"`
	// RDS or GaussDB(for MySQL) instance ID.
	// This parameter is mandatory when an RDS or GaussDB(for MySQL) instance is used.
	// In DR scenarios, this parameter is mandatory in source_endpoint
	// when job_direction is down and is mandatory in target_endpoint when job_direction is up.
	InstId string `json:"inst_id,omitempty"`
	// ID of the VPC where the database is located.
	VpcId string `json:"vpc_id,omitempty"`
	// ID of the subnet where the database is located.
	SubnetId string `json:"subnet_id,omitempty"`
	// ID of the security group to which the database belongs.
	SecurityGroupId string `json:"security_group_id,omitempty"`
	// The project ID of an RDS or GaussDB(for MySQL) instance.
	ProjectId string `json:"project_id,omitempty"`
	// Database password.
	DbPassword string `json:"db_password,omitempty"`
	// Database port. The value is an integer ranging from 1 to 65535.
	DbPort int `json:"db_port,omitempty"`
	// Database user.
	DbUser string `json:"db_user,omitempty"`
	// The name of an RDS or GaussDB(for MySQL) instance.
	InstName string `json:"inst_name,omitempty"`
	// Database IP address.
	Ip string `json:"ip,omitempty"`
	// Mongo HA mode.
	MongoHaMode string `json:"mongo_ha_mode,omitempty"`
	// SSL certificate password. The certificate file name extension is .p12.
	SslCertPassword string `json:"ssl_cert_password,omitempty"`
	// The checksum value of the SSL certificate, which is used for backend verification.
	// This parameter is mandatory for secure connection to the source database.
	SslCertCheckSum string `json:"ssl_cert_check_sum,omitempty"`
	// SSL certificate content, which is encrypted using Base64.
	SslCertKey string `json:"ssl_cert_key,omitempty"`
	// SSL certificate name.
	SslCertName string `json:"ssl_cert_name,omitempty"`
	// Whether SSL is enabled.
	SslLink bool `json:"ssl_link,omitempty"`
	// For MongoDB 4.0 or later, if the cluster instance cannot obtain the IP address of the sharded node,
	// set source_endpoint to Sharding4.0+. Default value: Sharding4.0+
	// Values: Sharding4.0+
	ClusterMode string `json:"cluster_mode,omitempty"`
}

func BatchCreateTasks(client *golangsdk.ServiceClient, opts BatchCreateTaskOpts) (*BatchCreateTasksResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/jobs/batch-creation
	raw, err := client.Post(client.ServiceURL("jobs", "batch-creation"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res BatchCreateTasksResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type BatchCreateTasksResponse struct {
	Results []CreateTaskResp `json:"results,omitempty"`
	Count   int              `json:"count,omitempty"`
}

type CreateTaskResp struct {
	// Task ID.
	Id string `json:"id"`
	// Task name.
	Name string `json:"name,omitempty"`
	// Task status.
	Status string `json:"status,omitempty"`
	// Creation time (timestamp).
	CreateTime string `json:"create_time,omitempty"`
	// Error code, which is optional and indicates the returned information about the failure status.
	ErrorCode string `json:"error_code,omitempty"`
	// Error message, which is optional and indicates the returned information about the failure status.
	ErrorMsg string `json:"error_msg,omitempty"`
	// Subtask ID set.
	ChildIds []string `json:"child_ids,omitempty"`
}
