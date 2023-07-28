package groups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ActionOpts struct {
	Action string `json:"action" required:"true"`
}

func Enable(client *golangsdk.ServiceClient, id string) error {
	return doAction(client, id, ActionOpts{
		Action: "resume",
	})
}

func Disable(client *golangsdk.ServiceClient, id string) error {
	return doAction(client, id, ActionOpts{
		Action: "pause",
	})
}

func doAction(client *golangsdk.ServiceClient, id string, opts ActionOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("scaling_group", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
