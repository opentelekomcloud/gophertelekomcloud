package connection

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// Delete is used to delete a connection.
// Send request DELETE /v1/{project_id}/connections/{connection_name}
func Delete(client *golangsdk.ServiceClient, connName, workspace string) error {

	reqOpts := &golangsdk.RequestOpts{
		OkCodes: []int{204},
	}

	if workspace != "" {
		reqOpts = &golangsdk.RequestOpts{
			MoreHeaders: map[string]string{HeaderWorkspace: workspace},
		}
	}

	_, err := client.Delete(client.ServiceURL(connectionsEndpoint, connName), reqOpts)

	return err
}
