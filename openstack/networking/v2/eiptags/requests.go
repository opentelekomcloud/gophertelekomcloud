package eiptags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToPublicIpTagCreateMap() (map[string]interface{}, error)
}

// CreateOpts is the common options struct used in this package's Create
// operation.
type CreateOpts struct {
	Tag Tag `json:"tag" required:"true"`
}

type Tag struct {
	Key   string `json:"key" required:"true"`
	Value string `json:"value"`
}

// ToPublicIpTagCreateMap casts a CreateOpts struct to a map.
func (opts CreateOpts) ToPublicIpTagCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create is a method of creating tags by eip id
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder, id string) (r CreateResult) {
	b, err := opts.ToPublicIpTagCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(rootURL(client, id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

// List is a method of getting the tags of the EIP
func List(client *golangsdk.ServiceClient, id string) (r ListResult) {
	_, r.Err = client.Get(rootURL(client, id), &r.Body, nil)
	return
}

// Delete is a method of deleting tags by key
func Delete(client *golangsdk.ServiceClient, id string, key string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, id, key), nil)
	return
}

// BatchActionOptsBuilder is an interface which can build the map parameter of action function
type BatchActionOptsBuilder interface {
	ToPublicIpBatchTagsActionMap() (map[string]interface{}, error)
}

// BatchActionOpts is the common options struct used in this package's Action
// operation.
type BatchActionOpts struct {
	Tags   []Tag  `json:"tags" required:"true"`
	Action string `json:"action" required:"true"`
}

// ToPublicIpBatchTagsActionMap casts a BatchActionOpts struct to a map.
func (opts BatchActionOpts) ToPublicIpBatchTagsActionMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Action is a method of create/delete multiple tags
func Action(client *golangsdk.ServiceClient, opts BatchActionOpts, id string) (r ActionResult) {
	b, err := opts.ToPublicIpBatchTagsActionMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(actionURL(client, id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
