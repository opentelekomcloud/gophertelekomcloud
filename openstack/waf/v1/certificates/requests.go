package certificates

import (
	"fmt"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToCertCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains all the values needed to create a new certificate.
type CreateOpts struct {
	// Certificate name
	Name string `json:"name" required:"true"`
	// Certificate content
	Content string `json:"content" required:"true"`
	// Private Key
	Key string `json:"key" required:"true"`
}

// ToCertCreateMap builds a create request body from CreateOpts.
func (opts CreateOpts) ToCertCreateMap() (map[string]interface{}, error) {
	return build.RequestBodyMap(opts, "")
}

// Create will create a new certificate based on the values in CreateOpts.
func Create(c *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToCertCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{200}}
	_, r.Err = c.Post(rootURL(c), b, &r.Body, reqOpt)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToCertUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contains all the values needed to update a certificate.
type UpdateOpts struct {
	// Certificate name
	Name string `json:"name,omitempty"`
}

// ToCertUpdateMap builds a update request body from UpdateOpts.
func (opts UpdateOpts) ToCertUpdateMap() (map[string]interface{}, error) {
	return build.RequestBodyMap(opts, "")
}

// Update accepts a UpdateOpts struct and uses the values to update a certificate.The response code from api is 200
func Update(c *golangsdk.ServiceClient, certID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToCertUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{200}}
	_, r.Err = c.Put(resourceURL(c, certID), b, &r.Body, reqOpt)
	return
}

// Get retrieves a particular certificate based on its unique ID.
func Get(c *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(resourceURL(c, id), &r.Body, openstack.StdRequestOpts())
	return
}

// Delete will permanently delete a particular certificate based on its unique ID.
func Delete(c *golangsdk.ServiceClient, id string) (r DeleteResult) {
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{204},
		MoreHeaders: openstack.StdRequestOpts().MoreHeaders}
	_, r.Err = c.Delete(resourceURL(c, id), reqOpt)
	return
}

type ListOptsBuilder interface {
	ToCertificateListQuery() (string, error)
}

type ListOpts struct {
	Offset int `q:"offset"`
	Limit  int `q:"limit"`
}

func (opts ListOpts) ToCertificateListQuery() (string, error) {
	if opts.Offset > 0xffff || opts.Offset < 0 {
		return "", fmt.Errorf("offset must be 0-65535")
	}
	if opts.Limit > 50 || opts.Offset < -1 {
		return "", fmt.Errorf("limit must be -1-50")
	}
	q, err := build.QueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func List(c *golangsdk.ServiceClient, opts ListOptsBuilder) (p pagination.Pager) {
	url := rootURL(c)
	if opts != nil {
		q, err := opts.ToCertificateListQuery()
		if err != nil {
			p.Err = err
			return
		}
		url += q
	}
	p = pagination.Pager{
		Client:     c,
		InitialURL: url,
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return CertificatePage{OffsetPageBase: pagination.OffsetPageBase{PageResult: r}}
		},
	}
	p.Headers = map[string]string{"content-type": "application/json"}
	return p
}
