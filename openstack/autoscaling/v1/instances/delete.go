package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type DeleteOpts struct {
	InstanceId string
	// Specifies whether an instance is deleted when it is removed from the AS group.
	// Options:
	// no (default): The instance will not be deleted.
	// yes: The instance will be deleted.
	DeleteInstance string `q:"instance_delete"`
}

func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) error {
	q, err := build.QueryString(opts)
	if err != nil {
		return err
	}

	_, err = client.Delete(client.ServiceURL("scaling_group_instance", opts.InstanceId)+q.String(), nil)
	return err
}
