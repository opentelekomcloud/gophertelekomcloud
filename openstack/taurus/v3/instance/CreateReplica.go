package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateReplicaOpts struct {
	Priorities []int `json:"priorities" required:"true"`
}

func CreateReplica(client *golangsdk.ServiceClient, instanceId string, opts CreateReplicaOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", instanceId, "nodes", "enlarge"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201, 202},
	})
	if err != nil {
		return nil, err
	}

	var res jobResponse
	return &res.JobID, extract.Into(raw.Body, &res)
}

type jobResponse struct {
	JobID string `json:"job_id"`
}
