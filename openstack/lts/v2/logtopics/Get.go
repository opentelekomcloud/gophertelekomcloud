package logtopics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type GetOpts struct {
	// Log group ID
	GroupId string `json:"group_id"`
	// ID of a log stream
	TopicId string `json:"topic_id"`
}

func Get(client *golangsdk.ServiceClient, opts GetOpts) (*GetResponse, error) {
	// GET /v2.0/{project_id}/log-groups/{group_id}/log-topics/{topic_id}
	raw, err := client.Get(client.ServiceURL("log-groups", opts.GroupId, "log-topics", opts.TopicId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GetResponse struct {
	// Log stream name
	LogTopicName string `json:"log_topic_name"`
	// Log group creation time
	CreationTime int64 `json:"creation_time"`
	// Search switch
	IndexEnabled bool `json:"index_enabled"`
}
