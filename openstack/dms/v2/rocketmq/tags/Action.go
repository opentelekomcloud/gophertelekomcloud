package tags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type actionOpts struct {
	// Operation: create or delete.
	Action string `json:"action" required:"true"`
	// Tag list. A maximum of 20 tags can be added to an instance.
	Tags []tags.ResourceTag `json:"tags" required:"true"`
}

// Create adds tags to a RocketMQ instance in batches.
// Send POST to /v2/{project_id}/rocketmq/{instance_id}/tags/action
func Create(client *golangsdk.ServiceClient, instanceID string, tagList []tags.ResourceTag) error {
	return doAction(client, instanceID, actionOpts{Action: "create", Tags: tagList})
}

// Delete removes tags from a RocketMQ instance in batches.
// Send POST to /v2/{project_id}/rocketmq/{instance_id}/tags/action
func Delete(client *golangsdk.ServiceClient, instanceID string, tagList []tags.ResourceTag) error {
	return doAction(client, instanceID, actionOpts{Action: "delete", Tags: tagList})
}

func doAction(client *golangsdk.ServiceClient, instanceID string, opts actionOpts) error {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("rocketmq", instanceID, "tags", "action").Build()
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
