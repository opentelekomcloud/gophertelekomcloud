package groups

import "github.com/opentelekomcloud/gophertelekomcloud"

type ActionOpts struct {
	Action string `json:"action" required:"true"`
}

func Disable(client *golangsdk.ServiceClient, id string) error {
	opts := ActionOpts{
		Action: "pause",
	}
	return doAction(client, id, opts)
}
