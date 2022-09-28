package tag

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListInstancesByTagsOpts struct {
	// Index position. The query starts from the next piece of data indexed by this parameter.
	// If action is set to count, this parameter is not transferred.
	// If action is set to filter, this parameter must be a positive integer. The default value is 0,
	// indicating that the query starts from the first piece of data.
	Offset string `json:"offset,omitempty"`
	// Number of records to be queried. l If action is set to count, this parameter is not transferred.
	// If action is set to filter, the value range is from 1 to 100. If this parameter is not transferred,
	// the first 100 instances are queried by default.
	Limit string `json:"limit,omitempty"`
	// Specifies the operation identifier.
	// If action is set to filter, instances are queried by tag filtering criteria.
	// If action is set to count, only the total number of records is returned.
	Action string `json:"action"`
	// Field to be matched.
	// If the value is left blank, the query is not based on the instance name or instance ID.
	// If the value is not empty
	Matches []MatchOption `json:"matches,omitempty"`
	// Specifies the included tags. Each tag contains up to 20 keys.
	Tags []TagOption `json:"tags,omitempty"`
}

type TagOption struct {
	// Tag key. It contains a maximum of 36 Unicode characters. key cannot be empty, an empty string, or spaces.
	// Before using key, delete spaces of single-byte character (SBC) before and after the value.
	// NOTE The character set of this parameter is not verified in the search process.
	Key string `json:"key"`
	// Lists the tag values. Each value contains a maximum of 43 Unicode characters and cannot contain spaces.
	// Before using values, delete SBC spaces before and after the value.
	// If the values are null, it indicates querying any value. The values are in OR relationship.
	Values []string `json:"values"`
}

type MatchOption struct {
	// Query criteria. The value can be instance_name or instance_id,
	// indicating that the query is based on the instance name or instance ID.
	Key string `json:"key"`
	// The name or ID of the instance to be matched.
	Value string `json:"value"`
}

func ListInstancesByResourceTags(client *golangsdk.ServiceClient, opts ListInstancesByTagsOpts) (*ListInstancesByResourceTagsResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/v3/{project_id}/instances/resource_instances/action
	raw, err := client.Post(client.ServiceURL("instances", "resource_instances", "action"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListInstancesByResourceTagsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListInstancesByResourceTagsResponse struct {
	Instances  []InstanceResult `json:"instances,omitempty"`
	TotalCount int32            `json:"total_count,omitempty"`
}

type InstanceResult struct {
	// Instance ID
	InstanceId string `json:"instance_id"`
	// Instance name
	InstanceName string `json:"instance_name"`
	// Tag list. If there is no tag in the list, tags is taken as an empty array
	Tags []InstanceTagResult `json:"tags"`
}

type InstanceTagResult struct {
	// Tag key. The value contains 36 Unicode characters and cannot be blank. Character set: 0-9, A-Z, a-z, "_", and "-".
	Key string `json:"key,omitempty"`
	// Tag value. The value contains a maximum of 43 Unicode characters and can also be an empty string.
	// Character set: 0-9, A-Z, a-z, "_", and "-".
	Value string `json:"value,omitempty"`
}
