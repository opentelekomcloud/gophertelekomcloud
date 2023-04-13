package certificates

import "github.com/opentelekomcloud/gophertelekomcloud"

// UpdateOptsBuilder is the interface options structs have to satisfy in order
// to be used in the main Update operation in this package. Since many
// extensions decorate or modify the common logic, it is useful for them to
// satisfy a basic interface in order for them to be used.
type UpdateOptsBuilder interface {
	ToCertificateUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts is the common options' struct used in this package's Update
// operation.
type UpdateOpts struct {
	Name        string  `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Domain      string  `json:"domain,omitempty"`
	PrivateKey  string  `json:"private_key,omitempty"`
	Certificate string  `json:"certificate,omitempty"`
}

// ToCertificateUpdateMap casts a UpdateOpts struct to a map.
func (opts UpdateOpts) ToCertificateUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "certificate")
}

// Update is an operation which modifies the attributes of the specified Certificate.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (r UpdateResult) {
	b, err := opts.ToCertificateUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(client.ServiceURL("certificates", id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
