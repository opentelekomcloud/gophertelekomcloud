package blackwhitelist

type ListQueryParameters struct {
	// Protected object ID, which is used to distinguish between Internet border protection and VPC border protection
	// after a cloud firewall is created. You can obtain the ID by calling the Get function in management package.
	// In the return value, find the ID in ProtectObjects[n].ObjectID.
	// If the value of type is 0, the protected object ID belongs to the Internet border.
	// If the value of type is 1, the protected object ID belongs to the VPC border.
	ObjectID string `q:"object_id" required:"true"`
	// Blacklist/Whitelist type: 4 (blacklist), 5 (whitelist).
	ListType int `q:"list_type" required:"true"`
	// IP address.
	Address string `q:"address,omitempty"`
	// Offset, which specifies the start position of the record to be returned. The value must be a number no less than 0. The default value is 0.
	Offset string `q:"offset" required:"true"`
	// Number of records displayed on each page. The value ranges from 1 to 1024.
	Limit int `q:"limit" required:"true"`
}

/*
########################## RESPONSE STRUCTS ##############################
*/

type BlackWhiteListId struct {
	// Blacklist/Whitelist ID.
	Id string `json:"id"`
	// Blacklist/Whitelist name.Which is the Address
	Name string `json:"name"`
}

type GetListResponse struct {
	// Return value for querying the blacklist/whitelist.
	Data BlacklistWhitelistQueryResponseData `json:"data"`
}

type BlacklistWhitelistQueryResponseData struct {
	// Number of records displayed on each page. The value ranges from 1 to 1024.
	Limit int `json:"limit"`
	// Offset, which specifies the start position of the record to be returned.
	Offset int `json:"offset"`
	// Query the total number of blacklist/whitelist records.
	Total int `json:"total"`
	// The list of Blacklist/Whitelist records
	Records []BlackWhiteListRecord `json:"records"`
}

type BlackWhiteListRecord struct {
	// Blacklist/Whitelist ID.
	ListId string `json:"list_id"`
	// Address direction: 0 (source), 1 (destination).
	Direction int `json:"direction"`
	// Internet protocol type of an address: 0 (IPv4), 1 (IPv6).
	AddressType int `json:"address_type"`
	// IP address.
	Address string `json:"address"`
	// Protocol type: 6 (TCP), 17 (UDP), 1 (ICMP), 58 (ICMPv6), or -1 (any).
	// Cannot be left blank when type is set to 0 (manual) and can be omitted when type is set to 1 (automatic).
	Protocol int `json:"protocol"`
	// Destination port.
	Port string `json:"port"`
	// Description.
	Description string `json:"description"`
}
