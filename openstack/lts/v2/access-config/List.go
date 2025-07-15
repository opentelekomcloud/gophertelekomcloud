package access_config

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type ListOpts struct {
	// List of ingestion configuration names.
	AccessConfigNames []string `json:"access_config_name_list"`
	// List of host group names.
	HostGroupNames []string `json:"host_group_name_list"`
	// List of log group names.
	LogGroupNames []string `json:"log_group_name_list"`
	// List of log stream names.
	LogStreamNames []string `json:"log_stream_name_list"`
	// Ingestion configuration tags. A tag key must be unique. Up to 20 tags are allowed.
	Tags []tags.ResourceTag `json:"access_config_tag_list,omitempty"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListResult, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	// POST /v3/{project_id}/lts/access-config-list
	raw, err := client.Post(client.ServiceURL("lts", "access-config-list"), b, nil, &golangsdk.RequestOpts{
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
	// Ingestion configuration list.
	Result []AccessConfigInfo `json:"result"`
	// Total number of ingestion configurations.
	Total int64 `json:"total"`
}
