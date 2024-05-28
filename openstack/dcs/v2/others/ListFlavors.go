package others

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListFlavorOpts struct {
	CacheMode     string `q:"cache_mode"`
	Engine        string `q:"engine"`
	EngineVersion string `q:"engine_version"`
	Capacity      string `q:"capacity"`
	SpecCode      string `q:"spec_code"`
	CPUType       string `q:"cpu_type"`
}

func ListFlavors(client *golangsdk.ServiceClient, opts ListFlavorOpts) ([]Product, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("flavors").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Product
	err = extract.IntoSlicePtr(raw.Body, &res, "flavors")
	return res, err
}

type Product struct {
	SpecCode              string           `json:"spec_code"`
	CloudServiceTypeCode  string           `json:"cloud_service_type_code"`
	CloudResourceTypeCode string           `json:"cloud_resource_type_code"`
	CacheMode             string           `json:"cache_mode"`
	Engine                string           `json:"engine"`
	EngineVersion         string           `json:"engine_version"`
	ProductType           string           `json:"product_type"`
	CpuType               string           `json:"cpu_type"`
	StorageType           string           `json:"storage_type"`
	Capacity              []string         `json:"capacity"`
	BillingMode           []string         `json:"billing_mode"`
	TenantIpCount         int              `json:"tenant_ip_count"`
	PricingType           string           `json:"pricing_type"`
	IsDec                 bool             `json:"is_dec"`
	Attrs                 []Attrs          `json:"attrs"`
	FlavorsAvailableZones []FlavorAzObject `json:"flavors_available_zones"`
	ReplicaCount          int              `json:"replica_count"`
}

type Attrs struct {
	Capacity string `json:"capacity"`
	Name     string `json:"name"`
	Value    string `json:"value"`
}

type FlavorAzObject struct {
	Capacity       string   `json:"capacity"`
	Unit           string   `json:"unit"`
	AvailableZones []string `json:"available_zones"`
	AzCodes        []string `json:"az_codes"`
}
