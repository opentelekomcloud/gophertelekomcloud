package servicegroup

type GetGroupMemberQueryParameters struct {
	// Service group ID. It is the same as ID retuned while creating an service group.
	SetID string `q:"set_id" required:"true"`
	// Offset, which specifies the start position of the record to be returned. The value must be a number no less than 0. The default value is 0.
	Offset string `q:"offset" required:"true"`
	// Number of records displayed on each page. The value ranges from 1 to 1024.
	Limit int `q:"limit" required:"true"`
}

/*
########################## RESPONSE STRUCTS ##############################
*/

type ServiceSetDataResponse struct {
	Data ServiceSetId `json:"data"`
}

type ServiceSetId struct {
	// Service group ID.
	Id string `json:"id"`
	// service group name.
	Name string `json:"name"`
}

type GetGroupMembersResponse struct {
	// Returned data for querying service group members.
	Data ServiceGroupMembersData `json:"data"`
}

type ServiceGroupMembersData struct {
	// Number of records displayed on each page. The value ranges from 1 to 1024.
	Limit int `json:"limit"`
	// Offset, which specifies the start position of the record to be returned.
	Offset int `json:"offset"`
	// Query the total number of group member records.
	Total int `json:"total"`
	// Service Group ID.
	SetId string `json:"set_id"`
	// The list of group member records
	Records []GroupMemberRecord `json:"records"`
}

type GroupMemberRecord struct {
	// ID of an service group member.
	ItemID string `json:"item_id"`
	// Protocol type: 6 (TCP), 17 (UDP), 1 (ICMP), 58 (ICMPv6), or -1 (any).
	Protocol int `json:"protocol"`
	// Source port.
	SourcePort string `json:"source_port"`
	// Destination port.
	DestPort string `json:"dest_port"`
	// Description.
	Description string `json:"description"`
}
