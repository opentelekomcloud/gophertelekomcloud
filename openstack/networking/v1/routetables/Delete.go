package routetables

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// Delete will permanently delete a particular route table based on its unique ID
func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	// DELETE /v1.1/{project_id}/clusters/{cluster_id}
	_, err = client.Delete(client.ServiceURL(client.ProjectID, "routetables", id), openstack.StdRequestOpts())
	return
}
