package direct_connect

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	PortType       string `json:"port_type" required:"true"`
	Bandwidth      int    `json:"bandwidth" required:"true"`
	Location       string `json:"location" required:"true"`
	PeerLocation   string `json:"peer_location,omitempty"`
	DeviceID       string `json:"device_id,omitempty"`
	InterfaceName  string `json:"interface_name,omitempty"`
	RedundantID    string `json:"redundant_id,omitempty"`
	Provider       string `json:"provider" required:"true"`
	ProviderStatus string `json:"provider_status,omitempty"`
	Type           string `json:"type,omitempty"`
	HostingID      string `json:"hosting_id,omitempty"`
	ChargeMode     string `json:"charge_mode,omitempty"`
	OrderID        string `json:"order_id,omitempty"`
	ProductID      string `json:"product_id,omitempty"`
	Status         string `json:"status,omitempty"`
	AdminStateUp   bool   `json:"admin_state_up,omitempty"`
}

func Create(c *golangsdk.ServiceClient, opts CreateOpts) (*DirectConnect, error) {
	b, err := build.RequestBody(opts, "direct_connect")
	if err != nil {
		return nil, err
	}
	raw, err := c.Post(c.ServiceURL("dcaas", "direct-connects"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res DirectConnect
	err = extract.IntoStructPtr(raw.Body, &res, "direct_connect")
	return &res, err
}
