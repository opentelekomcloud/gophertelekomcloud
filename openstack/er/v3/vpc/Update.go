package vpc

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	RouterID        string  `json:"-"`
	VpcAttachmentID string  `json:"-"`
	Description     *string `json:"description,omitempty"`
	Name            string  `json:"name,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*VpcAttachmentDetails, error) {
	b, err := build.RequestBody(opts, "vpc_attachment")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("enterprise-router", opts.RouterID, "vpc-attachments", opts.VpcAttachmentID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res VpcAttachmentDetails
	return &res, extract.IntoStructPtr(raw.Body, &res, "vpc_attachment")
}
