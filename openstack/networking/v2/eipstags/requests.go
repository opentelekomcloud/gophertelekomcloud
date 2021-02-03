package eipstags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// CreateOptsBuilder
type CreateOptsBuilder interface {
	ToPublicIpTagCreateMap() (map[string]interface{}, error)
}

// CreateOpts
type CreateOpts struct {
	Tag Tag `json:"tag" required:"true"`
}

type Tag struct {
	Key   string `json:"key" required:"true"`
	Value string `json:"value"`
}

func (opts CreateOpts) ToPublicIpTagCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create
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

// List
func List(client *golangsdk.ServiceClient, id string) (r ListResult) {
	_, r.Err = client.Get(rootURL(client, id), &r.Body, nil)
	return
}

// Delete
func Delete(client *golangsdk.ServiceClient, id string, key string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, id, key), nil)
	return
}

// BatchActionOptsBuilder
type BatchActionOptsBuilder interface {
	ToPublicIpBatchTagsActionMap() (map[string]interface{}, error)
}

type BatchActionOpts struct {
	Tags   []Tag  `json:"tags" required:"true"`
	Action string `json:"action" required:"true"`
}

func (opts BatchActionOpts) ToPublicIpBatchTagsActionMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Action
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
