package flavors

import "github.com/opentelekomcloud/gophertelekomcloud"

// DeleteExtraSpec will delete the key-value pair with the given key for the given flavor ID.
func DeleteExtraSpec(client *golangsdk.ServiceClient, flavorID, key string) (err error) {
	_, err = client.Delete(client.ServiceURL("flavors", flavorID, "os-extra_specs", key), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
