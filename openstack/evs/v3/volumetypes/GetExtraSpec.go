package volumetypes

import "github.com/opentelekomcloud/gophertelekomcloud"

// GetExtraSpec requests an extra-spec specified by key for the given volume type ID
func GetExtraSpec(client *golangsdk.ServiceClient, volumeTypeID string, key string) (r GetExtraSpecResult) {
	resp, err := client.Get(client.ServiceURL("types", volumeTypeID, "extra_specs", key), &r.Body, nil)
	_, r.Header, r.Err = golangsdk.ParseResponse(resp, err)
	return
}
