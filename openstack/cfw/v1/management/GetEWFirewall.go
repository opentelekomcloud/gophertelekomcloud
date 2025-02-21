package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetQueryParameters represents the query parameters for the firewall instance list.
type GetEWFirewallQueryParameters struct {
	// Offset, which specifies the start position of the record to be returned. The value must be a number no less than 0. The default value is 0.
	Offset int `q:"offset" required:"true"`
	// Number of records displayed on each page. The value ranges from 1 to 1024.
	Limit int `q:"limit" required:"true"`
	// Enterprise project ID
	EnterpriseProjectID string `q:"enterprise_project_id,omitempty"`
	// Firewall instance ID. This field is required.
	FwInstanceID string `q:"fw_instance_id" required:"true"`
}

// Get is used to query details about a Firewall instance.
func GetEWFirewall(client *golangsdk.ServiceClient, firewallId string) (*GetEastWestFirewallResponseBody, error) {
	// GET /v1/{project_id}/firewall/east-west
	url, err := golangsdk.NewURLBuilder().WithEndpoints("firewall", "east-west").WithQueryParams(&GetEWFirewallQueryParameters{
		FwInstanceID: firewallId,
	}).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetEWFirewallResponse
	err = extract.Into(raw.Body, &res)
	return &res.Data, err
}

type GetEWFirewallResponse struct {
	// Returned data for obtaining the east-west firewall list.
	Data GetEastWestFirewallResponseBody `json:"data"`
}

type GetEastWestFirewallResponseBody struct {
	// Protected object ID.
	ObjectID string `json:"object_id"`
	// Project ID.
	ProjectID string `json:"project_id"`
	// Protection status: 0 (enabled), 1 (disabled).
	Status int `json:"status"`
	// Information about the subnet associated with a cloud firewall.
	FirewallAssociatedSubnets []SubnetInfo `json:"firewall_associated_subnets"`
	// Information about the associated enterprise router in the outbound direction.
	ER ErInstance `json:"er"`
	// Information about the inspection VPC.
	InspectionVPC VpcDetail `json:"inspection_vpc"`
	// East-west protected resource information.
	ProtectInfos []EwProtectResourceInfo `json:"protect_infos"`
	// Total number of protected VPCs.
	Total int `json:"total"`
	// Offset specifying the start position of the record to be returned.
	Offset int `json:"offset"`
	// Number of records displayed on each page (range: 1–1024).
	Limit int `json:"limit"`
	// Protection mode. The value is "er".
	Mode string `json:"mode"`
}

// SubnetInfo represents information about a subnet associated with a cloud firewall.
type SubnetInfo struct {
	// ID of the AZ where a subnet is located.
	AvailabilityZone string `json:"availability_zone"`
	// Available IP address ranges for subnets in a VPC.
	CIDR string `json:"cidr"`
	// Subnet name.
	Name string `json:"name"`
	// Subnet ID.
	ID string `json:"id"`
	// Subnet gateway IP.
	GatewayIP string `json:"gateway_ip"`
	// UUID generated when a VPC is created.
	VPCID string `json:"vpc_id"`
	// Whether IPv6 is supported (true/false).
	IPv6Enable bool `json:"ipv6_enable"`
}

// ErInstance represents information about an enterprise router.
type ErInstance struct {
	// Enterprise router ID.
	ID string `json:"id"`
	// Enterprise router name.
	Name string `json:"name"`
	// Router status: pending, available, modifying, deleting, or failed.
	State string `json:"state"`
	// Enterprise project ID.
	EnterpriseProjectID string `json:"enterprise_project_id"`
	// Project ID.
	ProjectID string `json:"project_id"`
	// Whether IPv6 is enabled (true/false).
	EnableIPv6 bool `json:"enable_ipv6"`
	// Connection ID of the enterprise router.
	AttachmentID string `json:"attachment_id"`
}

// VpcDetail represents information about an inspection VPC.
type VpcDetail struct {
	// Random UUID generated when an inspection VPC is created.
	ID string `json:"id"`
	// Inspection VPC name.
	Name string `json:"name"`
	// Available subnet ranges in a VPC.
	CIDR string `json:"cidr"`
}

// EwProtectResourceInfo represents east-west protected resource information.
type EwProtectResourceInfo struct {
	// Protected resource type: 0 (VPC), 1 (VGW), 2 (VPN), 3 (peering).
	ProtectedResourceType int `json:"protected_resource_type"`
	// Protected resource name.
	ProtectedResourceName string `json:"protected_resource_name"`
	// Protected resource ID.
	ProtectedResourceID string `json:"protected_resource_id"`
	// Name of the NAT gateway to be protected.
	ProtectedResourceNATName string `json:"protected_resource_nat_name"`
	// ID of the NAT gateway to be protected.
	ProtectedResourceNATID string `json:"protected_resource_nat_id"`
	// Tenant ID of a protected resource.
	ProtectedResourceProjectID string `json:"protected_resource_project_id"`
	// Protected resource mode. The value is "er".
	ProtectedResourceMode string `json:"protected_resource_mode"`
	// Protection status: 0 (associated), 1 (not associated).
	Status int `json:"status"`
}
