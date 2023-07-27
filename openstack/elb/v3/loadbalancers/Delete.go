package loadbalancers

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a particular LoadBalancer based on its
// unique ID.
func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	// DELETE /v3/{project_id}/elb/loadbalancers/{loadbalancer_id}
	_, err = client.Delete(client.ServiceURL("loadbalancers", id), nil)
	return
}
