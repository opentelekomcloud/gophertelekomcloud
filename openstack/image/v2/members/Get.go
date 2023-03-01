package members

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v2/members"
)

// Get image member details.
func Get(client *golangsdk.ServiceClient, opts MemberOpts) (*members.Member, error) {
	// GET /v2/images/{image_id}/members/{member_id}
	raw, err := client.Get(client.ServiceURL("images", opts.ImageId, "members", opts.MemberId), nil, nil)
	return extra(err, raw)
}

func extra(err error, raw *http.Response) (*members.Member, error) {
	if err != nil {
		return nil, err
	}

	var res members.Member
	err = extract.Into(raw.Body, res)
	return &res, err
}
