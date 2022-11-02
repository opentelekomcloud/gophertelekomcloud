package volumetypes

import "github.com/opentelekomcloud/gophertelekomcloud"

// UpdateExtraSpecOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateExtraSpecOptsBuilder interface {
	ToVolumeTypeExtraSpecUpdateMap() (map[string]string, string, error)
}

// ToVolumeTypeExtraSpecUpdateMap assembles a body for an Update request based on
// the contents of a ExtraSpecOpts.
func (opts ExtraSpecsOpts) ToVolumeTypeExtraSpecUpdateMap() (map[string]string, string, error) {
	if len(opts) != 1 {
		err := golangsdk.ErrInvalidInput{}
		err.Argument = "volumetypes.ExtraSpecOpts"
		err.Info = "Must have one and only one key-value pair"
		return nil, "", err
	}

	var key string
	for k := range opts {
		key = k
	}

	return opts, key, nil
}

// UpdateExtraSpec will updates the value of the specified volume type's extra spec
// for the key in opts.
func UpdateExtraSpec(client *golangsdk.ServiceClient, volumeTypeID string, opts UpdateExtraSpecOptsBuilder) (r UpdateExtraSpecResult) {
	b, key, err := opts.ToVolumeTypeExtraSpecUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Put(client.ServiceURL("types", volumeTypeID, "extra_specs", key), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	_, r.Header, r.Err = golangsdk.ParseResponse(resp, err)
	return
}
