package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type deleteQueryParams struct {
	// Query param: ?delete_rds_data
	// This specifies whether data stored on the associated DB instances is deleted.
	// Default value: delete_rds_data=false.
	deleteRdsData string `q:"delete_rds_data"`
}

// This function is used to delete a DDM instance to release all its resources.
func Delete(client *golangsdk.ServiceClient, instanceId string, deleteRdsData bool) (*string, error) {

	deleteData := "false"
	if deleteRdsData {
		deleteData = "true"
	}
	// DELETE https://{Endpoint}/v1/{project_id}/instances/{instance_id}?delete_rds_data=false(OR true)
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceId).WithQueryParams(&deleteQueryParams{deleteRdsData: deleteData}).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Delete(client.ServiceURL(url.String()), nil)

	if err != nil {
		return nil, err
	}

	var res Job
	err = extract.Into(raw.Body, &res)
	return &res.JobId, err

}

type Job struct {

	// DDM instance ID
	Id string `json:"id"`
	// ID of the job for deleting an instance.
	JobId string `json:"job_id"`
}
