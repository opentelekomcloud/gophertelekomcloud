package grants

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

type CreateOptsBuilder interface {
	ToGrantCreateMap() (map[string]interface{}, error)
}

type CreateOpts struct {
	// 36-byte ID of a CMK that matches the
	// regular expression ^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$
	KeyID string `json:"key_id" required:"true"`
	// Indicates the ID of the authorized user.
	// The value is between 1 to 64 bytes and meets the regular expression "^[a-zA-Z0-9]{1,64}$".
	GranteePrincipal string `json:"grantee_principal" required:"true"`
	// Permissions that can be granted
	Operations []string `json:"operations" required:"true"`
	// Name of a grant which can be 1 to 255 characters in
	// length and matches the regular expression ^[a-zA-Z0-9:/_-]{1,255}$
	Name string `json:"name,omitempty"`
	// Indicates the ID of the retiring user. The value is between 1 to 64
	// bytes and meets the regular expression "^[a-zA-Z0-9]{1,64}$".
	RetiringPrincipal string `json:"retiring_principal,omitempty"`
	// Sequence represents 36-byte serial number of a request message
	Sequence string `json:"sequence,omitempty"`
}

// ToGrantCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToGrantCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create will create a new Grant based on the values in CreateOpts. To
// extract the Grant object from the response, call the Extract method on the
// CreateResult.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToGrantCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type DeleteOptsBuilder interface {
	ToGrantDeleteMap() (map[string]interface{}, error)
}

type DeleteOpts struct {
	// 36-byte ID of a CMK that matches the regular
	// expression ^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$
	KeyID string `json:"key_id" required:"true"`
	// 64-byte ID of a grant that meets the regular
	// expression ^[A-Fa-f0-9]{64}$
	GrantID string `json:"grant_id" required:"true"`
	// Sequence represents 36-byte serial number of a request message
	Sequence string `json:"sequence,omitempty"`
}

// ToGrantDeleteMap assembles a request body based on the contents of a
// DeleteOpts.
func (opts DeleteOpts) ToGrantDeleteMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Delete will delete the existing Grant based on the values in DeleteOpts. To
// extract result call the ExtractErr method on the DeleteResult.
func Delete(client *golangsdk.ServiceClient, opts DeleteOptsBuilder) (r DeleteResult) {
	b, err := opts.ToGrantDeleteMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(deleteURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type ListOptsBuilder interface {
	ToGrantListMap() (map[string]interface{}, error)
}

type ListOpts struct {
	KeyID    string `json:"key_id,omitempty"`
	Limit    string `json:"limit,omitempty"`
	Marker   string `json:"marker,omitempty"`
	Sequence string `json:"sequence,omitempty"`
}

func (opts ListOpts) ToGrantListMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// List will return a collection of Grants on a CMK.
func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) (r ListResult) {
	b, err := opts.ToGrantListMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(listURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
