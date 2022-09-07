package secgroups

import "github.com/opentelekomcloud/gophertelekomcloud"

// RemoveServer will disassociate a server from a security group.
func RemoveServer(client *golangsdk.ServiceClient, serverID, groupName string) (r RemoveServerResult) {
	raw, err := client.Post(serverActionURL(client, serverID), actionMap("remove", groupName), nil, nil)
	return
}
