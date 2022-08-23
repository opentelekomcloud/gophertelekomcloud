package groups

import "github.com/opentelekomcloud/gophertelekomcloud"

type ActionOpts struct {
	Action string `json:"action" required:"true"`
}

func Enable(client *golangsdk.ServiceClient, id string) error {
	opts := ActionOpts{
		Action: "resume",
	}
	return doAction(client, id, opts)
}

func Disable(client *golangsdk.ServiceClient, id string) error {
	opts := ActionOpts{
		Action: "pause",
	}
	return doAction(client, id, opts)
}

func doAction(client *golangsdk.ServiceClient, id string, opts ActionOpts) error {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("scaling_group", id, "action"), &b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
