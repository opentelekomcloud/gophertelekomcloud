package keys

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateDesOpts struct {
	// ID of a CMK
	KeyID string `json:"key_id" required:"true"`
	// CMK description
	KeyDescription string `json:"key_description" required:"true"`
}

func UpdateDes(client *golangsdk.ServiceClient, opts UpdateDesOpts) (*UpdateKeyDes, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("kms", "update-key-description"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res UpdateKeyDes

	err = extract.IntoStructPtr(raw.Body, &res, "key_info")
	return &res, err
}

type UpdateKeyDes struct {
	KeyID    string `json:"key_id"`
	KeyState string `json:"key_state"`
}
