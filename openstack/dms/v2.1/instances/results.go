package instances

// CrossVpc is the structure that represents the API response of 'UpdateCrossVpc' method.
type CrossVpc struct {
	// The result of cross-VPC access modification.
	Success bool `json:"success"`
	// The result list of broker cross-VPC access modification.
	Connections []Connection `json:"results"`
}

// Connection is the structure that represents the detail of the cross-VPC access.
type Connection struct {
	// advertised.listeners IP/domain name.
	AdvertisedIp string `json:"advertised_ip"`
	// The status of broker cross-VPC access modification.
	Success bool `json:"success"`
	// Listeners IP.
	ListenersIp string `json:"ip"`
}
