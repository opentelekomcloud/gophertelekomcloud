package eipstags

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// CreateOptsBuilder
type CreateOptsBuilder interface {
	ToPublicIpTagCreateMap() (map[string]interface{}, error)
}

// CreateOpts
type CreateOpts struct {
	Tag map[string]string `json:"tag" required:"true"`
}

func (opts CreateOpts) ToPublicIpTagCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder, id string) (r CreateResult) {
	b, err := opts.ToPublicIpTagCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(rootURL(client, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Get
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(rootURL(client, id), &r.Body, nil)
	return
}
