package clusters

import "github.com/opentelekomcloud/gophertelekomcloud"

// List returns collection of clusters.
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Clusters, error) {
	var res ListResult
	raw, err := client.Get(client.ServiceURL("clusters"), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts, JSONBody: nil,
	})

	allClusters, err := res.ExtractClusters()
	if err != nil {
		return nil, err
	}

	return filterClusters(allClusters, opts), nil
}
