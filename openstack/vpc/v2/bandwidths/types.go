package bandwidths

// PublicIPInfo describes an EIP that uses a shared bandwidth.
type PublicIPInfo struct {
	// Specifies the ID of the EIP that uses the bandwidth.
	PublicipId string `json:"publicip_id"`

	// Specifies the obtained EIP if only IPv4 EIPs are available.
	PublicipAddress string `json:"publicip_address"`

	// Specifies the obtained EIP if IPv6 EIPs are available. This parameter
	// does not exist if only IPv4 EIPs are available.
	Publicipv6Address string `json:"publicipv6_address"`

	// Specifies the IP address version. Possible values are 4 (IPv4) and 6
	// (IPv6).
	IPVersion int `json:"ip_version"`

	// Specifies the EIP type.
	PublicipType string `json:"publicip_type"`
}

// Bandwidth describes a shared bandwidth.
type Bandwidth struct {
	// Specifies the bandwidth name. The value can contain 1 to 64
	// characters, including letters, digits, underscores (_), hyphens (-),
	// and periods (.).
	Name string `json:"name"`

	// Specifies the bandwidth size in Mbit/s.
	Size int `json:"size"`

	// Specifies the bandwidth ID, which uniquely identifies the bandwidth.
	ID string `json:"id"`

	// Specifies whether the bandwidth is shared or dedicated. The value can
	// be PER (dedicated) or WHOLE (shared).
	ShareType string `json:"share_type"`

	// Specifies information about the EIP(s) that use the bandwidth. A
	// bandwidth whose type is WHOLE can be used by multiple EIPs, a
	// bandwidth whose type is PER can be used by only one EIP.
	PublicipInfo []PublicIPInfo `json:"publicip_info"`

	// Specifies the project ID.
	TenantId string `json:"tenant_id"`

	// Specifies the bandwidth type. The default value for a shared
	// bandwidth is "share".
	BandwidthType string `json:"bandwidth_type"`

	// Specifies that the bandwidth is billed by bandwidth. The value can be
	// "traffic".
	ChargeMode string `json:"charge_mode"`

	// Specifies the billing information. If specified, the bandwidth is in
	// yearly/monthly billing mode.
	BillingInfo string `json:"billing_info"`

	// Specifies the enterprise project ID.
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// Specifies the bandwidth status. Possible values are FREEZED (frozen)
	// and NORMAL (normal).
	Status string `json:"status"`

	// Specifies the time (UTC) when the bandwidth was created.
	CreatedAt string `json:"created_at"`

	// Specifies the time (UTC) when the bandwidth was updated.
	UpdatedAt string `json:"updated_at"`

	// Specifies whether the bandwidth is in a central site or an edge site.
	PublicBorderGroup string `json:"public_border_group"`
}
