package dnatrules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// CreateOptsBuilder is an interface must satisfy to be used as Create
// options.
type CreateOptsBuilder interface {
	ToDnatRuleCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains all the values needed to create a new dnat rule
// resource.
type CreateOpts struct {
	NatGatewayID        string `json:"nat_gateway_id" required:"true"`
	PortID              string `json:"port_id,omitempty"`
	PrivateIp           string `json:"private_ip,omitempty"`
	InternalServicePort *int   `json:"internal_service_port" required:"true"`
	FloatingIpID        string `json:"floating_ip_id" required:"true"`
	ExternalServicePort *int   `json:"external_service_port" required:"true"`
	Protocol            string `json:"protocol" required:"true"`
}

// ToDnatRuleCreateMap allows CreateOpts to satisfy the CreateOptsBuilder
// interface
func (opts CreateOpts) ToDnatRuleCreateMap() (map[string]interface{}, error) {
	return build.RequestBodyMap(opts, "dnat_rule")
}

// Create is a method by which can create a new dnat rule
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToDnatRuleCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(rootURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	return
}

// Get is a method by which can get the detailed information of the specified
// dnat rule.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(resourceURL(client, id), &r.Body, nil)
	return
}

// Delete is a method by which can be able to delete a dnat rule
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(resourceURL(client, id), nil)
	return
}
