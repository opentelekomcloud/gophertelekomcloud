package servicegroup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Service group name.
	Name string `json:"name,omitempty"`
	// Description.
	Description string `json:"description,omitempty"`
}

// This function is used to update service group information.
// groupId: Service group ID. It is the same as ID retuned while creating an service group.
func UpdateServiceGroup(client *golangsdk.ServiceClient, groupId string, opts UpdateOpts) (*ServiceSetId, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/service-sets/{set_id}
	raw, err := client.Put(client.ServiceURL("service-sets", groupId), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res ServiceSetDataResponse
	return &res.Data, extract.Into(raw.Body, &res)
}
