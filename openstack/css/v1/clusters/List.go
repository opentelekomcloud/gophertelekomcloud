package clusters

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func List(client *golangsdk.ServiceClient) ([]Cluster, error) {
	raw, err := client.Get(client.ServiceURL("clusters"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Cluster
	err = extract.IntoSlicePtr(raw.Body, &res, "clusters")
	return res, err
}
