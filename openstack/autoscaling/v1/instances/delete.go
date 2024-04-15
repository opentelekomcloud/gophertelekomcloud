package instances

import "github.com/opentelekomcloud/gophertelekomcloud"

type DeleteOpts struct {
	InstanceId string
	// Specifies whether an instance is deleted when it is removed from the AS group.
	// Options:
	// no (default): The instance will not be deleted.
	// yes: The instance will be deleted.
	DeleteInstance string `q:"instance_delete"`
}

func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) error {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("scaling_group_instance", opts.InstanceId).WithQueryParams(&opts).Build()
	if err != nil {
		return err
	}

	_, err = client.Delete(client.ServiceURL(url.String()), nil)
	return err
}
