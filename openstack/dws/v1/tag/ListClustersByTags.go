package tag

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type ListResourceReq struct {
	// The resources to be queried contain tags listed in tags. Each resource to be queried contains a maximum of 10 keys.
	// Each tag key can have a maximum of 10 tag values. The tag value corresponding to each tag key can be
	// an empty array but the structure cannot be missing. Each tag key must be unique, and each tag value in a tag must be unique.
	// The response returns resources containing all tags in this list. Keys in this list are in an
	// AND relationship while values in each key-value structure are in an OR relationship.
	// If no tag filtering condition is specified, full data is returned.
	Tags []TagWithMultiValue `json:"tags,omitempty"`
	// The resources to be queried contain any tags listed in tags_any. Each resource to be queried contains a maximum of 10 keys.
	// Each tag key can have a maximum of 10 tag values.
	// The tag value corresponding to each tag key can be an empty array but the structure cannot be missing.
	// Each tag key must be unique, and each tag value in a tag must be unique. The response returns resources containing the tags in this list.
	// Keys in this list are in an OR relationship and values in each key-value structure are also in an OR relationship.
	// If no tag filtering condition is specified, full data is returned.
	TagsAny []TagWithMultiValue `json:"tags_any,omitempty"`
	// The resources to be queried do not contain tags listed in not_tags. Each resource to be queried contains a maximum of 10 keys.
	// Each tag key can have a maximum of 10 tag values. The tag value corresponding to each tag key can be an empty array but the structure cannot be missing.
	// Each tag key must be unique, and each tag value in a tag must be unique. The response returns resources containing no tags in this list.
	// Keys in this list are in an AND relationship while values in each key-value structure are in an OR relationship.
	// If no tag filtering condition is specified, full data is returned.
	NotTags []TagWithMultiValue `json:"not_tags,omitempty"`
	// The resources to be queried do not contain any tags listed in not_tags_any.
	// Each resource to be queried contains a maximum of 10 keys. Each tag key can have a maximum of 10 tag values.
	// The tag value corresponding to each tag key can be an empty array but the structure cannot be missing.
	// Each tag key must be unique, and each tag value in a tag must be unique. The response returns resources containing no tags in this list.
	// Keys in this list are in an OR relationship and values in each key-value structure are also in an OR relationship.
	// If no tag filtering condition is specified, full data is returned.
	NotTagsAny []TagWithMultiValue `json:"not_tags_any,omitempty"`
	// Identifies the operation. The value can be filtered or count.
	// filter: indicates filtering. When both limit and offset are configured, the returned results are displayed in pages.
	// If both limit and offset are not configured, the returned results are displayed in pages only when the number of result records exceeds 1000.
	// count indicates the total number of returned records that meet the query criteria.
	Action string `json:"action"`
	// Maximum number of records returned to the query result. This parameter is not displayed when action is set to count.
	// If action is set to filter, this parameter takes effect. Its value ranges from 1 to 1000 (default).
	Limit int `json:"limit,omitempty"`
	// Start location of pagination query. The query starts from the next resource of the specified location.
	// When querying the data on the first page, you do not need to specify this parameter.
	// When querying the data on subsequent pages, set this parameter to the value in the response body returned by querying data of the previous page.
	// This parameter is not displayed when action is set to count. If action is set to filter, this parameter takes effect.
	// Its value can be 0 (default) or a positive integer.
	Offset int `json:"offset,omitempty"`
	// Search field. key indicates the field to be matched, for example, resource_name. value indicates the fuzzy match result.
	Matches []Match `json:"matches,omitempty"`
}

type Match struct {
	// Key. Currently, it can only be resource_name.
	Key string `json:"key,omitempty"`
	// Key value. Each value can contain a maximum of 255 Unicode characters.
	Value string `json:"value,omitempty"`
}

func ListClustersByTags(client *golangsdk.ServiceClient, opts ListResourceReq) (*ListClustersByTagsResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1.0/{project_id}/clusters/resource_instances/action
	raw, err := client.Get(client.ServiceURL("clusters", "resource_instances", "action"), b, nil)
	if err != nil {
		return nil, err
	}

	var res ListClustersByTagsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListClustersByTagsResponse struct {
	// Resources that meet the search criteria.
	Resources []MrsResource `json:"resources,omitempty"`
	// Total number of queried records.
	TotalCount int `json:"total_count,omitempty"`
}

type MrsResource struct {
	// Resource ID.
	ResourceId string `json:"resource_id,omitempty"`
	// Resource details. The value is a resource object, used for extension. This value is left empty by default.
	ResourceDetail string `json:"resource_detail,omitempty"`
	// List of tags. If no tag is matched, an empty array is returned.
	Tags []tags.ResourceTag `json:"tags,omitempty"`
	// Resource name. This parameter is an empty string by default if the resource name is not specified.
	ResourceName string `json:"resource_name,omitempty"`
}
