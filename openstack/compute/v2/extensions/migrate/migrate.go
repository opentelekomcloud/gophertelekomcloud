package migrate

import "github.com/opentelekomcloud/gophertelekomcloud"

// Migrate will initiate a migration of the instance to another host.
func Migrate(client *golangsdk.ServiceClient, id string) (r MigrateResult) {
	raw, err := client.Post(actionURL(client, id), map[string]interface{}{"migrate": nil}, nil, nil)
	return
}
