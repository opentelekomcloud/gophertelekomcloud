package addressgroup

/*
########################## RESPONSE STRUCTS ##############################
*/

type AddressSetId struct {
	// Address group ID.
	Id string `json:"id"`
	// IP address group name.
	Name string `json:"name"`
}
