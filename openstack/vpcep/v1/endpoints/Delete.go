package endpoints

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	// DELETE /v1/{project_id}/vpc-endpoints/{vpc_endpoint_id}
	_, err = client.Delete(client.ServiceURL("vpc-endpoints", id), &golangsdk.RequestOpts{
		OkCodes: []int{200, 204},
	})
	return
}
