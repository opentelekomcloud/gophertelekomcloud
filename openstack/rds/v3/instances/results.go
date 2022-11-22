package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type commonResult struct {
	golangsdk.Result
}

type CreateResult struct {
	commonResult
}

type RestartRdsInstanceResult struct {
	commonResult
}

type SingleToHaRdsInstanceResult struct {
	commonResult
}

type ResizeFlavorResult struct {
	commonResult
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

func (r CreateResult) Extract() (*CreateRds, error) {
	var response CreateRds
	err := r.ExtractInto(&response)
	return &response, err
}

type DeleteInstanceRdsResult struct {
	commonResult
}

func (r DeleteInstanceRdsResult) Extract() (*DeleteInstanceRdsResponse, error) {
	var response DeleteInstanceRdsResponse
	err := r.ExtractInto(&response)
	return &response, err
}

func (r RestartRdsInstanceResult) Extract() (*RestartRdsResponse, error) {
	var response RestartRdsResponse
	err := r.ExtractInto(&response)
	return &response, err
}

func (r SingleToHaRdsInstanceResult) Extract() (*SingleToHaResponse, error) {
	var response SingleToHaResponse
	err := r.ExtractInto(&response)
	return &response, err
}

func (r ResizeFlavorResult) Extract() (*ResizeFlavor, error) {
	var response ResizeFlavor
	err := r.ExtractInto(&response)
	return &response, err
}

type EnlargeVolumeResult struct {
	commonResult
}

func (r EnlargeVolumeResult) Extract() (*EnlargeVolumeResp, error) {
	var response EnlargeVolumeResp
	err := r.ExtractInto(&response)
	return &response, err
}

type ListRdsResult struct {
	commonResult
}

type ListRdsResponse struct {
	//
	Instances []RdsInstanceResponse `json:"instances"`
	//
	TotalCount int `json:"total_count"`
}

type RdsInstanceResponse struct {
	//
	Id string `json:"id"`
	//
	Name string `json:"name"`
	//
	Status string `json:"status"`
	//
	PrivateIps []string `json:"private_ips"`
	//
	PublicIps []string `json:"public_ips"`
	//
	Port int `json:"port"`
	//
	Type string `json:"type"`
	//
	Ha Ha `json:"ha"`
	//
	Region string `json:"region"`
	//
	DataStore Datastore `json:"datastore"`
	//
	Created string `json:"created"`
	//
	Updated string `json:"updated"`
	//
	DbUserName string `json:"db_user_name"`
	//
	VpcId string `json:"vpc_id"`
	//
	SubnetId string `json:"subnet_id"`
	//
	SecurityGroupId string `json:"security_group_id"`
	//
	FlavorRef string `json:"flavor_ref"`
	//
	Volume Volume `json:"volume"`
	//
	SwitchStrategy string `json:"switch_strategy"`
	//
	BackupStrategy BackupStrategy `json:"backup_strategy"`
	//
	MaintenanceWindow string `json:"maintenance_window"`
	//
	Nodes []Nodes `json:"nodes"`
	//
	RelatedInstance []RelatedInstance `json:"related_instance"`
	//
	DiskEncryptionId string `json:"disk_encryption_id"`
	//
	EnterpriseProjectId string `json:"enterprise_project_id"`
	//
	TimeZone string `json:"time_zone"`

	Tags []tags.ResourceTag `json:"tags"`
}

type Nodes struct {
	//
	Id string `json:"id"`
	//
	Name string `json:"name"`
	//
	Role string `json:"role"`
	//
	Status string `json:"status"`
	//
	AvailabilityZone string `json:"availability_zone"`
}

type RelatedInstance struct {
	//
	Id string `json:"id"`
	//
	Type string `json:"type"`
}

type RdsPage struct {
	pagination.SinglePageBase
}

func (r RdsPage) IsEmpty() (bool, error) {
	data, err := ExtractRdsInstances(r)
	if err != nil {
		return false, err
	}
	return len(data.Instances) == 0, err
}

// ExtractCloudServers is a function that takes a ListResult and returns the services' information.
func ExtractRdsInstances(r pagination.Page) (ListRdsResponse, error) {
	var s ListRdsResponse
	err := (r.(RdsPage)).ExtractInto(&s)
	return s, err
}

type ErrorLogResult struct {
	golangsdk.Result
}

type ErrorLogResp struct {
	//
	ErrorLogList []Errorlog `json:"error_log_list"`
	//
	TotalRecord int `json:"total_record"`
}

type Errorlog struct {
	//
	Time string `json:"time"`
	//
	Level string `json:"level"`
	//
	Content string `json:"content"`
}

type ErrorLogPage struct {
	pagination.SinglePageBase
}

// IsEmpty returns true if a ListResult contains no services.
func (r ErrorLogPage) IsEmpty() (bool, error) {
	data, err := ExtractErrorLog(r)
	if err != nil {
		return false, err
	}
	return len(data.ErrorLogList) == 0, err
}

func ExtractErrorLog(r pagination.Page) (ErrorLogResp, error) {
	var s ErrorLogResp
	err := (r.(ErrorLogPage)).ExtractInto(&s)
	return s, err
}

type SlowLogResp struct {
	//
	Slowloglist []Slowloglist `json:"slow_log_list"`
	//
	TotalRecord int `json:"total_record"`
}

type Slowloglist struct {
	//
	Count string `json:"count"`
	//
	Time string `json:"time"`
	//
	Locktime string `json:"lock_time"`
	//
	Rowssent string `json:"rows_sent"`
	//
	Rowsexamined string `json:"rows_examined"`
	//
	Database string `json:"database"`
	//
	Users string `json:"users"`
	//
	QuerySample string `json:"query_sample"`
	//
	Type string `json:"type"`
}

type SlowLogPage struct {
	pagination.SinglePageBase
}

// IsEmpty returns true if a ListResult contains no services.
func (r SlowLogPage) IsEmpty() (bool, error) {
	data, err := ExtractSlowLog(r)
	if err != nil {
		return false, err
	}
	return len(data.Slowloglist) == 0, err
}

// ExtractCloudServers is a function that takes a ListResult and returns the services' information.
func ExtractSlowLog(r pagination.Page) (SlowLogResp, error) {
	var s SlowLogResp
	err := (r.(SlowLogPage)).ExtractInto(&s)
	return s, err
}

type UpdateConfigurationResponse struct {
	RestartRequired bool `json:"restart_required"`
}

type UpdateConfigurationResult struct {
	golangsdk.Result
}

func (r UpdateConfigurationResult) Extract() (*UpdateConfigurationResponse, error) {
	restartRequired := new(UpdateConfigurationResponse)
	err := r.ExtractInto(restartRequired)
	if err != nil {
		return nil, err
	}
	return restartRequired, nil
}
