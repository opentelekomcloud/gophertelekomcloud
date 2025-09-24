package rules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Whether to only send the check request.
	// The value can be:
	// - true: The request will be checked and the security group rule will not be created.
	//         Check items include mandatory parameters, request format, and constraints.
	//         If the check fails, the system returns an error.
	//         If the check succeeds, response code 202 will be returned.
	// - false (default): A request will be sent and a security group rule will be created.
	DryRun bool `json:"dry_run,omitempty"`
	// Request body for creating a security group rule.
	SecurityGroupRule SecurityGroupRuleOptions `json:"security_group_rule" required:"true"`
}

type SecurityGroupRuleOptions struct {
	// ID of the security group to which the security group rule belongs.
	SecurityGroupID string `json:"security_group_id" required:"true"`
	// Provides supplementary information about the security group rule.
	// The value can contain no more than 255 characters and cannot contain angle brackets (< or >).
	Description string `json:"description,omitempty"`
	// Inbound or outbound direction of a security group rule.
	// The value can be: ingress (inbound) or egress (outbound).
	Direction string `json:"direction" required:"true"`
	// IP version.
	// The value can be IPv4 or IPv6.
	// If you do not set this parameter, IPv4 is used by default.
	Ethertype string `json:"ethertype,omitempty"`
	// Protocol type.
	// The value can be icmp, tcp, udp, icmpv6, or an IP number.
	// If protocol is left blank, all protocols are supported.
	Protocol string `json:"protocol,omitempty"`
	// Port or port range. Can be a single port (80), a port range (1-30), or non-consecutive ports separated by commas (22,3389,80).
	Multiport string `json:"multiport,omitempty"`
	// Remote IP address or CIDR block.
	// If direction is egress: source IP; if ingress: destination IP.
	// Mutually exclusive with remote_group_id and remote_address_group_id.
	RemoteIPPrefix string `json:"remote_ip_prefix,omitempty"`
	// ID of the remote security group, which allows or denies traffic to and from the security group.
	// Value range: ID of an existing security group.
	// Mutually exclusive with remote_ip_prefix and remote_address_group_id.
	RemoteGroupID string `json:"remote_group_id,omitempty"`
	// Action of the security group rule.
	// The value can be: allow or deny.
	// The default value is deny.
	Action string `json:"action,omitempty"`
	// Rule priority.
	// The value is from 1 to 100. The value 1 indicates the highest priority.
	Priority int `json:"priority,omitempty"`
}

// This function is used to create a security group rule.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*SecurityGroupRule, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/vpc/security-group-rules
	raw, err := client.Post(client.ServiceURL("security-group-rules"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res SecurityGroupRuleResponse
	return &res.SecurityGroupRule, extract.Into(raw.Body, &res)
}
