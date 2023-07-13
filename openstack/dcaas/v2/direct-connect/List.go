package direct_connect

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

type DirectConnect struct {
	TenantID       string `json:"tenant_id,omitempty"`
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

type ListOpts struct {
	ID string `q:"id"`
}

// List is used to obtain the DirectConnects list
func List(c *golangsdk.ServiceClient, opts ListOpts) ([]DirectConnect, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/v2.0/{project_id}/virtual-gateways
	raw, err := c.Get(c.ServiceURL("dcaas", "direct-connects")+q.String(), nil, openstack.StdRequestOpts())
	if err != nil {
		return nil, err
	}

	var res []DirectConnect
	err = extract.IntoSlicePtr(raw.Body, &res, "direct_connects")
	return res, err
}
