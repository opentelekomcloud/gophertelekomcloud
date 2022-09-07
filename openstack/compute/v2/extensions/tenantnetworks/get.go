package tenantnetworks

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// Get returns data about a previously created Network.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	raw, err := client.Get(client.ServiceURL("os-tenant-networks", id), nil, nil)
	return
}
