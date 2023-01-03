package groups

import "github.com/opentelekomcloud/gophertelekomcloud"

type DeleteOpts struct {
	ScalingGroupId string
	ForceDelete    *bool `q:"force_delete"`
}

func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) (err error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return
	}

	_, err = client.Delete(client.ServiceURL("scaling_group", opts.ScalingGroupId)+q.String(), nil)
	return
}
