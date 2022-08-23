package instances

import "github.com/opentelekomcloud/gophertelekomcloud"

type DeleteOpts struct {
	DeleteInstance bool `q:"instance_delete"`
}

func Delete(client *golangsdk.ServiceClient, id string, opts DeleteOpts) error {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return err
	}

	_, err = client.Delete(client.ServiceURL("scaling_group_instance", id)+q.String(), nil)
	return err
}
