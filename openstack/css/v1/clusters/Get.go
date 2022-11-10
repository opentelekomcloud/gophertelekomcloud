package clusters

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*Cluster, error) {
	raw, err := client.Get(client.ServiceURL("clusters", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Cluster
	err = extract.Into(raw.Body, &res)
	return &res, err
}
