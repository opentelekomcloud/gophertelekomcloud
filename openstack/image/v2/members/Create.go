package members

import (
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateOpts struct {
	ImageId string `json:"-" required:"true"`
	// Specifies the image member.
	// The value is the project ID of a tenant.
	Member string `json:"member" required:"true"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Member, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/images/{image_id}/members
	raw, err := client.Post(client.ServiceURL("images", opts.ImageId, "members"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}

type Member struct {
	// Specifies the image sharing status.
	Status string `json:"status"`
	// Specifies the time when a shared image was created. The value is in UTC format.
	CreatedAt time.Time `json:"created_at"`
	// Specifies the time when a shared image was updated. The value is in UTC format.
	UpdatedAt time.Time `json:"updated_at"`
	// Specifies the image ID.
	ImageId string `json:"image_id"`
	// Specifies the member ID, that is, the project ID of the tenant who is to accept the shared image.
	MemberId string `json:"member_id"`
	// Specifies the sharing schema.
	Schema string `json:"schema"`
}
