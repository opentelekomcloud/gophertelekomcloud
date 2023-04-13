package loadbalancers

import "github.com/opentelekomcloud/gophertelekomcloud"

// GetStatuses will return the status of a particular LoadBalancer.
func GetStatuses(client *golangsdk.ServiceClient, id string) (r GetStatusesResult) {
	_, r.Err = client.Get(client.ServiceURL("loadbalancers", id, "statuses"), &r.Body, nil)
	return
}
