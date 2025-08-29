package proxy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Resize(client *golangsdk.ServiceClient, instanceId string, proxyId string, flavorRef string) (*string, error) {
	b, err := build.RequestBody(flavorRef, "flavor_ref")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("instances", instanceId, "proxy", proxyId, "flavor"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res jobResponse
	return &res.JobId, extract.Into(raw.Body, &res)
}
