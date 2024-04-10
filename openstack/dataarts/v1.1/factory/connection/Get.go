package connection

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get is used to query configuration details of a specific connection.
// Send request GET /v1/{project_id}/connections/{connection_name}
func Get(client *golangsdk.ServiceClient, connectionName string, workspace string) (*Config, error) {
	var opts *golangsdk.RequestOpts
	if workspace != "" {
		opts = &golangsdk.RequestOpts{
			MoreHeaders: map[string]string{HeaderWorkspace: workspace},
		}
	}
	raw, err := client.Get(client.ServiceURL("clusters", connectionName), nil, opts)
	if err != nil {
		return nil, err
	}

	var res *Config
	err = extract.Into(raw.Body, res)
	return res, err
}
