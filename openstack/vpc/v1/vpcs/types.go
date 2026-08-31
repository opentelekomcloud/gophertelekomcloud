package vpcs

type Route struct {
	NextHop         string `json:"nexthop,omitempty"`
	DestinationCIDR string `json:"destination,omitempty"`
}

type Vpc struct {
	ID                  string  `json:"id"`
	Name                string  `json:"name"`
	CIDR                string  `json:"cidr"`
	Description         string  `json:"description"`
	Status              string  `json:"status"`
	Routes              []Route `json:"routes"`
	EnterpriseProjectID string  `json:"enterprise_project_id"`
	EnableSharedSnat    bool    `json:"enable_shared_snat"`
	TenantId            string  `json:"tenant_id"`
	CreatedAt           string  `json:"created_at"`
	UpdatedAt           string  `json:"updated_at"`
}
