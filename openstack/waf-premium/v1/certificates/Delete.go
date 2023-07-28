package certificates

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	// DELETE /v1/{project_id}/waf/certificate/{certificate_id}
	_, err = client.Delete(client.ServiceURL("waf", "certificate", id), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	return
}
