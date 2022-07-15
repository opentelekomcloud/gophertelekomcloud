package tags

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/vaults"
)

type Tag struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type TagResult struct {
	golangsdk.Result
}

func (r TagResult) Extract() ([]Tag, error) {
	var s struct {
		Tags []Tag `json:"tags"`
	}
	if r.Err != nil {
		return nil, r.Err
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return nil, fmt.Errorf("failed to extract Vault Project Tags")
	}
	return s.Tags, nil
}

// ----------------------------------------------------------------------------

type InstancesBody struct {
	Resources  []TagResource `json:"resources"`
	TotalCount int           `json:"total_count"`
}

type TagResource struct {
	ResourceID     string  `json:"resource_id"`
	ResourceDetail []Vault `json:"resource_detail"`
	Tags           []Tag   `json:"tags"`
	ResourceName   string  `json:"resource_name"`
	SysTags        []Tag   `json:"sys_tags"`
}

type Vault struct {
	vaults.Vault
}

type InstancesResult struct {
	golangsdk.Result
}

func (r InstancesResult) Extract() (*InstancesBody, error) {
	var s = InstancesBody{}
	if r.Err != nil {
		return nil, r.Err
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return nil, fmt.Errorf("failed to extract Vault Resource Instances")
	}
	return &s, nil
}
