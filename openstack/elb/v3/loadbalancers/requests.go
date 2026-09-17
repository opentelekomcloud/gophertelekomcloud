package loadbalancers

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// Delete will permanently delete a particular LoadBalancer based on its
// unique ID.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(resourceURL(client, id), nil)
	return
}

// GetStatuses will return the status of a particular LoadBalancer.
func GetStatuses(client *golangsdk.ServiceClient, id string) (r GetStatusesResult) {
	_, r.Err = client.Get(statusURL(client, id), &r.Body, nil)
	return
}
