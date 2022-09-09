package migrate

import "github.com/opentelekomcloud/gophertelekomcloud"

// Migrate will initiate a migration of the instance to another host.
func Migrate(client *golangsdk.ServiceClient, id string) (err error) {
	_, err = client.Post(client.ServiceURL("servers", id, "action"), map[string]interface{}{"migrate": nil}, nil, nil)
	return
}
