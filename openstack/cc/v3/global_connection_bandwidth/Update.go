package global_connection_bandwidth

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	// Instance ID.
	ID string `json:"-" required:"true"`
	// Instance name.
	Name string `json:"name,omitempty"`
	// Resource description. Angle brackets (<>) are not allowed.
	Description *string `json:"description,omitempty"`
	// Capacity in Mbit/s. Value range: 2 to 300.
	Size int `json:"size,omitempty"`
	// Billing option. Value options: bwd, 95, 95avr.
	ChargeMode string `json:"charge_mode,omitempty"`
	// Service tier. Value options: Pt, Au, Ag.
	SlaLevel string `json:"sla_level,omitempty"`
	// Instance type. Value options: CC, GEIP, GCN, GSN, ALL.
	BindingService string `json:"binding_service,omitempty"`
	// UUID of the line specification code.
	SpecCodeId string `json:"spec_code_id,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*GlobalConnectionBandwidth, error) {
	b, err := build.RequestBody(opts, "globalconnection_bandwidth")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL(client.DomainID, "gcb", "gcbandwidths", opts.ID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	return extra(raw)
}
