package instances

import (
	"net/http"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func StartupInstance(client *golangsdk.ServiceClient, instanceId string) (*string, error) {
	// POST https://{Endpoint}/v3/{project_id}/instances/{instance_id}/action/startup
	raw, err := client.Post(client.ServiceURL("instances", instanceId, "action", "startup"), struct{}{}, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extraJob(err, raw)
}

func extraJob(err error, raw *http.Response) (*string, error) {
	if err != nil {
		return nil, err
	}

	var res JobId
	err = extract.Into(raw.Body, &res)
	return &res.JobId, err
}
