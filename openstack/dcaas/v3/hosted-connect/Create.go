package hosted_connect

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Specifies the hosted connection name.
	Name string `json:"name,omitempty"`
	// Provides supplementary information about the hosted connection.
	Description string `json:"description,omitempty"`
	// Specifies the bandwidth of the hosted connection, in Mbit/s.
	Bandwidth int `json:"bandwidth" required:"true"`
	// Specifies the ID of the operations connection on which the hosted connection is created.
	HostingID string `json:"hosting_id" required:"true"`
	// Specifies the VLAN allocated to the hosted connection.
	Vlan int `json:"vlan" required:"true"`
	// Specifies the project ID of the specified tenant for whom a hosted connection is to be created.
	ResourceTenantId string `json:"resource_tenant_id" required:"true"`
	// Specifies the location of the on-premises facility at the other end of the connection,
	// specific to the street or data center name.
	PeerLocation string `json:"peer_location,omitempty"`
}

type HostedConnect struct {
	// Specifies the hosted connection ID.
	ID string `json:"id"`
	// Specifies the project ID.
	TenantID string `json:"tenant_id"`
	// Specifies the connection name.
	Name string `json:"name"`
	// Provides supplementary information about the connection.
	Description string `json:"description"`
	// Specifies the connection bandwidth, in Mbit/s.
	Bandwidth int `json:"bandwidth"`
	// Specifies information about the Direct Connect location.
	Location string `json:"location"`
	// Specifies the location of the on-premises facility at the other end of the connection,
	// specific to the street or data center name.
	PeerLocation string `json:"peer_location"`
	// Specifies the ID of the operations connection on which the hosted connection is created.
	HostingId string `json:"hosting_id"`
	// Specifies the provider of the leased line.
	Provider string `json:"provider"`
	// Specifies the provider of the leased line.
	AdminStateUp bool `json:"admin_state_up"`
	// Specifies the VLAN allocated to the hosted connection.
	Vlan int `json:"vlan"`
	// Specifies the operating status.
	Status string `json:"status"`
	// Specifies when the connection was requested.
	// The UTC time format is yyyy-MM-ddTHH:mm:ss.SSSZ.
	ApplyTime string `json:"apply_time"`
	// Specifies when the connection was created.
	// The UTC time format is yyyy-MM-ddTHH:mm:ss.SSSZ.
	CreatedAt string `json:"create_time"`
	// Specifies the carrier status. The status can be ACTIVE or DOWN.
	ProviderStatus string `json:"provider_status"`
	// Specifies the type of the port used by the connection.
	// The value can be 1G, 10G, 40G, or 100G.
	PortType string `json:"port_type"`
	// Specifies the type of the connection. The value is hosted.
	Type string `json:"type"`
}

func Create(c *golangsdk.ServiceClient, opts CreateOpts) (*HostedConnect, error) {
	b, err := build.RequestBody(opts, "direct_connect")
	if err != nil {
		return nil, err
	}
	raw, err := c.Post(c.ServiceURL("dcaas", "hosted-connects"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res HostedConnect
	err = extract.IntoStructPtr(raw.Body, &res, "hosted_connect")
	return &res, err
}
