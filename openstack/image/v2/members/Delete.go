package members

import "github.com/opentelekomcloud/gophertelekomcloud"

type MemberOpts struct {
	ImageId string `json:"-" required:"true"`
	// Specifies the image member.
	// The value is the project ID of a tenant.
	MemberId string `json:"member" required:"true"`
}

// Delete membership for given image. Callee should be image owner.
func Delete(client *golangsdk.ServiceClient, opts MemberOpts) (err error) {
	_, err = client.Delete(client.ServiceURL("images", opts.ImageId, "members", opts.MemberId), &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
