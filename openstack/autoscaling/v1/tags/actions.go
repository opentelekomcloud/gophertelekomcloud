package tags

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

type ActionOpts struct {
	Tags   []ResourceTag `json:"tags" required:"true"`
	Action string        `json:"action" required:"ture"`
}

func doAction(client *golangsdk.ServiceClient, id string, opts ActionOpts) (err error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
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
	Key   string `json:"key" required:"ture"`
	Value string `json:"value,omitempty"`
}
