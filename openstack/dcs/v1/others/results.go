package others

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// GetResponse response
type GetResponse struct {
	Products []Product `json:"products"`
}

// Product for dcs

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
