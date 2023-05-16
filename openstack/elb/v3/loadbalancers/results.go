package loadbalancers

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type EipInfo struct {
	// Eip ID
	EipID string `json:"eip_id"`
	// Eip Address
	EipAddress string `json:"eip_address"`
	// Eip Address
	IpVersion int `json:"ip_version"`
}

type PublicIpInfo struct {
	// Public IP ID
	PublicIpID string `json:"publicip_id"`
	// Public IP Address
	PublicIpAddress string `json:"publicip_address"`
	// IP Version
	IpVersion int `json:"ip_version"`
}

type commonResult struct {
	golangsdk.Result
}

// Extract is a function that accepts a result and extracts a loadbalancer.
func (r commonResult) Extract() (*LoadBalancer, error) {
	s := new(LoadBalancer)
	err := r.ExtractIntoStructPtr(s, "loadbalancer")
	if err != nil {
		return nil, err
	}
	return s, nil
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a LoadBalancer.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a LoadBalancer.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a LoadBalancer.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}
