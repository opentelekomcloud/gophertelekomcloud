package global_connection_bandwidth

import (
	"net/http"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Instance name.
	Name string `json:"name" required:"true"`
	// Resource description. Angle brackets (<>) are not allowed.
	Description string `json:"description,omitempty"`
	// Cross-border attribute.
	Bordercross *bool `json:"bordercross" required:"true"`
	// Bandwidth type. Value options: Area, TrsArea, SubArea, Region.
	Type string `json:"type" required:"true"`
	// ID of the enterprise project that the resource belongs to.
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
	// Billing option. Value options: bwd, 95, 95avr.
	ChargeMode string `json:"charge_mode" required:"true"`
	// Capacity of a global connection bandwidth, in Mbit/s. Value range: 2 to 300.
	Size int `json:"size" required:"true"`
	// Service tier. Value options: Pt (Platinum), Au (Gold), Ag (Silver).
	SlaLevel string `json:"sla_level,omitempty"`
	// Local access point.
	LocalArea string `json:"local_area,omitempty"`
	// Remote access point.
	RemoteArea string `json:"remote_area,omitempty"`
	// UUID of the line specification code.
	SpecCodeId string `json:"spec_code_id,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*GlobalConnectionBandwidth, error) {
	b, err := build.RequestBody(opts, "globalconnection_bandwidth")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL(client.DomainID, "gcb", "gcbandwidths"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	if err != nil {
		return nil, err
	}

	return extra(raw)
}

// GlobalConnectionBandwidth is a global connection bandwidth instance.
type GlobalConnectionBandwidth struct {
	// Instance ID.
	ID string `json:"id"`
	// Instance name.
	Name string `json:"name"`
	// Resource description.
	Description string `json:"description"`
	// Account ID.
	DomainId string `json:"domain_id"`
	// Cross-border attribute.
	Bordercross bool `json:"bordercross"`
	// Bandwidth type.
	Type string `json:"type"`
	// Service binding type. Value options: CC, GEIP, GCN, GSN, ALL.
	BindingService string `json:"binding_service"`
	// ID of the enterprise project that the resource belongs to.
	EnterpriseProjectId string `json:"enterprise_project_id"`
	// Billing option.
	ChargeMode string `json:"charge_mode"`
	// Capacity in Mbit/s.
	Size int `json:"size"`
	// Service tier.
	SlaLevel string `json:"sla_level"`
	// Local access point.
	LocalArea string `json:"local_area"`
	// Remote access point.
	RemoteArea string `json:"remote_area"`
	// Local access point code.
	LocalSiteCode string `json:"local_site_code"`
	// Remote access point code.
	RemoteSiteCode string `json:"remote_site_code"`
	// Status. Value options: NORMAL, FREEZED.
	AdminState string `json:"admin_state"`
	// Freeze status.
	Frozen bool `json:"frozen"`
	// UUID of the line specification code.
	SpecCodeId string `json:"spec_code_id"`
	// Time when the resource was created.
	CreatedAt string `json:"created_at"`
	// Time when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// Whether the bandwidth can be used by multiple instances.
	EnableShare bool `json:"enable_share"`
	// ID of the enterprise project that the resource belongs to.
	EpsId string `json:"eps_id"`
	// Associated instances.
	Instances []AssociatedInstance `json:"instances"`
	// Directional connections.
	DirectionalConnections []DirectionalConnection `json:"directional_connections"`
}

// AssociatedInstance is an instance associated with a global connection bandwidth.
type AssociatedInstance struct {
	// Instance ID.
	ID string `json:"id"`
	// Instance type.
	Type string `json:"type"`
	// Region ID.
	RegionId string `json:"region_id"`
	// Project ID.
	ProjectId string `json:"project_id"`
}

// DirectionalConnection is a directional connection of a global connection bandwidth.
type DirectionalConnection struct {
	// Connection name.
	Name string `json:"name"`
	// Connection ID.
	ID string `json:"id"`
	// Local access point code.
	LocalSiteCode string `json:"local_site_code"`
	// Remote access point code.
	RemoteSiteCode string `json:"remote_site_code"`
}

func extra(raw *http.Response) (*GlobalConnectionBandwidth, error) {
	var res struct {
		GlobalConnectionBandwidth GlobalConnectionBandwidth `json:"globalconnection_bandwidth"`
	}
	if err := extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return &res.GlobalConnectionBandwidth, nil
}
