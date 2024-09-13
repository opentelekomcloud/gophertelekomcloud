package schemas

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type deleteQueryParams struct {
	// Whether data stored on the associated DB instances is deleted. The value can be:
	// true: indicates that the data stored on the associated DB instances is deleted.
	// false: indicates that the data stored on the associated DB instances is not deleted. It is left blank by default.
	// Enumerated values: true, false
	deleteRdsData string `q:"delete_rds_data"`
}

// This function  is used to delete a schema to release all its resources.
// databaseName is the name of the schema to be queried, which is case-insensitive
func DeleteSchema(client *golangsdk.ServiceClient, instanceId string, databaseName string, deleteRdsData bool) (*string, error) {

	deleteData := "false"
	if deleteRdsData {
		deleteData = "true"
	}
	// DELETE /v1/{project_id}/instances/{instance_id}/databases/{ddm_dbname}?delete_rds_data={delete_rds_data}
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceId, "databases", databaseName).WithQueryParams(&deleteQueryParams{deleteRdsData: deleteData}).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Delete(client.ServiceURL(url.String()), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res DeleteSchemaResponse
	err = extract.Into(raw.Body, &res)
	return &res.JobId, err

}

type DeleteSchemaResponse struct {
	// ID of the job for deleting an schema.
	JobId string `json:"job_id"`
}
