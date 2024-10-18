package proxy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type EnlargeProxyOpts struct {
	InstanceID string `json:"-"`
	NodeNum    int    `json:"node_num" required:"true"`
	ProxyID    string `json:"proxy_id,omitempty"`
}

func EnlargeProxy(client *golangsdk.ServiceClient, opts EnlargeProxyOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceID, "proxy", "enlarge"), b, nil, nil)
	if err != nil {
		return "", err
	}

	var res JobId
	err = extract.Into(raw.Body, &res)
	return res.JobId, err
}
