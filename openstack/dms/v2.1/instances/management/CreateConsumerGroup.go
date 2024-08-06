package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2.1/instances"
)

type CreateConsumerGroupOpts struct {
	// Consumer group name.
	GroupName string `json:"group_name" required:"true"`
	// Consumer group description.
	// Minimum: 0
	// Maximum: 200
	Description string `json:"group_desc"`
}

// CreateConsumerGroup is used to create a consumer group.
// Send POST /v2/{project_id}/kafka/instances/{instance_id}/group
func CreateConsumerGroup(client *golangsdk.ServiceClient, instanceId string, opts *CreateConsumerGroupOpts) (*ErrorResp, error) {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL(instances.ResourcePath, instanceId, groupPath), body, nil, &golangsdk.RequestOpts{})
	if err != nil {
		return nil, err
	}

	var res ErrorResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ErrorResp struct {
	ErrorCode    string `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}
