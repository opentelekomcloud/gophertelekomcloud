package loadbalancers

type BandwidthRef struct {
	// Share Bandwidth ID
	ID string `json:"id" required:"true"`
}

type PublicIp struct {
	// IP Version.
	IpVersion int `json:"ip_version,omitempty"`

	// Network Type
	NetworkType string `json:"network_type" required:"true"`

	// Billing Info.
	BillingInfo string `json:"billing_info,omitempty"`

	// Description.
	Description string `json:"description,omitempty"`

	// Bandwidth
	Bandwidth Bandwidth `json:"bandwidth" required:"true"`
}

type Bandwidth struct {
	// Name
	Name string `json:"name" required:"true"`

	// Size
	Size int `json:"size" required:"true"`

	// Charge Mode
	ChargeMode string `json:"charge_mode" required:"true"`

	// Share Type
	ShareType string `json:"share_type" required:"true"`
}
