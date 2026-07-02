package host_groups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type DeleteOpts struct {
	HostGroupIds []string `json:"host_group_id_list" required:"true"`
}

// Delete a host group by id
func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) (*DeleteResult, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// DELETE /v3/{project_id}/lts/host-group
	raw, err := client.DeleteWithBody(client.ServiceURL("lts", "host-group"), b, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"content-type": "application/json",
		},
		OkCodes: []int{200, 204},
	})
	if err != nil {
		return nil, err
	}

	var res DeleteResult
	err = extract.Into(raw.Body, &res)
	return &res, err
}

// DeleteResult represents the API response after deletion
type DeleteResult struct {
	// Host group details.
	Result []HostGroupResponse `json:"result"`
	// Number of deleted host groups.
	Total int64 `json:"total"`
}
