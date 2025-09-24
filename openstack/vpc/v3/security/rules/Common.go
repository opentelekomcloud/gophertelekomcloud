package rules

// COMMON RESPONSE STRUCTS

type SecurityGroupRuleResponse struct {
	// Request ID.
	RequestID string `json:"request_id"`
	// Response body for creating a security group.
	SecurityGroupRule SecurityGroupRule `json:"security_group_rule"`
}

type SecurityGroupRule struct {
	// Security group rule ID, which uniquely identifies the security group rule.
	// The value is in UUID format with hyphens (-).
	ID string `json:"id"`
	// Provides supplementary information about the security group rule.
	// The value can contain no more than 255 characters and cannot contain angle brackets (< or >).
	Description string `json:"description"`
	// ID of the security group to which the security group rule belongs.
	SecurityGroupID string `json:"security_group_id"`
	// Inbound or outbound direction of a security group rule.
	// The value can be: ingress or egress.
	Direction string `json:"direction"`
	// Protocol type.
	// The value can be icmp, tcp, udp, icmpv6, or an IP number.
	// If protocol is left blank, all protocols are supported.
	Protocol string `json:"protocol"`
	// IP version.
	// The value can be IPv4 or IPv6.
	// If you do not set this parameter, IPv4 is used by default.
	Ethertype string `json:"ethertype"`
	// Port or port range. Can be a single port (80), a port range (1-30), or non-consecutive ports separated by commas (22,3389,80).
	Multiport string `json:"multiport"`
	// Action of the security group rule.
	// The value can be: allow or deny.
	// The default value is deny.
	Action string `json:"action"`
	// Rule priority.
	// The value is from 1 to 100. The value 1 indicates the highest priority.
	Priority int `json:"priority"`
	// Time when the security group rule is created.
	// UTC time in the format of yyyy-MM-ddTHH:mm:ssZ.
	CreatedAt string `json:"created_at"`
	// Time when the security group rule is updated.
	// UTC time in the format of yyyy-MM-ddTHH:mm:ssZ.
	UpdatedAt string `json:"updated_at"`
	// ID of the project to which the security group rule belongs.
	ProjectID string `json:"project_id"`
	// ID of the remote security group, which allows or denies traffic to and from the security group.
	// Value range: ID of an existing security group.
	// Mutually exclusive with remote_ip_prefix and remote_address_group_id.
	RemoteGroupID string `json:"remote_group_id"`
	// Remote IP address or CIDR block.
	// If direction is egress: source IP; if ingress: destination IP.
	// Mutually exclusive with remote_group_id and remote_address_group_id.
	RemoteIPPrefix string `json:"remote_ip_prefix"`
	// ID of the remote IP address group.
	// Mutually exclusive with remote_ip_prefix and remote_group_id.
	RemoteAddressGroupID string `json:"remote_address_group_id"`
}
