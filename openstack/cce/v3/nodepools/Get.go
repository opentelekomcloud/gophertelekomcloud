package nodepools

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get retrieves a particular node pool based on its unique ID and cluster ID.
func Get(client *golangsdk.ServiceClient, clusterId, nodepoolId string) (*NodePool, error) {
	// GET /api/v3/projects/{project_id}/clusters/{cluster_id}/nodepools/{nodepool_id}
	raw, err := client.Get(client.ServiceURL("clusters", clusterId, "nodepools", nodepoolId), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
		JSONBody:    nil,
	})
	if err != nil {
		return nil, err
	}

	var res NodePool
	return &res, extract.Into(raw.Body, &res)
}
