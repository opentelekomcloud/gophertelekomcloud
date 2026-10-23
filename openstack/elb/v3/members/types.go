package members

import "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/structs"

type Member struct {
	Name             string                `json:"name"`
	Weight           int                   `json:"weight"`
	AdminStateUp     bool                  `json:"admin_state_up"`
	ProjectID        string                `json:"project_id"`
	SubnetID         string                `json:"subnet_cidr_id"`
	PoolID           string                `json:"pool_id"`
	Address          string                `json:"address"`
	ProtocolPort     int                   `json:"protocol_port"`
	ID               string                `json:"id"`
	IpVersion        string                `json:"ip_version"`
	OperatingStatus  string                `json:"operating_status"`
	AvailabilityZone string                `json:"availability_zone"`
	Status           []MemberStatus        `json:"status"`
	CreatedAt        string                `json:"created_at"`
	UpdatedAt        string                `json:"updated_at"`
	MemberType       string                `json:"member_type"`
	InstanceID       string                `json:"instance_id"`
	DeviceOwner      string                `json:"device_owner"`
	DeviceID         string                `json:"device_id"`
	LoadbalancerID   string                `json:"loadbalancer_id"`
	Loadbalancers    []structs.ResourceRef `json:"loadbalancers"`
}

type MemberStatus struct {
	ListenerID      string                         `json:"listener_id"`
	OperatingStatus string                         `json:"operating_status"`
	Reason          *MemberHealthCheckFailedReason `json:"reason"`
	CreatedAt       string                         `json:"created_at"`
	UpdatedAt       string                         `json:"updated_at"`
}

type MemberHealthCheckFailedReason struct {
	ReasonCode          string `json:"reason_code"`
	ExpectedResponse    string `json:"expected_response"`
	HealthcheckResponse string `json:"healthcheck_response"`
}

type PageInfo struct {
	PreviousMarker string `json:"previous_marker"`
	NextMarker     string `json:"next_marker"`
	CurrentCount   int    `json:"current_count"`
}
