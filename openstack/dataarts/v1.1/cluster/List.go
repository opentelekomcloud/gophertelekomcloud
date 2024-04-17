package cluster

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// List Querying the Cluster List.
// Send request GET /v1.1/{project_id}/clusters
func List(client *golangsdk.ServiceClient) ([]*ClusterQuery, error) {

	raw, err := client.Get(client.ServiceURL(clustersURL), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []*ClusterQuery
	err = extract.IntoSlicePtr(raw.Body, &res, "clusters")
	return res, err
}
