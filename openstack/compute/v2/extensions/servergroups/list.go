package servergroups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// List returns a Pager that allows you to iterate over a collection of ServerGroups.
func List(client *golangsdk.ServiceClient) ([]ServerGroup, error) {
	raw, err := client.Get(client.ServiceURL("os-server-groups"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []ServerGroup
	err = extract.IntoSlicePtr(raw.Body, &res, "server_groups")
	return res, err
}
