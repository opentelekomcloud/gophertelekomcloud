package tags

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// ActionOptsBuilder is an interface from which can build the request of creating/deleting tags
type ActionOptsBuilder interface {
	ToTagsActionMap() (map[string]interface{}, error)
}

// ResourceTag is in key-value format
type ResourceTag struct {
	Key   string `json:"key" required:"true"`
	Value string `json:"value,omitempty"`
}

// ActionOpts is a struct contains the parameters of creating/deleting tags
type ActionOpts struct {
	Tags   []ResourceTag `json:"tags" required:"true"`
	Action string        `json:"action" required:"true"`
}

// ToTagsActionMap build the action request in json format
func (opts ActionOpts) ToTagsActionMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func doAction(client *golangsdk.ServiceClient, serviceType, id string, opts ActionOptsBuilder) (r ActionResult) {
	b, err := opts.ToTagsActionMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(actionURL(client, serviceType, id), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 204},
		MoreHeaders: openstack.StdRequestOpts().MoreHeaders,
	})
	return
}

// Create is a method of creating tags by id
func Create(client *golangsdk.ServiceClient, serviceType, id string, tags []ResourceTag) (r ActionResult) {
	opts := ActionOpts{
		Tags:   tags,
		Action: "create",
	}
	return doAction(client, serviceType, id, opts)
}

// Delete is a method of deleting tags by id
func Delete(client *golangsdk.ServiceClient, serviceType, id string, tags []ResourceTag) (r ActionResult) {
	opts := ActionOpts{
		Tags:   tags,
		Action: "delete",
	}
	return doAction(client, serviceType, id, opts)
}

// Get is a method of getting the tags by id
func Get(client *golangsdk.ServiceClient, serviceType, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, serviceType, id), &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{202, 200},
		MoreHeaders: openstack.StdRequestOpts().MoreHeaders,
	})
	return
}

// List is a method of getting the tags of all service
func List(client *golangsdk.ServiceClient, serviceType string) (r ListResult) {
	_, r.Err = client.Get(listURL(client, serviceType), &r.Body, openstack.StdRequestOpts())
	return
}
