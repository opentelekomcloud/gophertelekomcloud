package pools

import "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/structs"

type Pool struct {
	LBMethod                 string                `json:"lb_algorithm"`
	Protocol                 string                `json:"protocol"`
	Description              string                `json:"description"`
	Listeners                []structs.ResourceRef `json:"listeners"`
	Members                  []structs.ResourceRef `json:"members"`
	MonitorID                string                `json:"healthmonitor_id"`
	AdminStateUp             bool                  `json:"admin_state_up"`
	Name                     string                `json:"name"`
	ProjectID                string                `json:"project_id"`
	ID                       string                `json:"id"`
	Loadbalancers            []structs.ResourceRef `json:"loadbalancers"`
	Persistence              *SessionPersistence   `json:"session_persistence"`
	IpVersion                string                `json:"ip_version"`
	SlowStart                *SlowStart            `json:"slow_start"`
	DeletionProtectionEnable bool                  `json:"member_deletion_protection_enable"`
	VpcId                    string                `json:"vpc_id"`
	Type                     string                `json:"type"`
	CreatedAt                string                `json:"created_at"`
	UpdatedAt                string                `json:"updated_at"`
	ProtectionStatus         string                `json:"protection_status"`
	ProtectionReason         string                `json:"protection_reason"`
	AZAffinity               *AZAffinity           `json:"az_affinity"`
}

type SessionPersistence struct {
	Type               string `json:"type" required:"true"`
	CookieName         string `json:"cookie_name,omitempty"`
	PersistenceTimeout int    `json:"persistence_timeout,omitempty"`
}

type SlowStart struct {
	Enable   bool `json:"enable" required:"true"`
	Duration int  `json:"duration" required:"true"`
}

type AZAffinity struct {
	Enable                           bool   `json:"enable"`
	AZMinimumHealthyMemberPercentage int    `json:"az_minimum_healthy_member_percentage"`
	AZMinimumHealthyMemberCount      int    `json:"az_minimum_healthy_member_count"`
	AZUnhealthyFallbackStrategy      string `json:"az_unhealthy_fallback_strategy"`
}
