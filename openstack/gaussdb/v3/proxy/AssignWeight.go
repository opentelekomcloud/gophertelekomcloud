package proxy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type AssignWeightOpts struct {
	InstanceID    string          `json:"-"`
	ProxyID       string          `json:"-"`
	MasterWeight  *int            `json:"flavor_ref,omitempty"`
	ReadonlyNodes *[]ReadonlyNode `json:"readonly_nodes,omitempty"`
}

type ReadonlyNode struct {
	ID     string `json:"id"`
	Weight *int   `json:"weight"`
}

func AssignWeight(client *golangsdk.ServiceClient, opts AssignWeightOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	raw, err := client.Put(client.ServiceURL("instances", opts.InstanceID, "proxy", opts.ProxyID, "weight"), b, nil, nil)
	if err != nil {
		return "", err
	}

	var res JobId
	err = extract.Into(raw.Body, &res)
	return res.JobId, err
}
