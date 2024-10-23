package proxy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type EnableProxyOpts struct {
	InstanceID      string       `json:"-"`
	FlavorRef       string       `json:"flavor_ref" required:"true"`
	NodeNum         int          `json:"node_num" required:"true"`
	ProxyName       string       `json:"proxy_name,omitempty"`
	ProxyMode       string       `json:"proxy_mode,omitempty"`
	NodesReadWeight []NodeWeight `json:"nodes_read_weight,omitempty"`
}

type NodeWeight struct {
	ID     string `json:"id"`
	Weight int    `json:"weight"`
}

func EnableProxy(client *golangsdk.ServiceClient, opts EnableProxyOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceID, "proxy"), b, nil, nil)
	if err != nil {
		return "", err
	}

	var res JobId
	err = extract.Into(raw.Body, &res)
	return res.JobId, err
}

type JobId struct {
	JobId string `json:"job_id"`
}
