package tags

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// CreateOpts implements CreateOptsBuilder
type CreateOpts struct {
	// Tags is a set of tags.
	Tags map[string]string `json:"tags" required:"true"`
}

// Create implements create image request
func Create(client *golangsdk.ServiceClient, resource_type, resource_id string, opts CreateOpts) (r CreateResult) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		r.Err = err
		return r
	}

	_, r.Err = client.Put(client.ServiceURL("os-vendor-tags", resource_type, resource_id), b, &r.Body, &golangsdk.RequestOpts{OkCodes: []int{200}})
	return
}

// Get implements tags get request
func Get(client *golangsdk.ServiceClient, resource_type, resource_id string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("os-vendor-tags", resource_type, resource_id), &r.Body, nil)
	return
}

// Delete implements image delete request by creating empty tag map
func Delete(client *golangsdk.ServiceClient, resource_type, resource_id string) (err error) {
	_, err = Create(client, resource_type, resource_id, CreateOpts{
		Tags: map[string]string{},
	}).Extract()
	return
}
