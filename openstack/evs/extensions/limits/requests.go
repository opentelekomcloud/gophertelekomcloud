package limits

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// Get returns the limits about the currently scoped tenant.
func Get(client *golangsdk.ServiceClient) (r GetResult) {
	url := getURL(client)
	resp, err := client.Get(url, &r.Body, nil)
	_, r.Header, r.Err = golangsdk.ParseResponse(resp, err)
	return
}
