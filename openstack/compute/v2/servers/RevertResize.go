package servers

import "github.com/opentelekomcloud/gophertelekomcloud"

// RevertResize cancels a previous resize operation on a server. See Resize() for more details.
func RevertResize(client *golangsdk.ServiceClient, id string) (err error) {
	_, err = client.Post(client.ServiceURL("servers", id, "action"), map[string]interface{}{"revertResize": nil}, nil, nil)
	return
}
