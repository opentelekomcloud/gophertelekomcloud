package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Policy name
	Name string `json:"name"`
	// Policy action
	Action *PolicyAction `json:"action"`
	// Policy option
	Options *PolicyOption `json:"options"`
	// Feature-based anti-crawler protection mode.
	// The default protection mode is Log only.
	RobotAction *PolicyAction `json:"robot_action"`
	// Protection level
	Level int `json:"level"`
	// Detection mode in the precise protection rule
	FullDetection *bool `json:"full_detection"`
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*Policy, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/waf/policy/{policy_id}
	raw, err := client.Patch(client.ServiceURL("waf", "policy", id), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	if err != nil {
		return nil, err
	}

	var res Policy
	return &res, extract.Into(raw.Body, &res)
}
