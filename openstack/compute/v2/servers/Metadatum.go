package servers

import "github.com/opentelekomcloud/gophertelekomcloud"

// Metadatum requests the key-value pair with the given key for the given server ID.
func Metadatum(client *golangsdk.ServiceClient, id, key string) (map[string]string, error) {
	raw, err := client.Get(client.ServiceURL("servers", id, "metadata", key), nil, nil)
	return extraTum(err, raw)
}
