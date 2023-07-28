package domains

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToDomainCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains all the values needed to create a new backup.
type CreateOpts struct {
	// Domain name
	HostName string `json:"hostname" required:"true"`
	// Certificate ID
	CertificateId string `json:"certificate_id,omitempty"`
	// The original server information
	Server []ServerOpts `json:"server" required:"true"`
	// Whether proxy is configured
	Proxy *bool `json:"proxy" required:"true"`
	// TLS version
	TLS string `json:"tls,omitempty"`
	// Cipher suite version
	Cipher string `json:"cipher,omitempty"`
	// The type of the source IP header
	SipHeaderName string `json:"sip_header_name,omitempty"`
	// The HTTP request header for identifying the real source IP.
	SipHeaderList []string `json:"sip_header_list,omitempty"`
}

type ServerOpts struct {
	// Protocol type of the client
	ClientProtocol string `json:"client_protocol" required:"true"`
	// Protocol used by WAF to forward client requests to the server
	ServerProtocol string `json:"server_protocol" required:"true"`
	// IP address or domain name of the web server that the client accesses.
	Address string `json:"address" required:"true"`
	// Port number used by the web server
	Port int `json:"port" required:"true"`
}

// ToDomainCreateMap builds a create request body from CreateOpts.
func (opts CreateOpts) ToDomainCreateMap() (map[string]interface{}, error) {
	return build.RequestBodyMap(opts, "")
}

// Create will create a new Domain based on the values in CreateOpts.
func Create(c *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToDomainCreateMap()
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
	ToDomainUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contains all the values needed to update a Domain.
type UpdateOpts struct {
	// Certificate ID
	CertificateId string `json:"certificate_id,omitempty"`
	// The original server information
	Server []ServerOpts `json:"server,omitempty"`
	// Whether proxy is configured
	Proxy *bool `json:"proxy,omitempty"`
	// TLS version
	TLS string `json:"tls,omitempty"`
	// Cipher suite version
	Cipher string `json:"cipher,omitempty"`
	// The type of the source IP header
	SipHeaderName string `json:"sip_header_name,omitempty"`
	// The HTTP request header for identifying the real source IP.
	SipHeaderList []string `json:"sip_header_list,omitempty"`
	// Alarm page configuration
	BlockPage *BlockPage `json:"block_page,omitempty"`
}

type BlockPage struct {
	Template    string      `json:"template" required:"true"`
	CustomPage  *CustomPage `json:"custom_page,omitempty"`
	RedirectUrl string      `json:"redirect_url,omitempty"`
}

type CustomPage struct {
	StatusCode  string `json:"status_code" required:"true"`
	ContentType string `json:"content_type" required:"true"`
	Content     string `json:"content" required:"true"`
}

// ToDomainUpdateMap builds a update request body from UpdateOpts.
func (opts UpdateOpts) ToDomainUpdateMap() (map[string]interface{}, error) {
	return build.RequestBodyMap(opts, "")
}

// Update accepts a UpdateOpts struct and uses the values to update a Domain.The response code from api is 200
func Update(c *golangsdk.ServiceClient, domainID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToDomainUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(resourceURL(c, domainID), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Get retrieves a particular Domain based on its unique ID.
func Get(c *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(resourceURL(c, id), &r.Body, openstack.StdRequestOpts())
	return
}

// Delete will permanently delete a particular Domain based on its unique ID.
func Delete(c *golangsdk.ServiceClient, id string) (r DeleteResult) {
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{204},
		MoreHeaders: openstack.StdRequestOpts().MoreHeaders}
	_, r.Err = c.Delete(resourceURL(c, id), reqOpt)
	return
}
