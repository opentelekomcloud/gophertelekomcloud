package services

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	// DELETE /v1/{project_id}/vpc-endpoint-services/{vpc_endpoint_service_id}
	_, err = client.Delete(client.ServiceURL("vpc-endpoint-services", id), &golangsdk.RequestOpts{
		OkCodes: []int{200, 204},
	})
	return
}
