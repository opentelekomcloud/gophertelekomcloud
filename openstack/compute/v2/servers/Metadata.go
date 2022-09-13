package servers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// Metadata requests all the metadata for the given server ID.
func Metadata(client *golangsdk.ServiceClient, id string) (map[string]string, error) {
	raw, err := client.Get(client.ServiceURL("servers", id, "metadata"), nil, nil)
	return extraMet(err, raw)
}
