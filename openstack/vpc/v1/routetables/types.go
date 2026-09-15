package routetables

type RouteTable struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Default     bool     `json:"default"`
	Routes      []Route  `json:"routes"`
	Subnets     []Subnet `json:"subnets"`
	TenantID    string   `json:"tenant_id"`
	VpcID       string   `json:"vpc_id"`
	Description string   `json:"description"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

type Route struct {
	Type            string `json:"type"`
	DestinationCIDR string `json:"destination"`
	NextHop         string `json:"nexthop"`
	Description     string `json:"description"`
}

type Subnet struct {
	ID string `json:"id"`
}

type RouteOpts struct {
	Type        string  `json:"type" required:"true"`
	Destination string  `json:"destination" required:"true"`
	NextHop     string  `json:"nexthop" required:"true"`
	Description *string `json:"description,omitempty"`
}
