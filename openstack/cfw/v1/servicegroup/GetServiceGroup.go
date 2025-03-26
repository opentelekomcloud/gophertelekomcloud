package servicegroup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to query service group details.
// groupId: Service group ID. It is the same as ID retuned while creating an service group.
func GetServiceGroup(client *golangsdk.ServiceClient, groupId string) (*ServiceSetDetailResponseDto, error) {
	// GET /v1/{project_id}/service-sets/{set_id}
	raw, err := client.Get(client.ServiceURL("service-sets", groupId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetResponse
	err = extract.Into(raw.Body, &res)
	return &res.Data, err
}

type GetResponse struct {
	Data ServiceSetDetailResponseDto `json:"data"`
}

type ServiceSetDetailResponseDto struct {
	// Service Group ID
	ID string `json:"id"`
	// Service Group Name
	Name string `json:"name"`
	// Description
	Description string `json:"description"`
	// Service group type:
	// 0 - user-defined service group
	// 1 - common web service
	// 2 - common remote login and ping
	// 3 - common database
	ServiceSetType int `json:"service_set_type"`
}
