package resolverrules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts contains the parameters of the Create function.
type CreateOpts struct {
	// Rule name.
	// It can contain 1 to 64 characters. Only letters, digits, underscores (_), hyphens (-), and periods (.) are allowed.
	Name string `json:"name" required:"true"`
	// Domain name.
	DomainName string `json:"domain_name" required:"true"`
	// ID of the endpoint to which the current rule belongs.
	EndpointID string `json:"endpoint_id" required:"true"`
	// Destination IP address added to a rule.
	IPAddresses []IPInfo `json:"ipaddresses" required:"true"`
}

// IPInfo is a destination IP address of the rule.
type IPInfo struct {
	// IP address.
	IP string `json:"ip,omitempty"`
}

// Create is used to create an endpoint rule.
// Send POST /v2.1/resolverrules
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*ResolverRule, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(v21URL(client, "resolverrules"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res ResolverRule
	err = extract.IntoStructPtr(raw.Body, &res, "resolver_rule")
	return &res, err
}
