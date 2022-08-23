package groups_hcs

import "github.com/opentelekomcloud/gophertelekomcloud"

// Enable is an operation by which can make the group enable service
func Enable(client *golangsdk.ServiceClient, id string) (r ActionResult) {
	opts := ActionOpts{
		Action: "resume",
	}
	return doAction(client, id, opts)
}
