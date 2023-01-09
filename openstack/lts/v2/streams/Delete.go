package streams

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

type DeleteOpts struct {
	// ID of a created log group
	GroupId string
	// ID of a created log stream
	TopicId string
}

// Delete a log topic by id
func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) (err error) {
	// DELETE /v2.0/{project_id}/log-groups/{group_id}/log-topics/{topic_id}
	_, err = client.Delete(client.ServiceURL("log-groups", opts.GroupId, "log-topics", opts.TopicId), nil)
	return
}
