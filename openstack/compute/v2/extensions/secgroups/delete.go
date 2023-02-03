package secgroups

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a security group from the project.
func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	_, err = client.Delete(client.ServiceURL("os-security-groups", id), nil)
	return
}
