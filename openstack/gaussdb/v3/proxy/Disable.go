package proxy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type DisableProxyOpts struct {
	InstanceID string    `json:"-"`
	ProxyIDs   *[]string `json:"proxy_ids,omitempty"`
}

func Disable(client *golangsdk.ServiceClient, opts DisableProxyOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	raw, err := client.DeleteWithBody(client.ServiceURL("instances", opts.InstanceID, "proxy"), b, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 202},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return "", err
	}

	var res JobId
	err = extract.Into(raw.Body, &res)
	return res.JobId, err
}
