package others

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// GetResponse response
type GetResponse struct {
	Products []Product `json:"products"`
}

// Product for dcs
type Product struct {
	Price        float64 `json:"price"`
	Currency     string  `json:"currency"`
	ProductID    string  `json:"product_id"`
	SpecCode     string  `json:"spec_code"`
	SpecDetails  string  `json:"spec_details"`
	ChargingType string  `json:"charging_type"`
	SpecDetails2 string  `json:"spec_details2"`
	ProdType     string  `json:"prod_type"`
}

// GetResult contains the body of getting detailed
type GetResult struct {
	golangsdk.Result
}

// Extract from GetResult
func (r GetResult) Extract() (*GetResponse, error) {
	var s GetResponse
	err := r.Result.ExtractInto(&s)
	return &s, err
}

// GetResponse response
type GetResponse1 struct {
	RegionID       string          `json:"regionId"`
	AvailableZones []AvailableZone `json:"available_zones"`
}

// AvailableZone for dcs
type AvailableZone struct {
	ID                   string `json:"id"`
	Code                 string `json:"code"`
	Name                 string `json:"name"`
	Port                 string `json:"port"`
	ResourceAvailability string `json:"resource_availability"`
}

// GetResult contains the body of getting detailed
type GetResult1 struct {
	golangsdk.Result
}

// Extract from GetResult
func (r GetResult1) Extract() (*GetResponse1, error) {
	var s GetResponse1
	err := r.Result.ExtractInto(&s)
	return &s, err
}

// GetResponse response
type GetResponse2 struct {
	MaintainWindows []MaintainWindow `json:"maintain_windows"`
}

// MaintainWindow for dcs
type MaintainWindow struct {
	ID      int    `json:"seq"`
	Begin   string `json:"begin"`
	End     string `json:"end"`
	Default bool   `json:"default"`
}

// GetResult contains the body of getting detailed
type GetResult3 struct {
	golangsdk.Result
}

// Extract from GetResult
func (r GetResult3) Extract() (*GetResponse2, error) {
	var s GetResponse2
	err := r.Result.ExtractInto(&s)
	return &s, err
}
