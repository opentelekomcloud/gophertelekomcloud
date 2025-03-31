package host_groups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// CreateOpts is a struct that contains all the parameters.
type CreateOpts struct {
	// Host group name.
	// Use only letters, digits, underscores (_), hyphens (-), and periods (.).
	// Do not start with a period or underscore or end with a period.
	Name string `json:"host_group_name" required:"true"`
	// Host group type.
	// Windows
	// Linux
	Type string `json:"host_group_type" required:"true"`
	// List of host group IDs. The host type must be the same as the host group type.
	HostIdList []string `json:"host_id_list,omitempty"`
	// Host access type.
	// LABEL
	// IP
	AgentAccessType string `json:"agent_access_type,omitempty"`
	// Host group identifier. If the host access type is LABEL, this field saves the host group identifier.
	Labels []string `json:"labels,omitempty"`
	// Tag information. You can add up to 20 tags.
	Tags []tags.ResourceTag `json:"host_group_tag,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*HostGroupResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/lts/host-group
	raw, err := client.Post(client.ServiceURL("lts", "host-group"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	if err != nil {
		return nil, err
	}

	var res HostGroupResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type HostGroupResponse struct {
	// Host group ID.
	ID string `json:"host_group_id"`
	// Host group name.
	Name string `json:"host_group_name"`
	// Host group type.
	Type string `json:"host_group_type"`
	// Host ID list.
	HostIdList []string `json:"host_id_list"`
	// Tag information.
	Tags []tags.ResourceTag `json:"host_group_tag"`
	// Creation time.
	CreatedAt int64 `json:"create_time"`
	// Update time.
	UpdatedAt int64 `json:"update_time"`
	// Host group ID.
	Labels []string `json:"labels"`
	// Host access type.
	AgentAccessType string `json:"agent_access_type"`
}
