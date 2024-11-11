package keys

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateDescriptionOpts struct {
	// ID of a CMK
	KeyID string `json:"key_id" required:"true"`
	// CMK description
	KeyDescription string `json:"key_description" required:"true"`
}

func UpdateDescription(client *golangsdk.ServiceClient, opts UpdateDescriptionOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("kms", "update-key-description"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
