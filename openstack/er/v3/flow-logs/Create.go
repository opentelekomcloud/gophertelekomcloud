package flow_logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Enterprise router ID
	RouterID string `json:"-" required:"true"`
	// Request body for creating a flow log
	FlowLog *FlowLog `json:"flow_log" required:"true"`
}

type FlowLog struct {
	// Flow log name
	Name string `json:"name" required:"true"`
	// Flow log description
	Description string `json:"description,omitempty"`
	// Type of resource whose flow logs are collected.
	// VPC attachments
	// Virtual gateway attachments
	ResourceType string `json:"resource_type" required:"true"`
	// Resource ID
	ResourceId string `json:"resource_id" required:"true"`
	// Log group ID. Obtain the log group ID by referring to the Log Tank Service.
	LogGroupId string `json:"log_group_id" required:"true"`
	// Log stream ID. Obtain the log stream ID by referring to the Log Tank Service
	LogStreamId string `json:"log_stream_id" required:"true"`
	// Flow log storage type.
	// LTS: Logs are stored on LTS servers.
	LogStoreType string `json:"log_store_type" required:"true"`
	// Flow log storage name. This parameter is not supported for now.
	LogStoreName string `json:"log_store_name,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*FlowLogResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("enterprise-router", opts.RouterID, "flow-logs"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res FlowLogResponse
	return &res, extract.IntoStructPtr(raw.Body, &res, "flow_log")
}

type FlowLogResponse struct {
	// Flow log ID
	ID string `json:"id"`
	// Flow log name
	Name string `json:"name"`
	// Flow log description
	Description string `json:"description"`
	// Project ID of the flow log task creator
	ProjectId string `json:"project_id"`
	// Type of resource whose flow logs are collected.
	// VPC attachments
	// Virtual gateway attachments
	ResourceType string `json:"resource_type"`
	// Resource ID
	ResourceId string `json:"resource_id"`
	// Log group ID. Obtain the log group ID by referring to the Log Tank Service.
	LogGroupId string `json:"log_group_id"`
	// Log stream ID. Obtain the log stream ID by referring to the Log Tank Service
	LogStreamId string `json:"log_stream_id"`
	// Flow log storage type.
	// LTS: Logs are stored on LTS servers.
	LogStoreType string `json:"log_store_type"`
	// Flow log storage name. This parameter is not supported for now.
	LogStoreName string `json:"log_store_name"`
	// Log aggregation time, in seconds. The value ranges from 60 to 600.
	LogAggregationInterval int `json:"log_aggregation_interval"`
	// Creation time in the format YYYY-MM-DDTHH:mm:ss.sssZ
	CreatedAt string `json:"created_at"`
	// Update time in the format YYYY-MM-DDTHH:mm:ss.sssZ
	UpdatedAt string `json:"updated_at"`
	// Flow log status. Value options: pending, available, modifying, deleting, deleted, and failed
	Status string `json:"state"`
	// Whether to enable flow logs. The value can be true or false.
	Enabled bool `json:"enabled"`
}
