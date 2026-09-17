package fleets

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*ClusterGroup, error) {
	raw, err := client.Get(client.ServiceURL("clustergroups", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ClusterGroup
	return &res, extract.Into(raw.Body, &res)
}
