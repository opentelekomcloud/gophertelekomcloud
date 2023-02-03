package servers

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete requests that a server previously provisioned be removed from your account.
func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	_, err = client.Delete(client.ServiceURL("servers", id), nil)
	return
}
