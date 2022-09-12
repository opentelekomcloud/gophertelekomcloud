package secgroups

import "github.com/opentelekomcloud/gophertelekomcloud"

// DeleteRule will permanently delete a rule from a security group.
func DeleteRule(client *golangsdk.ServiceClient, id string) (err error) {
	_, err = client.Delete(client.ServiceURL("os-security-group-rules", id), nil)
	return
}
