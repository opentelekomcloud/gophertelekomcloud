package flavors

import "github.com/opentelekomcloud/gophertelekomcloud"

func GetExtraSpec(client *golangsdk.ServiceClient, flavorID string, key string) (r GetExtraSpecResult) {
	raw, err := client.Get(client.ServiceURL("flavors", flavorID, "os-extra_specs", key), nil, nil)
	return
}
