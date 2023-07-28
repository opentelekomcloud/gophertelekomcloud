package tags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/vaults"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type ActionType string
type CloudType string
type ObjectType string

const (
	Filter = "filter"
	Count  = "count"

	Public = "public"
	Hybrid = "hybrid"

	Server = "server"
	Disk   = "disk"
)

type ResourceInstancesRequest struct {
	// If this parameter is set to true, all resources without tags are queried.
	WithoutAnyTag bool `json:"without_any_tag,omitempty"`
	// Returns the full amount of data when there are no filter conditions
	Tags []tags.ListedTag `json:"tags,omitempty"`
	// Backups with any tags in this list will be filtered.
	TagsAny []tags.ListedTag `json:"tags_any,omitempty"`
	// Backups without these tags will be filtered.
	NotTags []tags.ListedTag `json:"not_tags,omitempty"`
	// Backups without any tags in this list will be filtered.
	NotTagsAny []tags.ListedTag `json:"not_tags_any,omitempty"`
	// Number of search records, default is 1000, the minimum value of limit is 1, the maximum value of limit is 1000
	Limit string `json:"limit,omitempty"`
	// Index position (no this parameter when action is count)
	Offset string `json:"offset,omitempty"`
	// filter is a paginated query. count simply returns the total number of items according to the criteria
	Action ActionType `json:"action"`
	// Query conditions supported by the resource itself
	Matches    []tags.ResourceTag `json:"matches,omitempty"`
	CloudType  CloudType          `json:"cloud_type,omitempty"`
	ObjectType ObjectType         `json:"object_type,omitempty"`
}

func ShowVaultResourceInstances(client *golangsdk.ServiceClient, req ResourceInstancesRequest) (*InstancesResponse, error) {
	b, err := build.RequestBodyMap(req, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/vault/resource_instances/action
	raw, err := client.Post(client.ServiceURL("vault", "resource_instances", "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res InstancesResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

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
