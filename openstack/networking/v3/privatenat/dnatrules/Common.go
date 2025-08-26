package dnatrules

// ############# COMMON RESPONSE STRUCTS #############

type DnatCommonResponse struct {
	// Specifies the response body of the DNAT rule
	DnatRule PrivateDnat `json:"dnat_rule"`
	// Specifies the request ID
	RequestId string `json:"request_id"`
}

type PrivateDnat struct {
	// Specifies the DNAT rule ID
	Id string `json:"id"`
	// Specifies the project ID
	ProjectId string `json:"project_id"`
	// Provides supplementary information about the DNAT rule.
	// The description can contain up to 255 characters and cannot contain angle brackets (<>).
	Description string `json:"description"`
	// Specifies the ID of the transit IP address
	TransitIpId string `json:"transit_ip_id"`
	// Specifies the private NAT gateway ID
	GatewayId string `json:"gateway_id"`
	// Specifies the network interface ID.
	// Network interfaces of a compute instance, load balancer (v2 or v3), or virtual IP address are supported.
	NetworkInterfaceId string `json:"network_interface_id"`
	// Specifies the backend resource type of the DNAT rule.
	// Enumeration values: COMPUTE, VIP, ELB, ELBv3, CUSTOMIZE
	Type string `json:"type"`
	// Specifies the protocol type.
	// Enumeration values: tcp, udp, any
	Protocol string `json:"protocol"`
	// Specifies the port IP address that the NAT gateway uses.
	// The resource can be a compute instance, load balancer (v2 or v3), or virtual IP address.
	PrivateIpAddress string `json:"private_ip_address"`
	// Specifies the port number of the resource (compute instance, load balancer, or virtual IP address).
	// Value range: 0-65535, Default: 0
	InternalServicePort string `json:"internal_service_port"`
	// Specifies the port number of the transit IP address.
	// Value range: 0-65535, Default: 0
	TransitServicePort string `json:"transit_service_port"`
	// Specifies the ID of the enterprise project that is associated with the DNAT rule when created
	EnterpriseProjectId string `json:"enterprise_project_id"`
	// Specifies the time when the DNAT rule was created (UTC, yyyy-mm-ddThh:mm:ssZ)
	CreatedAt string `json:"created_at"`
	// Specifies the time when the DNAT rule was updated (UTC, yyyy-mm-ddThh:mm:ssZ)
	UpdatedAt string `json:"updated_at"`
	// Specifies the DNAT rule status of a private NAT gateway.
	// Enumeration values: ACTIVE, FROZEN
	Status string `json:"status"`
}
