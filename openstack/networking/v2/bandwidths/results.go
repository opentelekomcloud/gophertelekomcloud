package bandwidths

import (
	"bytes"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type Bandwidth struct {
	// Specifies the bandwidth name. The value is a string of 1 to 64
	// characters that can contain letters, digits, underscores (_), and hyphens (-).
	Name string `json:"name"`

	// Specifies the bandwidth size. The value ranges from 1 Mbit/s to
	// 300 Mbit/s.
	Size int `json:"size"`

	// Specifies the bandwidth ID, which uniquely identifies the
	// bandwidth.
	ID string `json:"id"`

	// Specifies whether the bandwidth is shared or exclusive. The
	// value can be PER or WHOLE.
	ShareType string `json:"share_type"`

	// Specifies the elastic IP address of the bandwidth.  The
	// bandwidth, whose type is set to WHOLE, supports up to 20 elastic IP addresses. The
	// bandwidth, whose type is set to PER, supports only one elastic IP address.
	PublicIpInfo []PublicIpInfo `json:"publicip_info"`

	// Specifies the tenant ID of the user.
	TenantId string `json:"tenant_id"`

	// Specifies the bandwidth type.
	BandwidthType string `json:"bandwidth_type"`

	// Specifies the charging mode (by traffic or by bandwidth).
	ChargeMode string `json:"charge_mode"`

	// Specifies the billing information.
	BillingInfo string `json:"billing_info"`

	// Status
	Status string `json:"status"`

	// CreatedAt is the date when the Bandwidth was created.
	CreatedAt string `json:"created_at"`

	// UpdatedAt is the date when the last change was made to the Bandwidth.
	UpdatedAt string `json:"updated_at"`
}

type PublicIpInfo struct {
	// Specifies the tenant ID of the user.
	ID string `json:"publicip_id"`

	// Specifies the elastic IP address.
	Address string `json:"publicip_address"`

	// Specifies the elastic IP v6 address.
	AddressV6 string `json:"publicipv6_address"`

	// Specifies the elastic IP version.
	IPVersion int `json:"ip_version"`

	// Specifies the elastic IP address type.
	Type string `json:"publicip_type"`
}

type commonResult struct {
	golangsdk.Result
}

type CreateResult struct {
	commonResult
}

type UpdateResult struct {
	commonResult
}

type GetResult struct {
	commonResult
}

func (r commonResult) Extract() (*Bandwidth, error) {
	s := new(Bandwidth)
	err := r.ExtractIntoStructPtr(s, "bandwidth")
	if err != nil {
		return nil, err
	}
	return s, nil
}

type BandwidthPage struct {
	pagination.SinglePageBase
}

func (r BandwidthPage) IsEmpty() (bool, error) {
	is, err := ExtractBandwidths(r)
	return len(is) == 0, err
}

func ExtractBandwidths(r pagination.Page) ([]Bandwidth, error) {
	var s []Bandwidth

	err := extract.IntoSlicePtr(bytes.NewReader((r.(BandwidthPage)).Body), &s, "bandwidths")
	if err != nil {
		return nil, err
	}
	return s, err
}

type DeleteResult struct {
	golangsdk.ErrResult
}
