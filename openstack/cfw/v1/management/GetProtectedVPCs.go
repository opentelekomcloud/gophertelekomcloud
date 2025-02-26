package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetProtectedVPCsParameters represents the query parameters for the protected VPCs list.
type GetProtectedVPCsParameters struct {
	// Protected object ID, which is used to distinguish between Internet border protection and VPC border protection
	// after a cloud firewall is created. You can obtain the ID by calling the Get function in this package.
	// In the return value, find the ID in ProtectObjects[n].ObjectID.
	// If the value of type is 0, the protected object ID belongs to the Internet border.
	// If the value of type is 1, the protected object ID belongs to the VPC border.
	ObjectID string `q:"object_id" required:"true"`
	// Enterprise project ID
	EnterpriseProjectID string `q:"enterprise_project_id,omitempty"`
	// Firewall instance ID. This field is required.
	FwInstanceID string `q:"fw_instance_id,omitempty"`
}

// This function is used to query information about protected VPCs.
func GetProtectedVPCs(client *golangsdk.ServiceClient, objectId string) (*VPCProtectsVo, error) {
	// GET /v1/{project_id}/vpcs/protection
	url, err := golangsdk.NewURLBuilder().WithEndpoints("vpcs", "protection").WithQueryParams(&GetProtectedVPCsParameters{
		FwInstanceID: objectId,
	}).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetProtectedVPCsResponse
	err = extract.Into(raw.Body, &res)
	return &res.Data, err
}

type GetProtectedVPCsResponse struct {
	// Return value for querying protected VPCs.
	Data VPCProtectsVo `json:"data"`
}

type VPCProtectsVo struct {
	// Total number of protected VPCs.
	Total int `json:"total"`
	// The firewall can protect VPCs across accounts.
	// self_total indicates the total number of protected VPCs in the current project.
	SelfTotal int `json:"self_total"`
	// The east-west firewall protection can protect VPCs across accounts.
	// other_total indicates the number of protected VPCs in other projects.
	OtherTotal int `json:"other_total"`
	// The east-west firewall protection can protect VPCs across accounts.
	// protect_vpcs indicates the list of all protected VPCs.
	ProtectVPCs []VpcAttachmentDetail `json:"protect_vpcs"`
	// The east-west firewall protection can protect VPCs across accounts.
	// self_protect_vpcs indicates the list of protected VPCs in the current project.
	SelfProtectVPCs []VpcAttachmentDetail `json:"self_protect_vpcs"`
	// The east-west firewall protection can protect VPCs across accounts.
	// other_protect_vpcs indicates the list of protected VPCs of other projects.
	OtherProtectVPCs []VpcAttachmentDetail `json:"other_protect_vpcs"`
	// Total number of VPC assets of a tenant.
	TotalAssets int `json:"total_assets"`
}

type VpcAttachmentDetail struct {
	// ID of a protected VPC added for east-west protection.
	VPCID string `json:"vpc_id"`
}
