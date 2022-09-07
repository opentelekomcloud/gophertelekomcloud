package secgroups

import "github.com/opentelekomcloud/gophertelekomcloud"

// AddServer will associate a server and a security group, enforcing the
// rules of the group on the server.
func AddServer(client *golangsdk.ServiceClient, serverID, groupName string) (r AddServerResult) {
	raw, err := client.Post(client.ServiceURL("servers", serverID, "action"), actionMap("add", groupName), nil, nil)
	return
}
