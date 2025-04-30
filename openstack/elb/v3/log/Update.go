package log

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Specifies the log group ID.
	// This parameter is available for all services other than ELB.
	LogGroupId string `json:"log_group_id,omitempty"`
	// Specifies the ID of the log subscription topic.
	// This parameter is available for all services other than ELB.
	LogStreamId string `json:"log_topic_id,omitempty"`
}

func Update(c *golangsdk.ServiceClient, id string, opts UpdateOpts) (*Log, error) {
	b, err := build.RequestBody(opts, "logtank")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/elb/logtanks/{logtank_id}
	raw, err := c.Put(c.ServiceURL("logtanks", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	var res Log
	err = extract.Into(raw.Body, &res)
	return &res, err
}
