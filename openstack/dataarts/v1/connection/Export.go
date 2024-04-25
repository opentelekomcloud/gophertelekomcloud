package connection

import (
	"io"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

const exportEndpoint = "export"

// Export is used to export all connection information that is compressed in ZIP format.
// Send request POST /v1/{project_id}/connections/export
func Export(client *golangsdk.ServiceClient, workspace string) (io.ReadCloser, error) {

	opts := &golangsdk.RequestOpts{}
	if workspace != "" {
		opts.MoreHeaders[HeaderWorkspace] = workspace
	}
	raw, err := client.Post(client.ServiceURL(connectionsEndpoint, exportEndpoint), nil, nil, opts)
	if err != nil {
		return nil, err
	}

	return raw.Body, err
}
