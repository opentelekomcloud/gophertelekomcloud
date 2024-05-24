package script

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get is used to query a script, including the script type and script content.
// Send request GET /v1/{project_id}/scripts/{script_name}
func Get(client *golangsdk.ServiceClient, scriptName, workspace string) (*Script, error) {
	// strange behaviour, because workspace in the documentation is not required
	if workspace == "" {
		return nil, fmt.Errorf("workspace is required")
	}

	opts := &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{HeaderWorkspace: workspace},
	}

	raw, err := client.Get(client.ServiceURL(scriptsEndpoint, scriptName), nil, opts)
	if err != nil {
		return nil, err
	}

	var res Script
	err = extract.Into(raw.Body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
