package defsecrules

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a rule the project's default security group.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	raw, err := client.Delete(client.ServiceURL("os-security-group-default-rules", id), nil)
	return
}