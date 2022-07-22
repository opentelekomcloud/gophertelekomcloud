package tags

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

func ShowVaultProjectTag(client *golangsdk.ServiceClient) (r tags.ListResult) {
	return tags.List(client, baseURL)
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

func ShowVaultTag(client *golangsdk.ServiceClient, id string) (r tags.GetResult) {
	return tags.Get(client, baseURL, id)
}

// ----------------------------------------------------------------------------

func CreateVaultTags(client *golangsdk.ServiceClient, id string, req []tags.ResourceTag) (r tags.ActionResult) {
	return tags.Create(client, baseURL, id, req)
}

// ----------------------------------------------------------------------------

func DeleteVaultTag(client *golangsdk.ServiceClient, id string, req []tags.ResourceTag) (r tags.ActionResult) {
	return tags.Delete(client, baseURL, id, req)
}
