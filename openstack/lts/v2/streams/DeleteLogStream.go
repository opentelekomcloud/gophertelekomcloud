package streams

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

type DeleteOpts struct {
	// ID of a created log group
	GroupId string
	// ID of a created log stream
	StreamId string
}

// DeleteLogStream a log topic by id
func DeleteLogStream(client *golangsdk.ServiceClient, opts DeleteOpts) (err error) {
	// DELETE /v2/{project_id}/groups/{log_group_id}/streams/{log_stream_id}
	_, err = client.Delete(client.ServiceURL("groups", opts.GroupId, "streams", opts.StreamId), &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"content-type": "application/json",
		},
	})
	return
}
