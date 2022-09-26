package instance

import (
	"net/http"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// DeleteGaussMySqlInstance
// return ID of DB instance deletion task
func DeleteGaussMySqlInstance(client *golangsdk.ServiceClient, instanceId string) (string, error) {
	// DELETE https://{Endpoint}/mysql/v3/{project_id}/instances/{instance_id}
	raw, err := client.Delete(client.ServiceURL("instances", instanceId), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extraJob(err, raw)
}

func extraJob(err error, raw *http.Response) (string, error) {
	if err != nil {
		return "", err
	}

	var res struct {
		JobId string `json:"job_id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.JobId, err
}
