package propagation

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type DeleteOpts struct {
	RouterID     string `json:"-" required:"true"`
	RouteTableID string `json:"-" required:"true"`
	AttachmentID string `json:"attachment_id,omitempty"`
}

func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("enterprise-router", opts.RouterID, "route-tables", opts.RouteTableID, "disable-propagations"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
