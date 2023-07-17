package direct_connect

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

type DirectConnect struct {
	ID                    string `json:"id"`
	TenantID              string `json:"tenant_id"`
	Name                  string `json:"name"`
	Description           string `json:"description"`
	PortType              string `json:"port_type"`
	Bandwidth             int    `json:"bandwidth"`
	Location              string `json:"location"`
	PeerLocation          string `json:"peer_location"`
	DeviceID              string `json:"device_id"`
	InterfaceName         string `json:"interface_name"`
	RedundantID           string `json:"redundant_id"`
	Provider              string `json:"provider"`
	ProviderStatus        string `json:"provider_status"`
	Type                  string `json:"type"`
	HostingID             string `json:"hosting_id"`
	VLAN                  int    `json:"vlan"`
	ChargeMode            string `json:"charge_mode"`
	ApplyTime             string `json:"apply_time"`
	CreateTime            string `json:"create_time"`
	DeleteTime            string `json:"delete_time"`
	OrderID               string `json:"order_id"`
	ProductID             string `json:"product_id"`
	Status                string `json:"status"`
	AdminStateUp          bool   `json:"admin_state_up"`
	SpecCode              string `json:"spec_code"`
	Applicant             string `json:"applicant"`
	Mobile                string `json:"mobile"`
	Email                 string `json:"email"`
	RegionID              string `json:"region_id"`
	ServiceKey            string `json:"service_key"`
	CableLabel            string `json:"cable_label"`
	PeerPortType          string `json:"peer_port_type"`
	PeerProvider          string `json:"peer_provider"`
	OnestopProductID      string `json:"onestop_product_id"`
	BuildingLineProductID string `json:"building_line_product_id"`
	LastOnestopProductID  string `json:"last_onestop_product_id"`
	PeriodType            int    `json:"period_type"`
	PeriodNum             int    `json:"period_num"`
	Reason                string `json:"reason"`
	VGWType               string `json:"vgw_type"`
	LagID                 string `json:"lag_id"`
}

// List is used to obtain the DirectConnects list
func List(c *golangsdk.ServiceClient, id string) ([]DirectConnect, error) {

	// GET https://{Endpoint}/v2.0/dcaas/direct-connects?id={id}
	raw, err := c.Get(c.ServiceURL(fmt.Sprintf("dcaas/direct-connects?id=%s", id)), nil, openstack.StdRequestOpts())
	if err != nil {
		return nil, err
	}

	var res []DirectConnect
	err = extract.IntoSlicePtr(raw.Body, &res, "direct_connects")
	return res, err
}
