package instance

import (
	"net/http"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// DeleteInstance
// return ID of DB instance deletion task
func DeleteInstance(client *golangsdk.ServiceClient, instanceId string) (*string, error) {
	// DELETE https://{Endpoint}/mysql/v3/{project_id}/instances/{instance_id}
	raw, err := client.Delete(client.ServiceURL("instances", instanceId), &golangsdk.RequestOpts{
		OkCodes:     []int{200, 202},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	return extraJob(err, raw)
}

func extraJob(err error, raw *http.Response) (*string, error) {
	if err != nil {
		return nil, err
	}

	var res struct {
		JobId string `json:"job_id"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.JobId, err
}
