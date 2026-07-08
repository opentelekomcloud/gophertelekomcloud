package addressgroup

type IpExtraSetOption struct {
	// IP address entries in an IP address group. Both IPv4 and IPv6 address entries are supported.
	// A single IP address, for example, 192.168.21.25
	// An IP address range, for example, 192.168.21.25-192.168.21.30
	// A CIDR block, for example, 192.168.21.0/24
	Ip string `json:"ip" required:"true"`
	// Remarks of an IP address entry. The value can contain 0 to 255 characters and cannot contain angle brackets (< or >).
	Remarks string `json:"remarks,omitempty"`
}

// ################ RESPONSE STRUCTS ##################

type AddressGroupResponse struct {
	// Request ID.
	RequestID string `json:"request_id"`
	// Response body for creating an IP address group.
	AddressGroup AddressGroup `json:"address_group"`
}

type AddressGroup struct {
	// ID of an IP address group. After an IP address group is created, an IP address
	ID string `json:"id"`
	// The name of an IP address group.
	Name string `json:"name"`
	// Description about an IP address group.
	Description string `json:"description"`
	// Maximum number of IP address entries in an IP address group.
	MaxCapacity int `json:"max_capacity"`
	// IP address entries in an IP address group.
	IPSet []string `json:"ip_set"`
	// IP address version of an IP address group.
	// 4: IPv4 address group
	// 6: IPv6 address group
	IPVersion int `json:"ip_version"`
	// Time when an IP address group was created.
	// The value is a UTC time in the format of yyyy-MM-ddTHH:mm:ss.
	CreatedAt string `json:"created_at"`
	// Time when an IP address group was last updated.
	// The value is a UTC time in the format of yyyy-MM-ddTHH:mm:ss.
	UpdatedAt string `json:"updated_at"`
	// ID of the project that an IP address group belongs to.
	TenantID string `json:"tenant_id"`
	// ID of the enterprise project that an IP address group belongs to.
	// The value is 0 or a string in UUID format with hyphens (-).
	// 0 indicates the default enterprise project.
	EnterpriseProjectID string `json:"enterprise_project_id"`
	// Tags of an IP address group, including tag keys and tag values,
	// which can be used to classify and identify resources.
	Tags []ResponseTag `json:"tags"`
	// Status of an IP address group. If the IP address group is in the UPDATING
	// state, it cannot be updated again.
	// NORMAL: normal (default value)
	// UPDATING: being updated
	// UPDATE_FAILED: update failed
	Status string `json:"status"`
	// Details about the IP address group status.
	StatusMessage string `json:"status_message"`
	// IP address entries in an IP address group and their remarks.
	IPExtraSet []IpExtraSetRespOption `json:"ip_extra_set"`
}

type IpExtraSetRespOption struct {
	// IP address entry. Can be a single IP address (e.g. 192.168.21.25),
	// an IP address range (e.g. 192.168.21.25-192.168.21.30),
	// or a CIDR block (e.g. 192.168.21.0/24).
	IP string `json:"ip"`
	// Remark or description for the IP address entry.
	Remark string `json:"remark"`
}

// ResponseTag represents a tag object associated with an IP address group.
type ResponseTag struct {
	// Tag key.
	Key string `json:"key"`
	// Tag value.
	Value string `json:"value"`
}
