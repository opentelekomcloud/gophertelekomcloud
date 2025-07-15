package addressgroup

type GetGroupMemberQueryParameters struct {
	// Address group ID. It is the same as ID retuned while creating an address group.
	SetID string `q:"set_id" required:"true"`
	// Offset, which specifies the start position of the record to be returned. The value must be a number no less than 0. The default value is 0.
	Offset string `q:"offset" required:"true"`
	// Number of records displayed on each page. The value ranges from 1 to 1024.
	Limit int `q:"limit" required:"true"`
	// IP address.
	Address string `q:"address,omitempty"`
}

/*
########################## RESPONSE STRUCTS ##############################
*/

type AddressSetId struct {
	// Address group ID.
	Id string `json:"id"`
	// IP address group name.
	Name string `json:"name"`
}

type GetGroupMembersResponse struct {
	// Returned data for querying address group members.
	Data AddressGroupMembersData `json:"data"`
}

type AddressGroupMembersData struct {
	// Number of records displayed on each page. The value ranges from 1 to 1024.
	Limit int `json:"limit"`
	// Offset, which specifies the start position of the record to be returned.
	Offset int `json:"offset"`
	// Query the total number of blacklist/whitelist records.
	Total int `json:"total"`
	// Address Group ID.
	SetId string `json:"set_id"`
	// The list of Blacklist/Whitelist records
	Records []GroupMemberRecord `json:"records"`
}

type GroupMemberRecord struct {
	// ID of an address group member.
	ItemID string `json:"item_id"`
	// Name of an address group member.
	Name string `json:"name"`
	// Internet protocol type of an address: 0 (IPv4), 1 (IPv6).
	AddressType int `json:"address_type"`
	// IP address.
	Address string `json:"address"`
	// Description.
	Description string `json:"description"`
}
