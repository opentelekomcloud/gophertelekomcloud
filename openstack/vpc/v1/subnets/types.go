package subnets

type ExtraDHCPOpt struct {
	OptName  string `json:"opt_name"`
	OptValue string `json:"opt_value,omitempty"`
}

type Subnet struct {
	ID               string         `json:"id"`
	Name             string         `json:"name"`
	Description      string         `json:"description"`
	CIDR             string         `json:"cidr"`
	GatewayIP        string         `json:"gateway_ip"`
	EnableIpv6       bool           `json:"ipv6_enable"`
	CidrV6           string         `json:"cidr_v6"`
	GatewayIpV6      string         `json:"gateway_ip_v6"`
	EnableDHCP       bool           `json:"dhcp_enable"`
	PrimaryDNS       string         `json:"primary_dns"`
	SecondaryDNS     string         `json:"secondary_dns"`
	DNSList          []string       `json:"dnsList"`
	AvailabilityZone string         `json:"availability_zone"`
	VpcID            string         `json:"vpc_id"`
	Status           string         `json:"status"`
	NetworkID        string         `json:"neutron_network_id"`
	SubnetID         string         `json:"neutron_subnet_id"`
	SubnetIDV6       string         `json:"neutron_subnet_id_v6"`
	ExtraDHCPOpts    []ExtraDHCPOpt `json:"extra_dhcp_opts"`
	Scope            string         `json:"scope"`
	TenantID         string         `json:"tenant_id"`
	CreatedAt        string         `json:"created_at"`
	UpdatedAt        string         `json:"updated_at"`
}
