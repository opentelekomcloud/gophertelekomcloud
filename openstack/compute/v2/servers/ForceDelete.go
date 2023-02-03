package servers

import "github.com/opentelekomcloud/gophertelekomcloud"

// ForceDelete forces the deletion of a server.
func ForceDelete(client *golangsdk.ServiceClient, id string) (err error) {
	_, err = client.Post(client.ServiceURL("servers", id, "action"), map[string]interface{}{"forceDelete": ""}, nil, nil)
	return
}
