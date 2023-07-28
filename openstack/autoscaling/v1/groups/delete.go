package groups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"net/url"
)

type DeleteOpts struct {
	ScalingGroupId string
	ForceDelete    *bool `q:"force_delete"`
}

func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) (err error) {
	var opts2 interface{} = opts
	q, err := build.QueryString(opts2)
	if err != nil {
		return
	}

	_, err = client.Delete(client.ServiceURL("scaling_group", opts.ScalingGroupId)+q.String(), nil)
	return
}
