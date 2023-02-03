package defsecrules

import (
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// CreateOpts represents the configuration for adding a new default rule.
type CreateOpts struct {
	// The lower bound of the port range that will be opened.
	FromPort int `json:"from_port"`
	// The upper bound of the port range that will be opened.
	ToPort int `json:"to_port"`
	// The protocol type that will be allowed, e.g. TCP.
	IPProtocol string `json:"ip_protocol" required:"true"`
	// ONLY required if FromGroupID is blank. This represents the IP range that
	// will be the source of network traffic to your security group.
	//
	// Use 0.0.0.0/0 to allow all IPv4 addresses.
	// Use ::/0 to allow all IPv6 addresses.
	CIDR string `json:"cidr,omitempty"`
}

// ToRuleCreateMap builds the create rule options into a serializable format.
func (opts CreateOpts) ToRuleCreateMap() (map[string]interface{}, error) {
	if opts.FromPort == 0 && strings.ToUpper(opts.IPProtocol) != "ICMP" {
		return nil, golangsdk.ErrMissingInput{Argument: "FromPort"}
	}
	if opts.ToPort == 0 && strings.ToUpper(opts.IPProtocol) != "ICMP" {
		return nil, golangsdk.ErrMissingInput{Argument: "ToPort"}
	}
	return golangsdk.BuildRequestBody(opts, "security_group_default_rule")
}

// Create is the operation responsible for creating a new default rule.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*DefaultRule, error) {
	b, err := opts.ToRuleCreateMap()
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("os-security-group-default-rules"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
