package members

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get image member details.
func Get(client *golangsdk.ServiceClient, imageID string, memberID string) (r DetailsResult) {
	_, r.Err = client.Get(client.ServiceURL("images", imageID, "members", memberID), &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
