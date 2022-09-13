package servers

import "github.com/opentelekomcloud/gophertelekomcloud"

// DeleteMetadatum will delete the key-value pair with the given key for the given server ID.
func DeleteMetadatum(client *golangsdk.ServiceClient, id, key string) (err error) {
	_, err = client.Delete(client.ServiceURL("servers", id, "metadata", key), nil)
	return
}
