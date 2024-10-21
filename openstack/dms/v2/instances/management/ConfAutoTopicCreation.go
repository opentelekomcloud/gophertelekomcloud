package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ConfAutoTopicCreationOpts struct {
	EnableAutoTopic bool `json:"enable_auto_topic" required:"true"`
}

// ConfAutoTopicCreation is used to enable or disable automatic topic creation.
// Send POST /v2/{project_id}/instances/{instance_id}/autotopic
func ConfAutoTopicCreation(client *golangsdk.ServiceClient, instanceId string, opts ConfAutoTopicCreationOpts) error {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("instances", instanceId, "autotopic"), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	return err
}
