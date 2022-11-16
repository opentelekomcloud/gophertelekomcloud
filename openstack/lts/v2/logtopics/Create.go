package logtopics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// ID of a created log group
	GroupId string `json:"group_id"`
	// Log stream name.
	// The configuration rules are as follows:
	// - Must be a string of 1 to 64 characters.
	// - Only letters, digits, underscores (_), hyphens (-), and periods (.) are allowed.
	// The name cannot start or end with a period.
	LogTopicName string `json:"log_topic_name" required:"true"`
}

func Create(client *golangsdk.ServiceClient, ops CreateOpts) (string, error) {
	b, err := build.RequestBody(ops, "")
	if err != nil {
		return "", err
	}

	// POST /v2.0/{project_id}/log-groups/{group_id}/log-topics
	raw, err := client.Post(client.ServiceURL("log-groups", ops.GroupId, "log-topics"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})

	var res struct {
		ID string `json:"log_topic_id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.ID, err
}
