package peerings

type VpcInfo struct {
	VpcID    string `json:"vpc_id" required:"true"`
	TenantID string `json:"tenant_id,omitempty"`
}

type Peering struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Status         string  `json:"status"`
	RequestVpcInfo VpcInfo `json:"request_vpc_info"`
	AcceptVpcInfo  VpcInfo `json:"accept_vpc_info"`
	Description    string  `json:"description"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}
