package volumetypes

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will delete the existing Volume Type with the provided ID.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	resp, err := client.Delete(client.ServiceURL("types", id), nil)
	_, r.Header, r.Err = golangsdk.ParseResponse(resp, err)
	return
}
