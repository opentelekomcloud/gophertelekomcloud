package schedulerstats

import (
	"encoding/json"
	"math"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// ID of the tenant to look up storage pools for.
	TenantID string `q:"tenant_id"`
	// Whether to list extended details.
	Detail bool `q:"detail"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]StoragePool, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("scheduler-stats", "get_pools").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []StoragePool
	err = extract.IntoSlicePtr(raw.Body, &res, "pools")
	return res, err
}

type StoragePool struct {
	Name         string       `json:"name"`
	Capabilities Capabilities `json:"capabilities"`
}

type Capabilities struct {
	// The following fields should be present in all storage drivers.
	DriverVersion     string  `json:"driver_version"`
	FreeCapacityGB    float64 `json:"-"`
	StorageProtocol   string  `json:"storage_protocol"`
	TotalCapacityGB   float64 `json:"-"`
	VendorName        string  `json:"vendor_name"`
	VolumeBackendName string  `json:"volume_backend_name"`
	// The following fields are optional and may have empty values depending
	// on the storage driver in use.
	ReservedPercentage       int64   `json:"reserved_percentage"`
	LocationInfo             string  `json:"location_info"`
	QoSSupport               bool    `json:"QoS_support"`
	ProvisionedCapacityGB    float64 `json:"provisioned_capacity_gb"`
	MaxOverSubscriptionRatio string  `json:"max_over_subscription_ratio"`
	ThinProvisioningSupport  bool    `json:"thin_provisioning_support"`
	ThickProvisioningSupport bool    `json:"thick_provisioning_support"`
	TotalVolumes             int64   `json:"total_volumes"`
	FilterFunction           string  `json:"filter_function"`
	GoodnessFuction          string  `json:"goodness_function"`
	Mutliattach              bool    `json:"multiattach"`
	SparseCopyVolume         bool    `json:"sparse_copy_volume"`
}

func (r *Capabilities) UnmarshalJSON(b []byte) error {
	type tmp Capabilities
	var s struct {
		tmp
		FreeCapacityGB  interface{} `json:"free_capacity_gb"`
		TotalCapacityGB interface{} `json:"total_capacity_gb"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Capabilities(s.tmp)

	// Generic function to parse a capacity value which may be a numeric
	// value, "unknown", or "infinite"
	parseCapacity := func(capacity interface{}) float64 {
		if capacity != nil {
			switch capacity := capacity.(type) {
			case float64:
				return capacity
			case string:
				if capacity == "infinite" {
					return math.Inf(1)
				}
			}
		}
		return 0.0
	}

	r.FreeCapacityGB = parseCapacity(s.FreeCapacityGB)
	r.TotalCapacityGB = parseCapacity(s.TotalCapacityGB)

	return nil
}
