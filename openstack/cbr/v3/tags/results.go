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

type InstancesResponse struct {
	Resources  []TagResource `json:"resources"`
	TotalCount int           `json:"total_count"`
}

type TagResource struct {
	ResourceID     string       `json:"resource_id"`
	ResourceDetail BoxedVault   `json:"resource_detail"`
	Tags           []vaults.Tag `json:"tags"`
	ResourceName   string       `json:"resource_name"`
}

type BoxedVault struct {
	Vault vaults.Vault `json:"vault"`
}

type InstancesResult struct {
	golangsdk.Result
}

func (r InstancesResult) Extract() (*InstancesResponse, error) {
	var s = InstancesResponse{}
	if r.Err != nil {
		return nil, r.Err
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return nil, fmt.Errorf("failed to extract Vault Resource Instances")
	}
	return &s, nil
}

// ----------------------------------------------------------------------------

type ShowVaultTagResult struct {
	golangsdk.Result
}

type ShowVaultTagResponse struct {
	Tags []vaults.Tag `json:"tags"`
}

func (r ShowVaultTagResult) Extract() (*ShowVaultTagResponse, error) {
	var s = ShowVaultTagResponse{}
	if r.Err != nil {
		return nil, r.Err
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return nil, fmt.Errorf("failed to extract Vault Tag")
	}
	return &s, nil
}
