package bandwidths

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// RemovePublicIPInfo describes an EIP to be removed from a shared bandwidth.
type RemovePublicIPInfo struct {
	// Specifies the ID of the EIP that uses the bandwidth.
	PublicipId string `json:"publicip_id" required:"true"`
}

type RemoveEipOpts struct {
	// Specifies information about the EIP(s) to be removed from the shared
	// bandwidth.
	PublicipInfo []RemovePublicIPInfo `json:"publicip_info" required:"true"`

	// After an EIP is removed from a shared bandwidth, a dedicated
	// bandwidth is allocated to the EIP. Specifies whether that dedicated
	// bandwidth is billed by traffic or by bandwidth. The value can be
	// "bandwidth" or "traffic".
	ChargeMode string `json:"charge_mode" required:"true"`

	// After an EIP is removed from a shared bandwidth, a dedicated
	// bandwidth is allocated to the EIP. Specifies the size (Mbit/s) of
	// that dedicated bandwidth.
	Size int `json:"size" required:"true"`
}

// RemoveEip removes one or more EIPs from a shared bandwidth.
func RemoveEip(client *golangsdk.ServiceClient, id string, opts RemoveEipOpts) error {
	b, err := build.RequestBody(opts, "bandwidth")
	if err != nil {
		return err
	}
	_, err = client.Post(client.ServiceURL(client.ProjectID, "bandwidths", id, "remove"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
