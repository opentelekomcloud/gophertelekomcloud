package proxy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type DisableProxyOpts struct {
	ProxyIds []string `json:"proxy_ids,omitempty"`
}

func DisableProxy(client *golangsdk.ServiceClient, instanceID string, opts *DisableProxyOpts) (*string, error) {
	var body interface{}
	if opts != nil {
		body = opts
	}

	raw, err := client.Delete(client.ServiceURL("instances", instanceID, "proxy"), &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
		JSONBody: body,
	})
	if err != nil {
		return nil, err
	}

	var res jobResponse
	return &res.JobId, extract.Into(raw.Body, &res)
}
