package vpc

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, erID, vpcID string) (*VpcAttachmentsResp, error) {
	raw, err := client.Get(client.ServiceURL("enterprise-router", erID, "vpc-attachments", vpcID), nil, nil)
	if err != nil {
		return nil, err
	}

	var res VpcAttachmentsResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type VpcAttachmentsResp struct {
	// VPC attachment
	VpcAttachment *VpcAttachmentDetails `json:"vpc_attachment"`
	// Request ID
	RequestID string `json:"request_id"`
}
