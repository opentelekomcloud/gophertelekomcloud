package management

type Flavor struct {
	// Firewall version. Its value can only be 1 (professional edition).
	Version int `json:"version"`
	// Number of EIPs.
	EipCount int `json:"eip_count"`
	// Number of VPCs.
	VpcCount int `json:"vpc_count"`
	// Bandwidth, in Mbit/s.
	Bandwidth int `json:"bandwidth"`
	// Log storage, in bytes.
	LogStorage int `json:"log_storage"`
	// Default firewall bandwidth, in Mbit/s.
	// The value is 10 for the standard edition, 50 for the professional edition,
	// and 200 for the pay-per-use professional edition.
	DefaultBandwidth int `json:"default_bandwidth"`
	// Default number of EIPs.
	// The value is 20 for the standard edition, 50 for the professional edition,
	// and 1,000 for the pay-per-use professional edition.
	DefaultEipCount int `json:"default_eip_count"`
	// Default log storage, in bytes. The default value is 0.
	DefaultLogStorage int `json:"default_log_storage"`
	// Default number of VPCs.
	// The value is 0 for the standard edition, 2 for the professional edition,
	// and 5 for the pay-per-use professional edition.
	DefaultVpcCount int `json:"default_vpc_count"`
}
