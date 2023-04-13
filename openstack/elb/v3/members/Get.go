package members

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get retrieves a particular Pool Member based on its unique ID.
func Get(client *golangsdk.ServiceClient, poolID string, memberID string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("pools", poolID, "members", memberID), &r.Body, nil)
	return
}
