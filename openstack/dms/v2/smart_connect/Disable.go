package smart_connect

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This API is used to disable Smart Connect for a pay-per-use instance.
// Send POST /v2/{project_id}/kafka/instances/{instance_id}/delete-connector
func Disable(client *golangsdk.ServiceClient, id string) (*DisableResp, error) {
	// Providing empty body results to an error therefore filler is used
	emptyBody := struct{}{}

	b, err := build.RequestBody(emptyBody, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("kafka", "instances", id, "delete-connector"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res DisableResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type DisableResp struct {
	// Task ID.
	JobId string `json:"job_id"`
}
