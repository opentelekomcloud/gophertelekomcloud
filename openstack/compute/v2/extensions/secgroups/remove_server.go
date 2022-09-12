package secgroups

import "github.com/opentelekomcloud/gophertelekomcloud"

// RemoveServer will disassociate a server from a security group.
func RemoveServer(client *golangsdk.ServiceClient, serverID, groupName string) (err error) {
	_, err = client.Post(client.ServiceURL("servers", serverID, "action"), actionMap("remove", groupName), nil, nil)
	return
}
