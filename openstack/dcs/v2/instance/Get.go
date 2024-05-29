package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// Get a instance with detailed information by id
func Get(client *golangsdk.ServiceClient, id string) (*DcsInstance, error) {
	raw, err := client.Get(client.ServiceURL("instances", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res DcsInstance
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type DcsInstance struct {
	VpcName                   string               `json:"vpc_name"`
	ChargingMode              int                  `json:"charging_mode"`
	VpcId                     string               `json:"vpc_id"`
	UserName                  string               `json:"user_name"`
	CreatedAt                 string               `json:"created_at"`
	LaunchedAt                string               `json:"launched_at"`
	Description               string               `json:"description"`
	SecurityGroupId           string               `json:"security_group_id"`
	SecurityGroupName         string               `json:"security_group_name"`
	MaxMemory                 int                  `json:"max_memory"`
	UsedMemory                int                  `json:"used_memory"`
	Capacity                  float64              `json:"capacity"`
	CapacityMinor             string               `json:"capacity_minor"`
	MaintainBegin             string               `json:"maintain_begin"`
	MaintainEnd               string               `json:"maintain_end"`
	Engine                    string               `json:"engine"`
	NoPasswordAccess          string               `json:"no_password_access"`
	Ip                        string               `json:"ip"`
	BackupPolicy              InstanceBackupPolicy `json:"instance_backup_policy"`
	AzCodes                   []string             `json:"az_codes"`
	AccessUser                string               `json:"access_user"`
	InstanceID                string               `json:"instance_id"`
	Port                      int                  `json:"port"`
	UserId                    string               `json:"user_id"`
	Name                      string               `json:"name"`
	SpecCode                  string               `json:"spec_code"`
	SubnetId                  string               `json:"subnet_id"`
	SubnetName                string               `json:"subnet_name"`
	SubnetCidr                string               `json:"subnet_cidr"`
	EngineVersion             string               `json:"engine_version"`
	OrderId                   string               `json:"order_id"`
	Status                    string               `json:"status"`
	DomainName                string               `json:"domain_name"`
	EnablePublicIp            bool                 `json:"enable_publicip"`
	PublicIpId                string               `json:"publicip_id"`
	PublicIpAddress           string               `json:"publicip_address"`
	EnableSsl                 bool                 `json:"enable_ssl"`
	ServiceUpgrade            bool                 `json:"service_upgrade"`
	ServiceTaskId             string               `json:"service_task_id"`
	BackendAddress            string               `json:"backend_addrs"`
	BandWidthDetail           BandWidthInfo        `json:"bandwidth_info"`
	CacheMode                 string               `json:"cache_mode"`
	CpuType                   string               `json:"cpu_type"`
	ReplicaCount              int                  `json:"replica_count"`
	ReadOnlyDomainName        string               `json:"readonly_domain_name"`
	TransparentClientIpEnable bool                 `json:"transparent_client_ip_enable"`
	ShardingCount             int                  `json:"sharding_count"`
	ProductType               string               `json:"product_type"`
	InquerySpecCode           string               `json:"inquery_spec_code"`
	CloudResourceTypeCode     string               `json:"cloud_resource_type_code"`
	CloudServiceTypeCode      string               `json:"cloud_service_type_code"`
	DbNumber                  int                  `json:"db_number"`
	SupportSlowLogFlag        string               `json:"support_slow_log_flag"`
	StorageType               string               `json:"storage_type"`
	UpdateAt                  string               `json:"update_at"`
	Tags                      []tags.ResourceTag   `json:"tags"`
	SubStatus                 string               `json:"sub_status"`
	DomainNameInfo            DomainNameInfo       `json:"domain_name_info"`
	Features                  Features             `json:"features"`
}

type InstanceBackupPolicy struct {
	BackupPolicyId string          `json:"backup_policy_id"`
	Policy         DcsBackupPolicy `json:"policy"`
	CreatedAt      string          `json:"created_at"`
	UpdatedAt      string          `json:"updated_at"`
	TenantId       string          `json:"tenant_id"`
}

type DcsBackupPolicy struct {
	BackupType           string     `json:"backup_type"`
	SaveDays             int        `json:"save_days"`
	PeriodicalBackupPlan BackupPlan `json:"periodical_backup_plan"`
}

type BandWidthInfo struct {
	BandWidth          int  `json:"bandwidth"`
	BeginTime          int  `json:"begin_time"`
	CurrentTime        int  `json:"current_time"`
	EndTime            int  `json:"end_time"`
	ExpandCount        int  `json:"expand_count"`
	ExpandEffectTime   int  `json:"expand_effect_time"`
	ExpandIntervalTime int  `json:"expand_interval_time"`
	MaxExpandCount     int  `json:"max_expand_count"`
	NextExpandTime     int  `json:"next_expand_time"`
	TaskRunning        bool `json:"task_running"`
}

type DomainNameInfo struct {
	SupportPublicResolve bool           `json:"support_public_resolve"`
	IsLatestRules        bool           `json:"is_latest_rules"`
	ZoneName             string         `json:"zone_name"`
	HistoryDomainNames   []DomainEntity `json:"history_domain_names"`
}

type DomainEntity struct {
	DomainName string `json:"domain_name"`
	IsReadonly bool   `json:"is_readonly"`
}
