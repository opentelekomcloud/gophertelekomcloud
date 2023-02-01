package cluster

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	// DELETE /v1.1/{project_id}/clusters/{cluster_id}
	_, err = client.Delete(client.ServiceURL("clusters", id), openstack.StdRequestOpts())
	return
}
