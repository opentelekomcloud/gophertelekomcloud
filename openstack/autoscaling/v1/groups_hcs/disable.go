package groups_hcs

import "github.com/opentelekomcloud/gophertelekomcloud"

// Disable is an operation by which can be able to pause the group
func Disable(client *golangsdk.ServiceClient, id string) (r ActionResult) {
	opts := ActionOpts{
		Action: "pause",
	}
	return doAction(client, id, opts)
}
