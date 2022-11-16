package loggroups

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get a log group with detailed information by id
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("log-groups", id), &r.Body, nil)
	return
}
