package tags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ActionOpts struct {
	Action string      `json:"action,omitempty"`
	Tags   []TagEntity `json:"tags,omitempty"`
}

type TagEntity struct {
	Key   string  `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// Action sends POST /v2/{project_id}/kafka/{instance_id}/tags/action
func Action(client *golangsdk.ServiceClient, instanceID string, opts ActionOpts) error {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("kafka", instanceID, "tags", "action").
		Build()
	if err != nil {
		return err
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
