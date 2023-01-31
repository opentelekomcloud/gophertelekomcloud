package tags

import "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"

type ListImageByTagsOpts struct {
	// Identifies the operation. This parameter is case sensitive and its value can be filter or count.
	//
	// The value filter indicates pagination query.
	//
	// The value count indicates that the total number of query results meeting the search criteria will be returned.
	Action string `json:"action" required:"true"`
	// Includes all specified tags. A maximum of 10 tag keys are allowed for each query operation. Each tag key can contain a maximum of 10 tag values. Both tag keys and values must be unique. The tag keys cannot be left blank.
	Tags []tags.ListedTag `json:"tags,omitempty"`
	// Includes any of specified tags. A maximum of 10 tag keys are allowed for each query operation. Each tag key can contain a maximum of 10 tag values. Both tag keys and values must be unique. The tag keys cannot be left blank and set to an empty string.
	TagsAny []tags.ListedTag `json:"tags_any,omitempty"`
	// Excludes all specified tags. A maximum of 10 tag keys are allowed for each query operation. Each tag key can contain a maximum of 10 tag values. Both tag keys and values must be unique. The tag keys cannot be left blank.
	NotTags []tags.ListedTag `json:"not_tags,omitempty"`
	// Excludes any of specified tags. A maximum of 10 tag keys are allowed for each query operation. Each tag key can contain a maximum of 10 tag values. Both tag keys and values must be unique. The tag keys cannot be left blank.
	NotTagsAny []tags.ListedTag `json:"not_tags_any,omitempty"`
	// If this parameter is set to true, all resources without tags are queried. In this case, the tag, not_tags, tags_any, and not_tags_any fields are ignored.
	WithoutAnyTag bool `json:"without_any_tag,omitempty"`
	// Specifies the maximum number of query records. This parameter is invalid when action is set to count. If action is set to filter, the parameter limit takes effect and its default value is 10. The value of limit ranges from 1 to 1000.
	Limit string `json:"limit,omitempty"`
	// Specifies the index position. The query starts from the next image indexed by this parameter. This parameter is not required when data on the first page is queried, and it is invalid when action is set to count. If action is set to filter, the default value of offset is 0. The value of offset cannot be a negative number.
	Offset string `json:"offset,omitempty"`
	// Specifies the search criteria. The tag key is the field to match, for example, resource_name or resource_id. value indicates the matched value. Keys in this list must be unique. The parameter cannot be left blank and may not be transferred.
	Matches []tags.ResourceTag `json:"matches,omitempty"`
}

// POST /v2/{project_id}/images/resource_instances/action

// 200
