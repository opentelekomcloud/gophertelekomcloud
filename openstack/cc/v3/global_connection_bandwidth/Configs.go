package global_connection_bandwidth

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetConfigs queries the tenant configuration of a global connection bandwidth.
func GetConfigs(client *golangsdk.ServiceClient) (*ConfigsResp, error) {
	raw, err := client.Get(client.ServiceURL(client.DomainID, "gcb", "configs"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ConfigsResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ConfigsResp struct {
	// Request ID.
	RequestId string `json:"request_id"`
	// Dynamic configuration items for purchasing global connection bandwidth.
	Configs Configs `json:"configs"`
}

// Configs describes the tenant configuration of global connection bandwidth.
type Configs struct {
	// Capacity of global connection bandwidths by billing option.
	SizeRange []SizeRange `json:"size_range"`
	// List of supported billing options.
	ChargeMode []string `json:"charge_mode"`
	// Instance type.
	Services []string `json:"services"`
	// Bandwidth type.
	GcbType []string `json:"gcb_type"`
	// Percentage of minimum bandwidth in enhanced 95th percentile billing.
	Ratio95PeakPlus int `json:"ratio_95peak_plus"`
	// Percentage of minimum bandwidth in standard 95th percentile billing.
	Ratio95PeakGuar int `json:"ratio_95peak_guar"`
	// Whether a cross-border permit is approved.
	Crossborder bool `json:"crossborder"`
	// Quota information.
	Quotas []ConfigQuota `json:"quotas"`
	// Line grade.
	SlaLevel []string `json:"sla_level"`
	// Maximum number of instances that are allowed to use a shared bandwidth.
	BindLimit int `json:"bind_limit"`
	// Whether to enable the geographic region bandwidth.
	EnableAreaBandwidth bool `json:"enable_area_bandwidth"`
	// Whether standard 95th percentile bandwidth billing can be changed to billing by bandwidth.
	EnableChange95 bool `json:"enable_change_95"`
	// Whether multiple line specifications are supported.
	EnableSpecCode bool `json:"enable_spec_code"`
	// Whether to enable Cloud Eye monitoring.
	CesEnabled bool `json:"ces_enabled"`
}

// SizeRange describes the bandwidth capacity range for a billing option.
type SizeRange struct {
	// Billing option. Value options: bwd, 95, 95avr.
	Type string `json:"type"`
	// Minimum bandwidth in Mbit/s.
	Min int `json:"min"`
	// Maximum bandwidth in Mbit/s.
	Max int `json:"max"`
}

// ConfigQuota describes a quota in the tenant configuration.
type ConfigQuota struct {
	// Quotas.
	Quota int `json:"quota"`
	// Used quotas.
	Used int `json:"used"`
	// Quota type. Value options: gcb.size, gcb.count.
	Type string `json:"type"`
}
