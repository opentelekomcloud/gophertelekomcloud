package log

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Specifies the load balancer ID.
	LoadbalancerId string `json:"loadbalancer_id" required:"true"`
	// Specifies the log group ID.
	// This parameter is available for all services other than ELB.
	LogGroupId string `json:"log_group_id" required:"true"`
	// Specifies the ID of the log subscription topic.
	// This parameter is available for all services other than ELB.
	LogStreamId string `json:"log_topic_id" required:"true"`
}

func Create(c *golangsdk.ServiceClient, opts CreateOpts) (*Log, error) {
	b, err := build.RequestBody(opts, "logtank")
	if err != nil {
		return nil, err
	}
	// POST /v3/{project_id}/elb/logtanks
	raw, err := c.Post(c.ServiceURL("logtanks"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res Log
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Log struct {
	// Provides supplementary information.
	Logtank *Logtank `json:"logtank"`
	// Specifies the request ID. The value is automatically generated.
	RequestId string `json:"request_id"`
}

type Logtank struct {
	// Specifies the log ID.
	ID string `json:"id"`
	// Specifies the ID of a load balancer.
	ProjectId string `json:"project_id"`
	// Specifies the ID of a load balancer.
	LoadbalancerId string `json:"loadbalancer_id"`
	// Specifies the log group ID.
	LogGroupId string `json:"log_group_id"`
	// Specifies the log stream ID.
	LogStreamId string `json:"log_topic_id"`
}
