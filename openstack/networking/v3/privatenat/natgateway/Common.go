package natgateway

// ########## COMMON RESPONSE STRUCTS ###########

type GatewayCommonResponse struct {
	// Specifies the response body for the private NAT gateway.
	Gateway PrivateNATGateway `json:"gateway"`
	// Request ID
	RequestID string `json:"request_id"`
}

type PrivateNATGateway struct {
	// Specifies the private NAT gateway ID.
	Id string `json:"id"`
	// Project ID
	ProjectId string `json:"project_id"`
	// Specifies the private NAT gateway name.
	// Only digits, letters, underscores (_), and hyphens (-) are allowed.
	Name string `json:"name"`
	// Provides supplementary information about the private NAT gateway.
	// The description can contain up to 255 characters and cannot contain angle brackets (<>).
	Description string `json:"description"`
	// Specifies the private NAT gateway specifications.
	// The value can be: Small, Medium, Large, Extra-large.
	// Default value: Small.
	Spec string `json:"spec"`
	// Specifies the private NAT gateway status.
	// The value can be:
	// ACTIVE: The private NAT gateway is running properly.
	// FROZEN: The private NAT gateway is frozen.
	Status string `json:"status"`
	// Specifies the time when the private NAT gateway was created. It is a UTC time in yyyy-mm-ddThh:mm:ssZ format.
	CreatedAt string `json:"created_at"`
	// Specifies the time when the private NAT gateway was updated. It is a UTC time in yyyy-mm-ddThh:mm:ssZ format.
	UpdatedAt string `json:"updated_at"`
	// Specifies the VPC where the private NAT gateway works.
	DownlinkVpcs []DownlinkVpc `json:"downlink_vpcs"`
	// Specifies the tag list.
	Tags []Tag `json:"tags"`
	// Specifies the ID of the enterprise project associated with the private NAT gateway.
	// Default value: 0.
	EnterpriseProjectID string `json:"enterprise_project_id"`
	// Specifies the maximum number of rules. Value range: 0-65535
	RuleMax int `json:"rule_max"`
	// Specifies the maximum number of transit IP addresses in a transit IP address pool. Value range: 1-100
	TransitIpPoolSizeMax int `json:"transit_ip_pool_size_max"`
}

type DownlinkVpc struct {
	// VPC ID
	VpcId string `json:"vpc_id"`
	// Specifies the ID of the subnet where the private NAT gateway works.
	VirSubnetID string `json:"virsubnet_id"`
	// Specifies the private IP address of the private NAT gateway.
	NgPortIPAddress string `json:"ngport_ip_address"`
}

type Tag struct {
	// Specifies the tag key. A key can contain up to 128 Unicode characters.
	// Key cannot be left blank.
	Key string `json:"key"`
	// Specifies the tag value. Each value can contain up to 255 Unicode characters.
	Value string `json:"value"`
}
