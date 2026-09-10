package securitygrouprules

type SecurityGroupRule struct {
	ID                   string `json:"id"`
	Description          string `json:"description"`
	SecurityGroupID      string `json:"security_group_id"`
	Direction            string `json:"direction"`
	EtherType            string `json:"ethertype"`
	Protocol             string `json:"protocol"`
	PortRangeMin         *int   `json:"port_range_min"`
	PortRangeMax         *int   `json:"port_range_max"`
	RemoteIPPrefix       string `json:"remote_ip_prefix"`
	RemoteGroupID        string `json:"remote_group_id"`
	RemoteAddressGroupID string `json:"remote_address_group_id"`
	TenantID             string `json:"tenant_id"`
}
