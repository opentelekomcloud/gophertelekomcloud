package privateips

type PrivateIP struct {
	Status      string `json:"status"`
	ID          string `json:"id"`
	SubnetID    string `json:"subnet_id"`
	TenantID    string `json:"tenant_id"`
	DeviceOwner string `json:"device_owner"`
	IPAddress   string `json:"ip_address"`
}
