package tags

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

/*
ShowVaultProjectTag
Query the set of all tags of the tenant in the specified Region and instance type

@author Aloento
@since 0.5.17
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
	// If this parameter is set to true, all resources without tags are queried.
	WithoutAnyTag bool `json:"without_any_tag,omitempty"`
	// Returns the full amount of data when there are no filter conditions
	Tags []Tag `json:"tags,omitempty"`
	// Backups with any tags in this list will be filtered.
	TagsAny []Tag `json:"tags_any,omitempty"`
	// Backups without these tags will be filtered.
	NotTags []Tag `json:"not_tags,omitempty"`
	// Backups without any tags in this list will be filtered.
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
	Matches    []SysTag   `json:"matches,omitempty"`
	CloudType  CloudType  `json:"cloud_type,omitempty"`
	ObjectType ObjectType `json:"object_type,omitempty"`
}

/*
ShowVaultResourceInstances
Use tags to filter instances.

@author Aloento
@since 0.5.17
@version 0.1.0
*/
func ShowVaultResourceInstances(client *golangsdk.ServiceClient, req ResourceInstancesRequest) (r InstancesResult) {
	reqBody, err := golangsdk.BuildRequestBody(req, "")
	if err != nil {
		r.Err = fmt.Errorf("failed to create vault map: %s", err)
		return
	}
	_, err = client.Post(showVaultResourceInstancesURL(client), reqBody, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	r.Err = err
	return
}

// ----------------------------------------------------------------------------

/*
ShowVaultTag
Query the label information of the specified instance

@author Aloento
@since 0.5.17
@version 0.1.0
*/
func ShowVaultTag(client *golangsdk.ServiceClient, id string) (r ShowVaultTagResult) {
	_, r.Err = client.Get(vaultTagsURL(client, id), &r.Body, nil)
	return
}

// ----------------------------------------------------------------------------

/*
CreateVaultTags
Add repository resource tags

@author Aloento
@since 0.5.17
@version 0.1.0
*/
func CreateVaultTags(client *golangsdk.ServiceClient, id string, req SysTag) (r golangsdk.ErrResult) {
	reqBody, err := golangsdk.BuildRequestBody(req, "tag")
	if err != nil {
		r.Err = fmt.Errorf("failed to create vault map: %s", err)
		return
	}
	_, err = client.Post(vaultTagsURL(client, id), reqBody, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	r.Err = err
	return
}

// ----------------------------------------------------------------------------

/*
DeleteVaultTag
Delete repository resource tags

@author Aloento
@since 0.5.17
@version 0.1.0
*/
func DeleteVaultTag(client *golangsdk.ServiceClient, id string, key string) (r golangsdk.ErrResult) {
	_, err := client.Delete(deleteVaultTagURL(client, id, key), &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	r.Err = err
	return
}

// ----------------------------------------------------------------------------

type BulkCreateAndDeleteVaultTagsRequest struct {
	Tags    []SysTag       `json:"tags,omitempty"`
	SysTags []SysTag       `json:"sys_tags,omitempty"`
	Action  BulkActionType `json:"action"`
}

type BulkActionType string

const (
	Create = "create"
	Delete = "delete"
)

/*
BatchCreateAndDeleteVaultTags
Add or remove tags in bulk for specified instances

@author Aloento
@since 0.5.17
@version 0.1.0
*/
func BatchCreateAndDeleteVaultTags(client *golangsdk.ServiceClient, id string, req BulkCreateAndDeleteVaultTagsRequest) (r golangsdk.ErrResult) {
	reqBody, err := golangsdk.BuildRequestBody(req, "")
	if err != nil {
		r.Err = fmt.Errorf("failed to create vault map: %s", err)
		return
	}
	_, err = client.Post(batchCreateAndDeleteVaultTagsURL(client, id), reqBody, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	r.Err = err
	return
}
