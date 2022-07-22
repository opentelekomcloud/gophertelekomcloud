package tags

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/vaults"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type InstancesResponse struct {
	Resources  []TagResource `json:"resources"`
	TotalCount int           `json:"total_count"`
}

type TagResource struct {
	ResourceID     string             `json:"resource_id"`
	ResourceDetail BoxedVault         `json:"resource_detail"`
	Tags           []tags.ResourceTag `json:"tags"`
	ResourceName   string             `json:"resource_name"`
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
