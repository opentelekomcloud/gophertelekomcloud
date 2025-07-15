package nodes

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// Delete will permanently delete a particular node based on its unique ID and cluster ID.
func Delete(client *golangsdk.ServiceClient, clusterID, nodeID string) error {
	// DELETE /api/v3/projects/{project_id}/clusters/{cluster_id}/nodes/{node_id}
	_, err := client.Delete(client.ServiceURL("clusters", clusterID, "nodes", nodeID), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
		JSONBody:    nil,
	})
	if err != nil {
		return err
	}
	return nil
}
