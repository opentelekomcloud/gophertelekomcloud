package volumetypes

import "github.com/opentelekomcloud/gophertelekomcloud"

// CreateExtraSpecsOptsBuilder allows extensions to add additional parameters to the
// CreateExtraSpecs requests.
type CreateExtraSpecsOptsBuilder interface {
	ToVolumeTypeExtraSpecsCreateMap() (map[string]interface{}, error)
}

// ExtraSpecsOpts is a map that contains key-value pairs.
type ExtraSpecsOpts map[string]string

// ToVolumeTypeExtraSpecsCreateMap assembles a body for a Create request based on
// the contents of ExtraSpecsOpts.
func (opts ExtraSpecsOpts) ToVolumeTypeExtraSpecsCreateMap() (map[string]interface{}, error) {
	return map[string]interface{}{"extra_specs": opts}, nil
}

// CreateExtraSpecs will create or update the extra-specs key-value pairs for
// the specified volume type.
func CreateExtraSpecs(client *golangsdk.ServiceClient, volumeTypeID string, opts CreateExtraSpecsOptsBuilder) (r CreateExtraSpecsResult) {
	b, err := opts.ToVolumeTypeExtraSpecsCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(client.ServiceURL("types", volumeTypeID, "extra_specs"), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	_, r.Header, r.Err = golangsdk.ParseResponse(resp, err)
	return
}
