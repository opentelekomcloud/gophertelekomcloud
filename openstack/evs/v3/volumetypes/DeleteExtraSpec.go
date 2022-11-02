package volumetypes

import "github.com/opentelekomcloud/gophertelekomcloud"

// DeleteExtraSpec will delete the key-value pair with the given key for the given
// volume type ID.
func DeleteExtraSpec(client *golangsdk.ServiceClient, volumeTypeID, key string) (r DeleteExtraSpecResult) {
	resp, err := client.Delete(client.ServiceURL("types", volumeTypeID, "extra_specs", key), &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	_, r.Header, r.Err = golangsdk.ParseResponse(resp, err)
	return
}
