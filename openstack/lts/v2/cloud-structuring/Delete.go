package cloud_structuring

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// Delete is used to delete a structuring rule of a log stream.
func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	type opts struct {
		// Structuring rule ID.
		ID string `json:"id" required:"true"`
	}
	b, err := build.RequestBody(opts{ID: id}, "")
	if err != nil {
		return err
	}
	// DELETE /v2/{project_id}/lts/struct/template
	_, err = client.DeleteWithBody(client.ServiceURL("lts", "struct", "template"), b, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"content-type": "application/json",
		},
		OkCodes: []int{200},
	})
	if err != nil {
		return err
	}
	return
}
