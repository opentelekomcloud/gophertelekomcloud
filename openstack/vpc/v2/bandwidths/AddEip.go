package bandwidths

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// InsertPublicIPInfo describes an EIP to be added to a shared bandwidth.
type InsertPublicIPInfo struct {
	// Specifies the ID of the EIP that uses the bandwidth.
	PublicipId string `json:"publicip_id" required:"true"`

	// Specifies the EIP type.
	PublicipType string `json:"publicip_type,omitempty"`
}

type AddEipOpts struct {
	// Specifies information about the EIP(s) to be added to the shared
	// bandwidth. A shared bandwidth can be used by up to 20 EIPs by
	// default.
	PublicipInfo []InsertPublicIPInfo `json:"publicip_info" required:"true"`
}

// AddEip adds one or more EIPs to a shared bandwidth.
func AddEip(client *golangsdk.ServiceClient, id string, opts AddEipOpts) (*Bandwidth, error) {
	b, err := build.RequestBody(opts, "bandwidth")
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL(client.ProjectID, "bandwidths", id, "insert"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}
	var res struct {
		Bandwidth Bandwidth `json:"bandwidth"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.Bandwidth, err
}
