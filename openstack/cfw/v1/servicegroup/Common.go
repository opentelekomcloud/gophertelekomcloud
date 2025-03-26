package servicegroup

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
