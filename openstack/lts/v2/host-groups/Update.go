package host_groups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type UpdateLogGroupOpts struct {
	// Host group ID.
	ID string `json:"host_group_id" required:"true"`
	// Host group name.
	// Use only letters, digits, underscores (_), hyphens (-), and periods (.).
	// Do not start with a period or underscore or end with a period.
	Name string `json:"host_group_name,omitempty"`
	// Host ID list. The host type must be the same as the host group type.
	HostIdList []string `json:"host_id_list,omitempty"`
	// Host group tags. A key must be unique. Up to 20 keys are allowed.
	Tags []tags.ResourceTag `json:"host_group_tag,omitempty"`
	// Host group identifier. If the host access type is LABEL, this field saves the host group identifier.
	Labels []string `json:"labels,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateLogGroupOpts) (*HostGroupResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/lts/host-group
	raw, err := client.Put(client.ServiceURL("lts", "host-group"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res HostGroupResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
