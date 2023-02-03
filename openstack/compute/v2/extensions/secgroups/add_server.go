package secgroups

import "github.com/opentelekomcloud/gophertelekomcloud"

func actionMap(prefix, groupName string) map[string]map[string]string {
	return map[string]map[string]string{
		prefix + "SecurityGroup": {"name": groupName},
	}
}

// AddServer will associate a server and a security group, enforcing the rules of the group on the server.
func AddServer(client *golangsdk.ServiceClient, serverID, groupName string) (err error) {
	_, err = client.Post(client.ServiceURL("servers", serverID, "action"), actionMap("add", groupName), nil, nil)
	return
}
