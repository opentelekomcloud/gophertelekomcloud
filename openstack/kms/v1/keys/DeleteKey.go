package keys

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type DeleteOpts struct {
	// ID of a CMK
	KeyID string `json:"key_id" required:"true"`
	// Number of days after which a CMK is scheduled to be deleted
	// (The value ranges from 7 to 1096.)
	PendingDays string `json:"pending_days" required:"true"`
}

func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("kms", "schedule-key-deletion"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
