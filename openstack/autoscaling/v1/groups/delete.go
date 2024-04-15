package groups

import "github.com/opentelekomcloud/gophertelekomcloud"

type DeleteOpts struct {
	ScalingGroupId string
	ForceDelete    *bool `q:"force_delete"`
}

func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) (err error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("scaling_group", opts.ScalingGroupId).WithQueryParams(&opts).Build()
	if err != nil {
		return
	}

	_, err = client.Delete(client.ServiceURL(url.String()), nil)
	return
}
