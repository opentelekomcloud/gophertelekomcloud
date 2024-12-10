package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get retrieves a particular cluster based on its unique ID.
func Get(client *golangsdk.ServiceClient, clusterId string) (*Clusters, error) {
	// GET /api/v3/projects/{project_id}/clusters/{cluster_id}
	raw, err := client.Get(client.ServiceURL("clusters", clusterId), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
		JSONBody:    nil,
	})
	if err != nil {
		return nil, err
	}

	var res Clusters
	return &res, extract.Into(raw.Body, &res)
}
