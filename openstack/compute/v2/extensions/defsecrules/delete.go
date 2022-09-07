package defsecrules

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a rule the project's default security group.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	raw, err := client.Delete(resourceURL(client, id), nil)
	return
}
