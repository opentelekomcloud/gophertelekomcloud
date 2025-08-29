package proxy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateWeightOpts struct {
	MasterWeight  *int                                  `json:"master_weight,omitempty"`
	ReadonlyNodes []TaurusModifyProxyWeightReadonlyNode `json:"readonly_nodes,omitempty"`
}

type TaurusModifyProxyWeightReadonlyNode struct {
	Id     string `json:"id,omitempty"`
	Weight int    `json:"weight,omitempty"`
}

func UpdateWeight(client *golangsdk.ServiceClient, instanceId string, proxyId string, opts UpdateWeightOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("instances", instanceId, "proxy", proxyId, "weight"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res jobResponse
	return &res.JobId, extract.Into(raw.Body, &res)
}
