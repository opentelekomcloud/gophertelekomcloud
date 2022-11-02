package volumes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UpdateOpts contain options for updating an existing Volume. This object is passed
// to the volumes.Update function. For more information about the parameters, see
// the Volume object.
type UpdateOpts struct {
	// Specifies the disk name. The value can contain a maximum of 255 bytes.
	Name string `json:"name,omitempty"`
	// Specifies the disk description. The value can contain a maximum of 255 bytes.
	Description string `json:"description,omitempty"`
	// Specifies the disk metadata.
	// The length of the key or value in the metadata cannot exceed 255 bytes.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Specifies also the disk name. You can specify either parameter name or display_name.
	// If both parameters are specified, the name value is used. The value can contain a maximum of 255 bytes.
	DisplayName string `json:"display_name,omitempty"`
	// Specifies also the disk description. You can specify either parameter description or display_description.
	// If both parameters are specified, the description value is used. The value can contain a maximum of 255 bytes.
	DisplayDescription string `json:"display_description,omitempty"`
}

// Update will update the Volume with provided information. To extract the updated
// Volume from the response, call the Extract method on the UpdateResult.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (r UpdateResult) {
	b, err := build.RequestBody(opts, "volume")
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Put(client.ServiceURL("volumes", id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	_, r.Header, r.Err = golangsdk.ParseResponse(resp, err)
	return
}
