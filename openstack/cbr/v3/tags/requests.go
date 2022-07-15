package tags

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

/*
ShowVaultProjectTag
Query the set of all tags of the tenant in the specified Region and instance type

@author Aloento
@since 0.4.17
@version 0.1.0
*/
func ShowVaultProjectTag(client *golangsdk.ServiceClient) (r TagResult) {
	_, r.Err = client.Get(showVaultProjectTagURL(client), &r.Body, nil)
	return
}

// ----------------------------------------------------------------------------

type ActionType string

const (
	Filter = "filter"
	Count  = "count"
)

type CloudType string

const (
	Public = "public"
	Hybrid = "hybrid"
)

type ObjectType string

const (
	Server = "server"
	Disk   = "disk"
)

type ResourceInstancesRequest struct {
	// Does not contain any of the tags?
	WithoutAnyTag bool `json:"without_any_tag,omitempty"`
	// Returns the full amount of data when there are no filter conditions
	Tags []Tag `json:"tags,omitempty"`
	// Include any tag
	TagsAny []Tag `json:"tags_any,omitempty"`
	// Label not included
	NotTags []Tag `json:"not_tags,omitempty"`
	// Does not contain any tags
	NotTagsAny []Tag `json:"not_tags_any,omitempty"`
	// Only the op_service permission can use this field as a resource instance filter condition
	SysTags []Tag `json:"sys_tags,omitempty"`
	// Number of search records, default is 1000, the minimum value of limit is 1, the maximum value of limit is 1000
	Limit string `json:"limit,omitempty"`
	// Index position (no this parameter when action is count)
	Offset string `json:"offset,omitempty"`
	// filter is a paginated query. count simply returns the total number of items according to the criteria
	Action ActionType `json:"action"`
	// Query conditions supported by the resource itself
	Matches    []Tag      `json:"matches,omitempty"`
	CloudType  CloudType  `json:"cloud_type,omitempty"`
	ObjectType ObjectType `json:"object_type,omitempty"`
}

/*
ShowVaultResourceInstances
Use tags to filter instances.

@author Aloento
@since 0.4.17
@version 0.1.0
*/
func ShowVaultResourceInstances(client *golangsdk.ServiceClient, req ResourceInstancesRequest) (r InstancesResult) {
	reqBody, err := golangsdk.BuildRequestBody(req, "")
	if err != nil {
		r.Err = fmt.Errorf("failed to create vault create map: %s", err)
		return
	}
	_, err = client.Post(showVaultResourceInstancesURL(client), reqBody, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	r.Err = err
	return
}

// ----------------------------------------------------------------------------
