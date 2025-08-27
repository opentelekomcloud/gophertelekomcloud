package transitip

// ########## COMMON RESPONSE STRUCTS ###########

type TransitIPCommonResponse struct {
	// Specifies the response body of the transit IP address.
	TransitIp TransitIP `json:"transit_ip"`
	// Request ID
	RequestID string `json:"request_id"`
}

type TransitIP struct {
	// Specifies the ID of the transit IP address.
	Id string `json:"id"`
	// Project ID
	ProjectId string `json:"project_id"`
	// Specifies the network interface ID of the transit IP address.
	NetworkInterfaceId string `json:"network_interface_id"`
	// Specifies the transit IP address.
	IpAddress string `json:"ip_address"`
	// Specifies the time when the transit IP address was assigned. It is a UTC time in yyyy-mm-ddThh:mm:ssZ format.
	CreatedAt string `json:"created_at"`
	// Specifies the time when the transit IP address was updated. It is a UTC time in yyyy-mm-ddThh:mm:ssZ format.
	UpdatedAt string `json:"updated_at"`
	// Specifies the subnet ID of the current tenant.
	VirSubnetID string `json:"virsubnet_id"`
	// Specifies the tag list.
	Tags []Tag `json:"tags"`
	// Specifies the ID of the private NAT gateway associated with the transit IP address.
	GatewayId string `json:"gateway_id"`
	// Specifies the ID of the enterprise project that is associated with the transit IP address when the transit IP address is being assigned.
	EnterpriseProjectID string `json:"enterprise_project_id"`
	// Specifies the transit IP address status.
	// The value can be:
	// ACTIVE: The transit IP address is running properly.
	// FROZEN: The transit IP address is frozen.
	Status string `json:"status"`
}

type Tag struct {
	// Specifies the tag key. A key can contain up to 128 Unicode characters.
	// Key cannot be left blank.
	Key string `json:"key"`
	// Specifies the tag value. Each value can contain up to 255 Unicode characters.
	Value string `json:"value"`
}
