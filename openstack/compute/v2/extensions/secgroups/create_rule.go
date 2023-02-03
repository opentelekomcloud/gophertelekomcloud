package secgroups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateRuleOpts represents the configuration for adding a new rule to an
// existing security group.
type CreateRuleOpts struct {
	// ID is the ID of the group that this rule will be added to.
	ParentGroupID string `json:"parent_group_id" required:"true"`

	// FromPort is the lower bound of the port range that will be opened.
	// Use -1 to allow all ICMP traffic.
	FromPort int `json:"from_port"`

	// ToPort is the upper bound of the port range that will be opened.
	// Use -1 to allow all ICMP traffic.
	ToPort int `json:"to_port"`

	// IPProtocol the protocol type that will be allowed, e.g. TCP.
	IPProtocol string `json:"ip_protocol" required:"true"`

	// CIDR is the network CIDR to allow traffic from.
	// This is ONLY required if FromGroupID is blank. This represents the IP
	// range that will be the source of network traffic to your security group.
	// Use 0.0.0.0/0 to allow all IP addresses.
	CIDR string `json:"cidr,omitempty" or:"FromGroupID"`

	// FromGroupID represents another security group to allow access.
	// This is ONLY required if CIDR is blank. This value represents the ID of a
	// group that forwards traffic to the parent group. So, instead of accepting
	// network traffic from an entire IP range, you can instead refine the
	// inbound source by an existing security group.
	FromGroupID string `json:"group_id,omitempty" or:"CIDR"`
}

// CreateRule will add a new rule to an existing security group (whose ID is
// specified in CreateRuleOpts). You have the option of controlling inbound
// traffic from either an IP range (CIDR) or from another security group.
func CreateRule(client *golangsdk.ServiceClient, opts CreateRuleOpts) (*Rule, error) {
	b, err := build.RequestBody(opts, "security_group_rule")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("os-security-group-rules"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Rule
	err = extract.IntoStructPtr(raw.Body, &res, "security_group_rule")
	return &res, err
}
