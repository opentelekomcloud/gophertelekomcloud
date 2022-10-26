package tags

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ActionOpts struct {
	// Specifies the tag list.
	// If action is set to delete, the tag structure cannot be missing, and the key cannot be left blank or an empty string.
	Tags []ResourceTag `json:"tags" required:"true"`
	// Operation ID (case sensitive)
	// delete: indicates deleting a tag.
	// create: indicates creating a tag. If the same key value already exists, it will be overwritten.
	Action string `json:"action" required:"ture"`
}

func doAction(client *golangsdk.ServiceClient, id string, opts ActionOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	_, err = client.Post(client.ServiceURL("scaling_group_tag", id, "tags/action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 204},
	})
	return
}

func Create(client *golangsdk.ServiceClient, id string, tags []ResourceTag) (err error) {
	opts := ActionOpts{
		Tags:   tags,
		Action: "create",
	}
	return doAction(client, id, opts)
}

func Update(client *golangsdk.ServiceClient, id string, tags []ResourceTag) (err error) {
	opts := ActionOpts{
		Tags:   tags,
		Action: "update",
	}
	return doAction(client, id, opts)
}

func Delete(client *golangsdk.ServiceClient, id string, tags []ResourceTag) (err error) {
	opts := ActionOpts{
		Tags:   tags,
		Action: "delete",
	}
	return doAction(client, id, opts)
}

type ResourceTag struct {
	// Specifies the resource tag key. Tag keys of a resource must be unique.
	// A tag key contains a maximum of 36 Unicode characters and cannot be left blank.
	// It can contain only digits, letters, hyphens (-), underscores (_), and at signs (@).
	// When action is set to delete, the tag character set is not verified, and a key contains a maximum of 127 Unicode characters.
	Key string `json:"key" required:"ture"`
	// Specifies the resource tag value.
	// A tag value contains a maximum of 43 Unicode characters and can be left blank.
	// It can contain only digits, letters, hyphens (-), underscores (_), and at signs (@).
	// When action is set to delete, the tag character set is not verified, and a value contains a maximum of 255 Unicode characters.
	// If value is specified, tags are deleted by key and value. If value is not specified, tags are deleted by key.
	Value string `json:"value,omitempty"`
}
