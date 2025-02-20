package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type EWFirewallQueryParams struct {
	// Enterprise project ID
	EnterpriseProjectId string `q:"enterprise_project_id,omitempty"`
	// Firewall ID
	FwInstanceId string `q:"fw_instance_id" required:"true"`
}

type CreateEWFirewallOpts struct {
	// ID of the associated enterprise router in the outbound direction.
	ERID string `json:"er_id" required:"true"`
	// Inspection VPC ID.
	InspectionVPCID string `json:"inspection_vpc_id,omitempty"`
	// Subnet associated with an enterprise router.
	ERAssociatedSubnet *AssociatedSubnet `json:"er_associated_subnet,omitempty"`
	// List of subnets associated with a firewall.
	FirewallAssociatedSubnets []AssociatedSubnet `json:"firewall_associated_subnets,omitempty"`
}

// AssociatedSubnet represents a subnet associated with a router or firewall.
type AssociatedSubnet struct {
	// AZ.
	AZ string `json:"az" required:"true"`

	// Subnet CIDR block.
	SubnetSegment string `json:"subnet_segment" required:"true"`

	// Subnet name.
	SubnetName string `json:"subnet_name" required:"true"`
}

// Create function is used to create an east-west firewall
func CreateEWFirewall(client *golangsdk.ServiceClient, firewallId string, opts CreateEWFirewallOpts) (*CreateEWFirewallResp, error) {
	// POST /v1/{project_id}/firewall/east-west?fw_instance_id=XXXXX&enterprise_project_id=default
	url, err := golangsdk.NewURLBuilder().WithEndpoints("firewall", "east-west").WithQueryParams(&EWFirewallQueryParams{
		FwInstanceId: firewallId,
	}).Build()
	if err != nil {
		return nil, err
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res CreateEWFResponse
	return &res.Data, extract.Into(raw.Body, &res)
}

type CreateEWFResponse struct {
	// Return value for creating an east-west firewall.
	Data CreateEWFirewallResp `json:"data"`
}

type CreateEWFirewallResp struct {
	// East-west protection ID, corresponding to the object_id field.
	ID string `json:"id"`
	// Enterprise router information.
	ER ER `json:"er"`
	// Information about the inspection VPC.
	InspVPC CreateEWFirewallInspectVpcResp `json:"inspection_vpc"`
}

type ER struct {
	// Enterprise router ID, which is referenced when east-west protection is created.
	ERID string `json:"er_id"`
	// Connection ID of an enterprise router.
	ERAttachID string `json:"er_attach_id"`
}

type CreateEWFirewallInspectVpcResp struct {
	// ID of an inspection VPC.
	VPCID string `json:"vpc_id"`
	// Subnet ID list of the created inspection VPC.
	SubnetIDs []string `json:"subnet_ids"`
}
