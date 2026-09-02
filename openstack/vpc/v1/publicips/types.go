package publicips

type Profile struct {
	OrderID   string `json:"order_id"`
	ProductID string `json:"product_id"`
	RegionID  string `json:"region_id"`
	UserID    string `json:"user_id"`
}

type PublicIP struct {
	ID              string  `json:"id"`
	Status          string  `json:"status"`
	Profile         Profile `json:"profile"`
	Type            string  `json:"type"`
	PublicIpAddress string  `json:"public_ip_address"`
	// PublicIpV6Address is not documented in the reviewed OTC documentation.
	PublicIpV6Address        string   `json:"public_ipv6_address"`
	IPVersion                int      `json:"ip_version"`
	PrivateIpAddress         string   `json:"private_ip_address"`
	PortId                   string   `json:"port_id"`
	TenantId                 string   `json:"tenant_id"`
	CreateTime               string   `json:"create_time"`
	BandwidthId              string   `json:"bandwidth_id"`
	BandwidthSize            int      `json:"bandwidth_size"`
	BandwidthShareType       string   `json:"bandwidth_share_type"`
	BandwidthName            string   `json:"bandwidth_name"`
	Alias                    string   `json:"alias"`
	EnterpriseProjectId      string   `json:"enterprise_project_id"`
	PublicBorderGroup        string   `json:"public_border_group"`
	AllowShareBandwidthTypes []string `json:"allow_share_bandwidth_types"`
	Tags                     []string `json:"tags"`
}

type PublicIPCreateResp = PublicIP

type PublicIPUpdateResp = PublicIP
