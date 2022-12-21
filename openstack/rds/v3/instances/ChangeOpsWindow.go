package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ChangeOpsWindowOpts struct {
	InstanceId string `json:"-"`
	// Specifies the start time.
	// The value must be a valid value in the "HH:MM" format. The current time is in the UTC format.
	StartTime string `json:"start_time"`
	// Specifies the end time.
	// The value must be a valid value in the "HH:MM" format. The current time is in the UTC format.
	// NOTE
	// The interval between the start time and end time must be four hours.
	EndTime string `json:"end_time"`
}

func ChangeOpsWindow(client *golangsdk.ServiceClient, opts ChangeOpsWindowOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// PUT https://{Endpoint}/v3/{project_id}/instances/{instance_id}/ops-window
	_, err = client.Put(client.ServiceURL("instances", opts.InstanceId, "ops-window"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
