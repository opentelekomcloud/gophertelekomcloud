package acl

// Note for devs: This struct is almost same in create and update api with the exception of additional redundant parameter
// "address_group_names" in Update of type []AddressGroupVO which can be replaced easily by "address_group" parameter.
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

/*
########################## RESPONSE STRUCTS ##############################
*/
type RuleId struct {
	// Rule ID
	ID string `json:"id"`
	// Rule name
	Name string `json:"name"`
}
