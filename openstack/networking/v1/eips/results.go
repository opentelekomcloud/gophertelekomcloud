package eips

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ApplyResult is a struct which represents the result of apply public ip
type ApplyResult struct {
	golangsdk.Result
}

func (r ApplyResult) Extract() (*PublicIp, error) {
	s := new(PublicIp)
	err := r.ExtractIntoStructPtr(s, "publicip")
	return s, err
}

// PublicIp is a struct that represents a public ip
type PublicIp struct {
	ID                 string `json:"id"`
	Status             string `json:"status"`
	Type               string `json:"type"`
	PublicAddress      string `json:"public_ip_address"`
	PrivateAddress     string `json:"private_ip_address"`
	PortID             string `json:"port_id"`
	TenantID           string `json:"tenant_id"`
	CreateTime         string `json:"create_time"`
	BandwidthID        string `json:"bandwidth_id"`
	BandwidthSize      int    `json:"bandwidth_size"`
	BandwidthShareType string `json:"bandwidth_share_type"`
	IpVersion          int    `json:"ip_version"`
	Name               string `json:"alias"`
}

// GetResult is a return struct of get method
type GetResult struct {
	golangsdk.Result
}

func (r GetResult) Extract() (*PublicIp, error) {
	s := new(PublicIp)
	err := r.ExtractIntoStructPtr(s, "publicip")
	return s, err
}

// DeleteResult is a struct of delete result
type DeleteResult struct {
	golangsdk.ErrResult
}

// UpdateResult is a struct which contains the result of update method
type UpdateResult struct {
	golangsdk.Result
}

func (r UpdateResult) Extract() (*PublicIp, error) {
	s := new(PublicIp)
	err := r.ExtractIntoStructPtr(s, "publicip")
	return s, err
}

// EipPage is a single page of Flavor results.
type EipPage struct {
	pagination.MarkerPageBase
}

// IsEmpty returns true if a page contains no results.
func (r EipPage) IsEmpty() (bool, error) {
	eips, err := ExtractEips(r)
	return len(eips) == 0, err
}

// LastMarker returns the last Eip ID in a ListResult.
func (r EipPage) LastMarker() (string, error) {
	eips, err := ExtractEips(r)
	if err != nil {
		return "", err
	}
	if len(eips) == 0 {
		return "", nil
	}
	return eips[len(eips)-1].ID, nil
}

// ExtractEips extracts and returns Public IPs. It is used while iterating over a public ips.
func ExtractEips(r pagination.Page) ([]PublicIp, error) {
	var s []PublicIp
	err := (r.(EipPage)).ExtractIntoSlicePtr(&s, "publicips")
	return s, err
}
