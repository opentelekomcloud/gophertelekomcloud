package acl

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateACLRuleOpts struct {
	// Protected object ID, which is used to distinguish between Internet border protection and VPC border protection
	// after a cloud firewall is created. You can obtain the ID by calling the Get function in management package.
	// In the return value, find the ID in ProtectObjects[n].ObjectID.
	// If the value of type is 0, the protected object ID belongs to the Internet border.
	// If the value of type is 1, the protected object ID belongs to the VPC border.
	ObjectID string `json:"object_id" required:"true"`
	// Rule type: 0 (Internet border rule), 1 (inter-VPC rule), or 2 (NAT rule).
	// When type is set to 0, the source and destination addresses of the rule
	// must be EIPs or domain names of the public network.
	// For an inter-VPC rule, the source and destination addresses must be private IP addresses.
	// For a NAT rule, the source address must be a private IP address, and the destination address
	// must be an EIP or domain name of the public network.
	Type int `json:"type" required:"true"`
	// Rules in a rule addition request.
	Rules []Rule `json:"rules" required:"true"`
}

type Rule struct {
	// Rule name.
	Name string `json:"name" required:"true"`
	// Request body for changing the rule sequence.
	Sequence OrderRuleAclDto `json:"sequence" required:"true"`
	// Internet protocol type of an address (0: IPv4, 1: IPv6).
	AddressType int `json:"address_type" required:"true"`
	// Rule action (0: Permit, 1: Deny).
	ActionType int `json:"action_type" required:"true"`
	// Rule status (0: Disabled, 1: Enabled)
	Status int `json:"status" required:"true"`
	// Rule application list. Rule application type:
	// HTTP, HTTPS, TLS1, DNS, SSH, MYSQL, SMTP, RDP, RDPS, VNC, POP3, IMAP4, SMTPS, POP3S, FTPS, ANY, or BGP.
	Applications []string `json:"applications,omitempty"`
	// JSON string converted from the applications field in the application list.
	ApplicationsJsonString string `json:"applicationsJsonString,omitempty"`
	// Persistent connection duration.
	LongConnectTime int64 `json:"long_connect_time,omitempty"`
	// Persistent connection duration (hour).
	LongConnectTimeHour int64 `json:"long_connect_time_hour,omitempty"`
	// Persistent connection duration (minutes).
	LongConnectTimeMinute int64 `json:"long_connect_time_minute,omitempty"`
	// Persistent connection duration (seconds).
	LongConnectTimeSecond int64 `json:"long_connect_time_second,omitempty"`
	// Whether to enable long connection (0: No, 1: Yes).
	LongConnectEnable int `json:"long_connect_enable" required:"true"`
	// Description of the rule.
	Description string `json:"description,omitempty"`
	// Direction of rule (0: Inbound, 1: Outbound). This parameter is mandatory when type is set to 0 (Internet rule) or 2 (NAT rule).
	Direction int `json:"direction,omitempty"`
	// Source address Data Transport Object.
	Source RuleAddressDtoRequest `json:"source" required:"true"`
	// Destination address Data Transport Object.
	Destination RuleAddressDtoRequest `json:"destination" required:"true"`
	// Service object associated with the rule.
	Service RuleServiceDto `json:"service" required:"true"`
	// Tag object attached to a rule.
	Tag TagsVO `json:"tag,omitempty"`
}

// OrderRuleAclDto represents the ordering of rule actions.
type OrderRuleAclDto struct {
	// ID of the target rule. The added rule is placed after this rule.
	// This parameter cannot be left blank when the added rule is not pinned on top,
	// and can be left blank when the added rule is pinned on top.
	DestRuleId string `json:"dest_rule_id,omitempty"`
	// Whether to pin on top: 0 (no), 1 (yes).
	Top int `json:"top,omitempty"`
	// Whether to pin the rule to the bottom (0: No, 1: Yes).
	Bottom int `json:"bottom,omitempty"`
}

