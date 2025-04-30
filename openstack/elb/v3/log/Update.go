package log

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	// Specifies the log group ID.
	// This parameter is available for all services other than ELB.
	LogGroupId string `json:"log_group_id,omitempty"`
	// Specifies the ID of the log subscription topic.
	// This parameter is available for all services other than ELB.
	LogStreamId string `json:"log_topic_id,omitempty"`
}

func Update(c *golangsdk.ServiceClient, id string, opts UpdateOpts) (err error) {
	b, err := build.RequestBody(opts, "logtank")
	if err != nil {
		return
	}

	// PUT /v3/{project_id}/elb/logtanks/{logtank_id}
	_, err = c.Put(c.ServiceURL("logtanks", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	return
}
