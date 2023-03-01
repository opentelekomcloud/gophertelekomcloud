package members

import "github.com/opentelekomcloud/gophertelekomcloud"

type DeleteOpts struct {
	ImageId  string `json:"-" required:"true"`
	MemberId string `json:"-" required:"true"`
}

// Delete membership for given image. Callee should be image owner.
func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) (err error) {
	_, err = client.Delete(client.ServiceURL("images", opts.ImageId, "members", opts.MemberId), &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
