package flow_logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	ResourceType string `json:"resource_type" required:"true"`
	ResourceID   string `json:"resource_id" required:"true"`
	TrafficType  string `json:"traffic_type" required:"true"`
	LogGroupID   string `json:"log_group_id" required:"true"`
	LogTopicID   string `json:"log_topic_id" required:"true"`
	IndexEnabled *bool  `json:"index_enabled,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*FlowLog, error) {
	b, err := build.RequestBody(opts, "flow_log")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL(client.ProjectID, "fl", "flow_logs"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res flowLogResponse
	err = extract.Into(raw.Body, &res)
	return &res.FlowLog, err
}
