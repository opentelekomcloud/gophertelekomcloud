package volumetypes

import "github.com/opentelekomcloud/gophertelekomcloud"

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToVolumeTypeUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contain options for updating an existing Volume Type. This object is passed
// to the volumetypes.Update function. For more information about the parameters, see
// the Volume Type object.
type UpdateOpts struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	IsPublic    *bool   `json:"is_public,omitempty"`
}

// ToVolumeTypeUpdateMap assembles a request body based on the contents of an
// UpdateOpts.
func (opts UpdateOpts) ToVolumeTypeUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "volume_type")
}

// Update will update the Volume Type with provided information. To extract the updated
// Volume Type from the response, call the Extract method on the UpdateResult.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToVolumeTypeUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Put(client.ServiceURL("types", id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	_, r.Header, r.Err = golangsdk.ParseResponse(resp, err)
	return
}
