package volumes

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type DeleteOpts struct {
	// Cascade specifies whether to delete all snapshots created for the volume.
	Cascade bool `q:"cascade"`
}

// Delete deletes the volume with the provided ID.
func Delete(client *golangsdk.ServiceClient, id string, opts DeleteOpts) error {
	q, err := build.QueryString(opts)
	if err != nil {
		return err
	}

	_, err = client.Delete(client.ServiceURL("volumes", id)+q.String(), &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return err
}
