package snatrules

// ############# COMMON RESPONSE STRUCTS #############

type SnatCommonResponse struct {
	// Specifies the response body of the SNAT rule
	SnatRule PrivateSnat `json:"snat_rule"`
	// Specifies the request ID
	RequestId string `json:"request_id"`
}

type PrivateSnat struct {
	// Specifies the SNAT rule ID
	Id string `json:"id"`
	// Specifies the project ID
	ProjectId string `json:"project_id"`
	// Specifies the private NAT gateway ID
	GatewayId string `json:"gateway_id"`
	// Specifies the CIDR block that matches the SNAT rule.
	// Constraint: Either this parameter or virsubnet_id must be specified.
	Cidr string `json:"cidr"`
	// Specifies the ID of the subnet that matches the SNAT rule.
	// Constraint: Either this parameter or cidr must be specified.
	VirSubnetId string `json:"virsubnet_id"`
	// Provides supplementary information about the SNAT rule.
	// The description can contain up to 255 characters and cannot contain angle brackets (<>).
	Description string `json:"description"`
	// Specifies the list of details of associated transit IP addresses.
	TransitIpAssociations []TransitIPAssociation `json:"transit_ip_associations"`
	// Specifies the time when the SNAT rule was created (UTC, yyyy-mm-ddThh:mm:ssZ)
	CreatedAt string `json:"created_at"`
	// Specifies the time when the SNAT rule was updated (UTC, yyyy-mm-ddThh:mm:ssZ)
	UpdatedAt string `json:"updated_at"`
	// Specifies the ID of the enterprise project that is associated with the SNAT rule when created
	EnterpriseProjectId string `json:"enterprise_project_id"`
	// Specifies the SNAT rule status of a private NAT gateway.
	// Enumeration values: ACTIVE, FROZEN
	Status string `json:"status"`
}

type TransitIPAssociation struct {
	// Specifies the ID of the transit IP address
	TransitIpId string `json:"transit_ip_id"`
	// Specifies the transit IP address
	TransitIpAddress string `json:"transit_ip_address"`
}
