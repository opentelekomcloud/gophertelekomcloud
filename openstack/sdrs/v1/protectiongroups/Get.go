package protectiongroups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get retrieves a particular Group based on its unique ID.
func Get(client *golangsdk.ServiceClient, ServerGroupId string) (*ServerGroupResponseInfo, error) {
	// GET /v1/{project_id}/server-groups/{server_group_id}
	raw, err := client.Get(client.ServiceURL("server-groups", ServerGroupId), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}
	var res GetResponse
	err = extract.Into(raw.Body, &res)
	return &res.ServerGroup, err
}

type GetResponse struct {
	// Specifies the information about a protection group.
	ServerGroup ServerGroupResponseInfo `json:"server_group"`
	// Defined in Update.go
}
