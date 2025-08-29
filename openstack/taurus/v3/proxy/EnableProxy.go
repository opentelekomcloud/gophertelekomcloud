package proxy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type EnableProxyOpts struct {
	FlavorRef       string        `json:"flavor_ref" required:"true"`
	NodeNum         int           `json:"node_num" required:"true"`
	ProxyName       string        `json:"proxy_name,omitempty"`
	ProxyMode       string        `json:"proxy_mode,omitempty"`
	NodesReadWeight []NodesWeight `json:"nodes_read_weight,omitempty"`
}

type NodesWeight struct {
	Id     string `json:"id,omitempty"`
	Weight int    `json:"weight,omitempty"`
}

func EnableProxy(client *golangsdk.ServiceClient, instanceId string, opts EnableProxyOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", instanceId, "proxy"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res jobResponse
	return &res.JobId, extract.Into(raw.Body, &res)
}

type jobResponse struct {
	JobId string `json:"job_id"`
}
