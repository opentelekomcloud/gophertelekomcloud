package servers

import "github.com/opentelekomcloud/gophertelekomcloud"

// GetPassword makes a request against the nova API to get the encrypted
// administrative password.
func GetPassword(client *golangsdk.ServiceClient, serverId string) (r GetPasswordResult) {
	raw, err := client.Get(client.ServiceURL("servers", serverId, "os-server-password"), nil, nil)
	return
}
