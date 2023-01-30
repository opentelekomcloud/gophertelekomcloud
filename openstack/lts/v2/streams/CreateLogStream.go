package streams

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// ID of a created log group
	GroupId string `json:"-" required:"true"`
	// Name of the log stream to be created.
	// Minimum length: 1 character
	// Maximum length: 64 characters
	// Enumerated value:
	// lts-stream-13ci
	LogStreamName string `json:"log_stream_name" required:"true"`
}

func CreateLogStream(client *golangsdk.ServiceClient, opts CreateOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	// POST /v2/{project_id}/groups/{log_group_id}/streams
	raw, err := client.Post(client.ServiceURL("groups", opts.GroupId, "streams"), b, nil, nil)
	if err != nil {
		return "", err
	}

	var res struct {
		ID string `json:"log_stream_id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.ID, err
}
