package tags

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// CreateOptsBuilder describes struct types that can be accepted by the Create call.
// The CreateOpts struct in this package does.
type CreateOptsBuilder interface {
	// Returns value that can be passed to json.Marshal
	ToTagsCreateMap() (map[string]interface{}, error)
}

// CreateOpts implements CreateOptsBuilder
type CreateOpts struct {
	// Tags is a set of tags.
	Tags []string `json:"tags" required:"true"`
}

// ToImageCreateMap assembles a request body based on the contents of
// a CreateOpts.
func (opts CreateOpts) ToTagsCreateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Create implements create tags request
func Create(client *golangsdk.ServiceClient, server_id string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToTagsCreateMap()
	if err != nil {
		r.Err = err
		return r
	}
	_, r.Err = client.Put(createURL(client, server_id), b, &r.Body, &golangsdk.RequestOpts{OkCodes: []int{200}})
	return
}

// Get implements tags get request
func Get(client *golangsdk.ServiceClient, server_id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, server_id), &r.Body, nil)
	return
}

// Delete implements image delete request
func Delete(client *golangsdk.ServiceClient, server_id string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, server_id), nil)
	return
}
