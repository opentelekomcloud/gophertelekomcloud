package members

import (
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type GetMemberOpts struct {
	ImageId  string `json:"-" required:"true"`
	MemberId string `json:"-" required:"true"`
}

func GetMember(client *golangsdk.ServiceClient, opts GetMemberOpts) (*Member, error) {
	// GET /v2/images/{image_id}/members/{member_id}
	raw, err := client.Get(client.ServiceURL("images", opts.ImageId, "members", opts.MemberId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Member
	err = extract.Into(raw.Body, &res)
	return &res, err
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
	// Specifies the member ID.
	MemberId string `json:"member_id"`
	// Specifies the sharing schema.
	Schema string `json:"schema"`
}
