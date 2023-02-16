package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// InstanceCreate response
type InstanceCreate struct {
	InstanceID string `json:"instance_id"`
}

type ListResponse struct {
	Instances  []Instance `json:"instances"`
	TotalCount int        `json:"instance_num"`
}

// Instance response
type Instance struct {
	Name                       string             `json:"name"`
	Description                string             `json:"description"`
	Engine                     string             `json:"engine"`
	EngineVersion              string             `json:"engine_version"`
	Specification              string             `json:"specification"`
	StorageSpace               int                `json:"storage_space"`
	PartitionNum               string             `json:"partition_num"`
	BrokerNum                  int                `json:"broker_num"`
	NodeNum                    int                `json:"node_num"`
	UsedStorageSpace           int                `json:"used_storage_space"`
	ConnectAddress             string             `json:"connect_address"`
	Port                       int                `json:"port"`
	Status                     string             `json:"status"`
	InstanceID                 string             `json:"instance_id"`
	ResourceSpecCode           string             `json:"resource_spec_code"`
	ChargingMode               int                `json:"charging_mode"`
	VPCID                      string             `json:"vpc_id"`
	VPCName                    string             `json:"vpc_name"`
	CreatedAt                  string             `json:"created_at"`
	UserID                     string             `json:"user_id"`
	UserName                   string             `json:"user_name"`
	MaintainBegin              string             `json:"maintain_begin"`
	MaintainEnd                string             `json:"maintain_end"`
	EnablePublicIP             bool               `json:"enable_publicip"`
	ManagementConnectAddress   string             `json:"management_connect_address"`
	SslEnable                  bool               `json:"ssl_enable"`
	Type                       string             `json:"type"`
	ProductID                  string             `json:"product_id"`
	SecurityGroupID            string             `json:"security_group_id"`
	SecurityGroupName          string             `json:"security_group_name"`
	SubnetID                   string             `json:"subnet_id"`
	SubnetName                 string             `json:"subnet_name"`
	SubnetCIDR                 string             `json:"subnet_cidr"`
	AvailableZones             []string           `json:"available_zones"`
	TotalStorageSpace          int                `json:"total_storage_space"`
	PublicConnectionAddress    string             `json:"public_connect_address"`
	StorageResourceID          string             `json:"storage_resource_id"`
	StorageSpecCode            string             `json:"storage_spec_code"`
	ServiceType                string             `json:"service_type"`
	StorageType                string             `json:"storage_type"`
	RetentionPolicy            string             `json:"retention_policy"`
	KafkaPublicStatus          string             `json:"kafka_public_status"`
	PublicBandWidth            int                `json:"public_bandwidth"`
	CrossVpcInfo               string             `json:"cross_vpc_info"`
	MessageQueryInstEnable     bool               `json:"message_query_inst_enable"`
	SupportFeatures            string             `json:"support_features"`
	DiskEncrypted              bool               `json:"disk_encrypted"`
	DiskEncryptedKey           string             `json:"disk_encrypted_key"`
	KafkaPrivateConnectAddress string             `json:"kafka_private_connect_address"`
	PublicAccessEnabled        string             `json:"public_access_enabled"`
	AccessUser                 string             `json:"access_user"`
	Tags                       []tags.ResourceTag `json:"tags"`
}

// CrossVpc is the structure that represents the API response of 'UpdateCrossVpc' method.
type CrossVpc struct {
	// The result of cross-VPC access modification.
	Success bool `json:"success"`
	// The result list of broker cross-VPC access modification.
	Connections []Connection `json:"results"`
}

// Connection is the structure that represents the detail of the cross-VPC access.
type Connection struct {
	// advertised.listeners IP/domain name.
	AdvertisedIp string `json:"advertised_ip"`
	// The status of broker cross-VPC access modification.
	Success bool `json:"success"`
	// Listeners IP.
	ListenersIp string `json:"ip"`
}
