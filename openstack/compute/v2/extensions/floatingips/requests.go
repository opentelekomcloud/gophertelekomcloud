package floatingips

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// List returns a Pager that allows you to iterate over a collection of FloatingIPs.
func List(client *golangsdk.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, listURL(client), func(r pagination.PageResult) pagination.Page {
		return FloatingIPPage{pagination.SinglePageBase(r)}
	})
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToFloatingIPCreateMap() (map[string]interface{}, error)
}

// CreateOpts specifies a Floating IP allocation request.
type CreateOpts struct {
	// Pool is the pool of Floating IPs to allocate one from.
	Pool string `json:"pool" required:"true"`
}

// ToFloatingIPCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToFloatingIPCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// AssociateOptsBuilder allows extensions to add additional parameters to the
// Associate request.
type AssociateOptsBuilder interface {
	ToFloatingIPAssociateMap() (map[string]interface{}, error)
}

// AssociateOpts specifies the required information to associate a Floating IP with an instance
type AssociateOpts struct {
	// FloatingIP is the Floating IP to associate with an instance.
	FloatingIP string `json:"address" required:"true"`

	// FixedIP is an optional fixed IP address of the server.
	FixedIP string `json:"fixed_address,omitempty"`
}

// ToFloatingIPAssociateMap constructs a request body from AssociateOpts.
func (opts AssociateOpts) ToFloatingIPAssociateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "addFloatingIp")
}

// DisassociateOptsBuilder allows extensions to add additional parameters to
// the Disassociate request.
type DisassociateOptsBuilder interface {
	ToFloatingIPDisassociateMap() (map[string]interface{}, error)
}

// DisassociateOpts specifies the required information to disassociate a
// Floating IP with a server.
type DisassociateOpts struct {
	FloatingIP string `json:"address" required:"true"`
}

// ToFloatingIPDisassociateMap constructs a request body from DisassociateOpts.
func (opts DisassociateOpts) ToFloatingIPDisassociateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "removeFloatingIp")
}
