package bandwidths

// PublicIpinfo describes an EIP that uses a bandwidth.
type PublicIpinfo struct {
	// Specifies the ID of the EIP that uses the bandwidth.
	PublicipId string `json:"publicip_id"`

	// Specifies the obtained EIP if only IPv4 EIPs are available.
	PublicipAddress string `json:"publicip_address"`

	// Publicipv6Address is not documented for this operation in the reviewed
	// OTC documentation.
	Publicipv6Address string `json:"publicipv6_address"`

	// Specifies the IP address version. Possible values are 4 (IPv4) and 6
	// (IPv6).
	IPVersion int `json:"ip_version"`

	// Specifies the EIP type.
	PublicipType string `json:"publicip_type"`
}

// BandWidth describes a dedicated or shared bandwidth.
type BandWidth struct {
	// Specifies the bandwidth name. The value is a string of 1 to 64
	// characters that can contain letters, digits, underscores (_), and
	// hyphens (-).
	Name string `json:"name"`

	// Specifies the bandwidth size in Mbit/s.
	Size int `json:"size"`

	// Specifies the bandwidth ID, which uniquely identifies the bandwidth.
	ID string `json:"id"`

	// Specifies whether the bandwidth is shared or exclusive. The value can
	// be PER (dedicated) or WHOLE (shared).
	ShareType string `json:"share_type"`

	// Specifies information about the EIP(s) that use the bandwidth. A
	// bandwidth whose type is set to WHOLE supports up to 20 EIPs. A
	// bandwidth whose type is set to PER supports only one EIP.
	PublicipInfo []PublicIpinfo `json:"publicip_info"`

	// Specifies the project ID of the user.
	TenantId string `json:"tenant_id"`

	// Specifies the bandwidth type.
	BandwidthType string `json:"bandwidth_type"`

	// Specifies the charging mode (by traffic or by bandwidth).
	ChargeMode string `json:"charge_mode"`

	// Specifies the billing information.
	BillingInfo string `json:"billing_info"`

	// Specifies the enterprise project ID.
	EnterpriseProjectID string `json:"enterprise_project_id"`

	// Specifies the bandwidth status.
	Status string `json:"status"`

	// Specifies the time (UTC) when the bandwidth was created.
	CreatedAt string `json:"created_at"`

	// Specifies the time (UTC) when the bandwidth was updated.
	UpdatedAt string `json:"updated_at"`

	// Specifies whether the bandwidth is in a central site or an edge site.
	PublicBorderGroup string `json:"public_border_group"`
}
