package host_groups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type ListOpts struct {
	// List of host group IDs. The host type must be the same as the host group type.
	HostIdList []string `json:"host_id_list,omitempty"`
	// Host group filters.
	Filter *Filter `json:"filter,omitempty"`
}

type Filter struct {
	// Host group type.
	// Windows
	// Linux
	Type string `json:"host_group_type,omitempty"`
	// List of host group names.
	HostGroupNameList []string `json:"host_group_name_list,omitempty"`
	// Host name list.
	HostNameList []string `json:"host_name_list,omitempty"`
	// Host group tags.
	HostGroupTag []ListTag `json:"host_group_tag,omitempty"`
}

type ListTag struct {
	// Tag type. Tag filtering logic: `AND` or `OR`.
	Type string `json:"tag_type,omitempty"`
	// Host group tags.
	Tags []tags.ResourceTag `json:"tag_list,omitempty"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListResult, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	// POST /v3/{project_id}/lts/host-group-list
	raw, err := client.Post(client.ServiceURL("lts", "host-group-list"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ListResult
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListResult struct {
	// Host group details.
	Result []HostGroupResponse `json:"result"`
	// Number of deleted host groups.
	Total int64 `json:"total"`
}
