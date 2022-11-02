package volumes

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get retrieves the Volume with the provided ID. To extract the Volume object
// from the response, call the Extract method on the GetResult.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	resp, err := client.Get(client.ServiceURL("volumes", id), &r.Body, nil)
	_, r.Header, r.Err = golangsdk.ParseResponse(resp, err)
	return
}
