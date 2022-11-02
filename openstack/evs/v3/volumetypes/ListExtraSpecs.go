package volumetypes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// ListExtraSpecs requests all the extra-specs for the given volume type ID.
func ListExtraSpecs(client *golangsdk.ServiceClient, volumeTypeID string) (r ListExtraSpecsResult) {
	resp, err := client.Get(client.ServiceURL("types", volumeTypeID, "extra_specs"), &r.Body, nil)
	_, r.Header, r.Err = golangsdk.ParseResponse(resp, err)
	return
}
