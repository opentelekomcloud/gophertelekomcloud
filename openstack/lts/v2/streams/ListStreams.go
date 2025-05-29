package streams

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListStreamsOpts struct {
	GroupName  string `q:"log_group_name,omitempty"`
	StreamName string `q:"log_stream_name,omitempty"`
}

func ListStreams(client *golangsdk.ServiceClient, opts ListStreamsOpts) ([]LogStream, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("log-streams").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	// GET /v2/{project_id}/log-streams
	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"content-type": "application/json",
		},
	})
	if err != nil {
		return nil, err
	}

	var res []LogStream
	err = extract.IntoSlicePtr(raw.Body, &res, "log_streams")
	return res, err
}