type RuleAddressDtoRequest struct {
	// Address type: 0 (manual input), 1 (associated IP address group), 2 (domain name),
	// 3 (geographical location), 4 (domain name group) 5 (multiple objects),
	// 6 (domain name group - network), 7 (domain name group - application).
	Type int `json:"type" required:"true"`
	// Internet protocol type of an address (0: IPv4, 1: IPv6).
	// If type is 0, this parameter cannot be left blank.
	AddressType int `json:"address_type,omitempty"`
	// IP address information. It cannot be left blank if type is set to 0.
	Address string `json:"address,omitempty"`
	// ID of an associated IP address group. This parameter cannot be left blank when type is set to 1.
	AddressSetID string `json:"address_set_id,omitempty"`
	// Name of an associated IP address.
	// Name of an associated IP address group. This parameter cannot be left blank when type is set to 1.
	AddressSetName string `json:"address_set_name,omitempty"`
	// Name of a domain name address. This parameter is valid when type is set to 2 (domain name) or 7 (application domain name group).
	DomainAddressName string `json:"domain_address_name,omitempty"`
	// JSON value of the rule region list.
	RegionListJson string `json:"region_list_json,omitempty"`
	// Rule region list.
	RegionList []IpRegionDto `json:"region_list,omitempty"`
	// Domain group ID. The value cannot be left blank when type is set to 4 (domain name group) or 7 (domain name group - application).
	DomainSetID string `json:"domain_set_id,omitempty"`
	// Domain group name. The value cannot be left blank when type is set to 4 (domain name group) or 7 (domain name group - application).
	DomainSetName string `json:"domain_set_name,omitempty"`
	// IP address list. This parameter cannot be left blank when type is set to 5 (multiple objects).
	IPAddresses []string `json:"ip_address,omitempty"`
	// Address group type. It cannot be left blank when type is set to 1 (associated IP address group).
	// It value can be 0 (user-defined address group), 1 (WAF back-to-source IP address group),
	// 2 (DDoS back-to-source IP address group), or 3 (NAT64 address group).
	AddressSetType int `json:"address_set_type,omitempty"`
	// Pre-defined address group ID list. This parameter cannot be left blank when type is set to 5 (multiple objects).
	PredefinedGroup []string `json:"predefined_group,omitempty"`
	// Address group ID list. This parameter cannot be left blank when type is set to 5 (multiple objects).
	AddressGroup []string `json:"address_group,omitempty"`
}

type IpRegionDto struct {
	// Region ID.
	RegionID string `json:"region_id,omitempty"`
	// Region type: 0 (country), 1 (province), and 2 (continent).
	RegionType int `json:"region_type,omitempty"`
}

type RuleServiceDto struct {
	// Service input type (0: manual, 1: automatic).
	Type int `json:"type" required:"true"`
	// Protocol type (6: TCP, 17: UDP, 1: ICMP, 58: ICMPv6, -1: any).
	// It cannot be left blank when type is set to 0 (manual).
	Protocol int `json:"protocol,omitempty"`
	// List of protocols (6: TCP, 17: UDP, 1: ICMP, 58: ICMPv6, -1: any).
	// It cannot be left blank when type is set to 0 (manual).
	Protocols []int `json:"protocols,omitempty"`
	// Source port.
	SourcePort string `json:"source_port,omitempty"`
	// Destination port.
	DestPort string `json:"dest_port,omitempty"`
	// Service group ID. This parameter cannot be left blank when type is set to 1 (associated IP address group).
	ServiceSetID string `json:"service_set_id,omitempty"`
	// Service group name. This parameter cannot be left blank when type is set to 1 (associated IP address group).
	ServiceSetName string `json:"service_set_name,omitempty"`
	// Custom service list.
	CustomService []ServiceItem `json:"custom_service,omitempty"`
	// Predefined service group ID list.
	PredefinedGroup []string `json:"predefined_group,omitempty"`
	// Service group ID list.
	ServiceGroup []string `json:"service_group,omitempty"`
	// Service group name list.
	ServiceGroupNames []ServiceGroupVO `json:"service_group_names,omitempty"`
	// Service group type (0: user-defined service group, 1: common web service, 2: common remote login & ping, 3: common database).
	ServiceSetType int `json:"service_set_type,omitempty"`
}

type ServiceItem struct {
	// Protocol type (6: TCP, 17: UDP, 1: ICMP, 58: ICMPv6, -1: any).
	// It cannot be left blank when RuleServiceDto.type is set to 0 (manual).
	Protocol int `json:"protocol,omitempty"`
	// Source port.
	SourcePort string `json:"source_port,omitempty"`
	// Destination port.
	DestPort string `json:"dest_port,omitempty"`
	// Service member description.
	Description string `json:"description,omitempty"`
	// Service member name.
	Name string `json:"name,omitempty"`
}

type ServiceGroupVO struct {
	// Service group name.
	Name string `json:"name,omitempty"`
	// Protocol list. Protocol type: 6 (TCP), 17 (UDP), 1 (ICMP), 58 (ICMPv6), or -1 (any).
	Protocols []int `json:"protocols,omitempty"`
	// Service group type (0: user-defined service group, 1: predefined service group).
	ServiceSetType int `json:"service_set_type,omitempty"`
	// Service group ID.
	SetID string `json:"set_id,omitempty"`
}

type TagsVO struct {
	// Rule tag ID.
	TagID string `json:"tag_id,omitempty"`
	// Rule tag key.
	TagKey string `json:"tag_key,omitempty"`
	// Rule tag value.
	TagValue string `json:"tag_value,omitempty"`
}

// This function is used to create an ACL rule.
func CreateACLRule(client *golangsdk.ServiceClient, opts CreateACLRuleOpts) ([]RuleId, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/acl-rule
	raw, err := client.Post(client.ServiceURL("acl-rule"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res CreateResponse
	return res.Data.Rules, extract.Into(raw.Body, &res)
}

type CreateResponse struct {
	// Data of the return value for creating a rule.
	Data RuleIdList `json:"data"`
}

type RuleIdList struct {
	// Rule ID list
	Rules []RuleId `json:"rules"`
}

type RuleId struct {
	// Rule ID
	ID string `json:"id"`
	// Rule name
	Name string `json:"name"`
}
