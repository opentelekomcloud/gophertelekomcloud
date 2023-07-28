package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type ListOpts struct {
	// Specifies the DB instance ID.
	// The asterisk (*) is reserved for the system. If the instance ID starts with *, it indicates that fuzzy match is performed based on the value following * Otherwise, the exact match is performed based on the instance ID. The value cannot contain only asterisks (*).
	Id string `q:"id"`
	// Specifies the DB instance name.
	// The asterisk (*) is reserved for the system. If the instance name starts with *, it indicates that fuzzy match is performed based on the value following * Otherwise, the exact match is performed based on the instance name. The value cannot contain only asterisks (*).
	Name string `q:"name"`
	// Specifies the instance type based query. The value is Single, Ha, or Replica, which correspond to single instance, primary/standby instances, and read replica, respectively.
	Type string `q:"type"`
	// Specifies the database type. Its value can be any of the following and is case-sensitive:
	// MySQL
	// PostgreSQL
	// SQLServer
	DataStoreType string `q:"datastore_type"`
	// Specifies the VPC ID.
	// Method 1: Log in to VPC console and view the VPC ID in the VPC details.
	// Method 2: See the "Querying VPCs" section in the Virtual Private Cloud API Reference.
	VpcId string `q:"vpc_id"`
	// Specifies the network ID of the subnet.
	// Method 1: Log in to VPC console and click the target subnet on the Subnets page. You can view the network ID on the displayed page.
	// Method 2: See the "Querying Subnets" section under "APIs" or the "Querying Networks" section under "OpenStack Neutron APIs" in Virtual Private Cloud API Reference.
	SubnetId string `q:"subnet_id"`
	// Specifies the index position. If offset is set to N, the resource query starts from the N+1 piece of data. The value is 0 by default, indicating that the query starts from the first piece of data. The value must be a positive number.
	Offset int `q:"offset"`
	// Specifies the number of records to be queried. The default value is 100. The value cannot be a negative number. The minimum value is 1 and the maximum value is 100.
	Limit int `q:"limit"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListResponse, error) {
	q, err := build.QueryString(opts)
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
	// Indicates the created DB instance name.
	Name string `json:"name"`
	// Indicates the DB instance status.
	// Value:
	// If the value is BUILD, the instance is being created.
	// If the value is ACTIVE, the instance is normal.
	// If the value is FAILED, the instance is abnormal.
	// If the value is MODIFYING, the instance is being scaled up.
	// If the value is REBOOTING, the instance is being rebooted.
	// If the value is RESTORING, the instance is being restored.
	// If the value is MODIFYING INSTANCE TYPE, the instance is changing from primary to standby.
	// If the value is SWITCHOVER, the primary/standby switchover is being performed.
	// If the value is MIGRATING, the instance is being migrated.
	// If the value is BACKING UP, the instance is being backed up.
	// If the value is MODIFYING DATABASE PORT, the database port is being changed.
	// If the value is SHUTDOWN, the DB instance is stopped.
	Status string `json:"status"`
	// Indicates the DB instance alias.
	Alias string `json:"alias"`
	// Indicates the private IP address list. It is a blank string until an ECS is created.
	PrivateIps []string `json:"private_ips"`
	// Indicates the public IP address list.
	PublicIps []string `json:"public_ips"`
	// Indicates the database port number.
	// The MySQL database port ranges from 1024 to 65535 (excluding 12017 and 33071, which are occupied by the RDS system and cannot be used).
	// The PostgreSQL database port ranges from 2100 to 9500.
	// The Microsoft SQL Server database port is 1433 or ranges from 2100 to 9500 (excluding 5355 and 5985).
	// If this parameter is not set, the default value is as follows:
	// For MySQL, the default value is 3306.
	// For PostgreSQL, the default value is 5432.
	// For Microsoft SQL Server, the default value is 1433.
	Port int `json:"port"`
	// The value is Single, Ha, or Replica, which correspond to single instance, primary/standby instances, and read replica, respectively.
	Type string `json:"type"`
	// Indicates the primary/standby DB instance information. Returned only when you obtain a primary/standby DB instance list.
	Ha Ha `json:"ha"`
	// Indicates the region where the DB instance is deployed.
	Region string `json:"region"`
	// Indicates the database information.
	DataStore Datastore `json:"datastore"`
	// Indicates the creation time in the "yyyy-mm-ddThh:mm:ssZ" format.
	// T is the separator between the calendar and the hourly notation of time. Z indicates the time zone offset.
	// The value is empty when the DB instance is being created. After the DB instance is created, the value is not empty.
	Created string `json:"created"`
	// Indicates the update time. The format is the same as that of the created field.
	// The value is empty when the DB instance is being created. After the DB instance is created, the value is not empty.
	Updated string `json:"updated"`
	// Indicates the default username.
	DbUserName string `json:"db_user_name"`
	// Indicates the VPC ID.
	VpcId string `json:"vpc_id"`
	// Indicates the network ID of the subnet.
	SubnetId string `json:"subnet_id"`
	// Indicates the security group ID.
	SecurityGroupId string `json:"security_group_id"`
	// Indicates the number of CPUs. For example, the value 1 indicates 1 vCPU.
	Cpu string `json:"cpu"`
	// Indicates the memory size in GB.
	Mem string `json:"mem"`
	// Indicates the specification code.
	FlavorRef string `json:"flavor_ref"`
	// Indicates the volume information.
	Volume Volume `json:"volume"`
	// Indicates the database switchover policy. The value can be reliability or availability, indicating the reliability first and availability first, respectively.
	SwitchStrategy string `json:"switch_strategy"`
	// Indicates the backup policy.
	BackupStrategy BackupStrategy `json:"backup_strategy"`
	// Indicates the start time of the maintenance time window in the UTC format.
	MaintenanceWindow string `json:"maintenance_window"`
	// Indicates the primary/standby DB instance information.
	Nodes []Nodes `json:"nodes"`
	// Indicates the list of associated DB instances
	RelatedInstance []RelatedInstance `json:"related_instance"`
	// Indicates the disk encryption key ID.
	DiskEncryptionId string `json:"disk_encryption_id"`
	// Indicates the time zone.
	TimeZone string `json:"time_zone"`
	// Indicates the billing information, which is pay-per-use.
	ChargeInfo ChargeInfo `json:"charge_info"`
	// Indicates the tag list. If there is no tag in the list, an empty array is returned.
	Tags []tags.ResourceTag `json:"tags"`
	// Indicates whether a DDM instance has been associated.
	AssociatedWithDdm bool `json:"associated_with_ddm"`
}

type Nodes struct {
	// Indicates the node ID.
	Id string `json:"id"`
	// Indicates the node name.
	Name string `json:"name"`
	// Indicates the node type. The value can be master, slave, or readreplica, indicating the primary node, standby node, and read replica node, respectively.
	Role string `json:"role"`
	// Indicates the node status.
	Status string `json:"status"`
	// Indicates the AZ.
	AvailabilityZone string `json:"availability_zone"`
}

type RelatedInstance struct {
	// Indicates the associated DB instance ID.
	Id string `json:"id"`
	// Indicates the associated DB instance type.
	// replica_of: indicates the primary DB instance.
	// replica: indicates read replicas.
	Type string `json:"type"`
}
