package keys

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type DeleteOpts struct {
	// ID of a CMK
	KeyID string `json:"key_id" required:"true"`
	// Number of days after which a CMK is scheduled to be deleted
	// (The value ranges from 7 to 1096.)
	PendingDays string `json:"pending_days" required:"true"`
}

func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) (*UpdateKeyState, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("kms", "schedule-key-deletion"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	if err != nil {
		return nil, err
	}

	var res UpdateKeyState

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type UpdateKeyState struct {
	KeyID    string `json:"key_id"`
	KeyState string `json:"key_state"`
}
