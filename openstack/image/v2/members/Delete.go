package members

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete membership for given image. Callee should be image owner.
func Delete(client *golangsdk.ServiceClient, imageID string, memberID string) (r DeleteResult) {
	_, r.Err = client.Delete(client.ServiceURL("images", imageID, "members", memberID), &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
