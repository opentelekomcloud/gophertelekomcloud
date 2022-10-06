package groups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	Name                      string          `json:"scaling_group_name" required:"true"`
	ConfigurationID           string          `json:"scaling_configuration_id,omitempty"`
	DesireInstanceNumber      int             `json:"desire_instance_number,omitempty"`
	MinInstanceNumber         int             `json:"min_instance_number,omitempty"`
	MaxInstanceNumber         int             `json:"max_instance_number,omitempty"`
	CoolDownTime              int             `json:"cool_down_time,omitempty"`
	LBListenerID              string          `json:"lb_listener_id,omitempty"`
	LBaaSListeners            []LBaaSListener `json:"lbaas_listeners,omitempty"`
	AvailableZones            []string        `json:"available_zones,omitempty"`
	Networks                  []ID            `json:"networks" required:"true"`
	SecurityGroup             []ID            `json:"security_groups,omitempty"`
	VpcID                     string          `json:"vpc_id" required:"true"`
	HealthPeriodicAuditMethod string          `json:"health_periodic_audit_method,omitempty"`
	HealthPeriodicAuditTime   int             `json:"health_periodic_audit_time,omitempty"`
	HealthPeriodicAuditGrace  int             `json:"health_periodic_audit_grace_period,omitempty"`
	InstanceTerminatePolicy   string          `json:"instance_terminate_policy,omitempty"`
	Notifications             []string        `json:"notifications,omitempty"`
	IsDeletePublicip          *bool           `json:"delete_publicip,omitempty"`
	IsDeleteVolume            *bool           `json:"delete_volume,omitempty"`
	EnterpriseProjectID       string          `json:"enterprise_project_id,omitempty"`
	MultiAZPriorityPolicy     string          `json:"multi_az_priority_policy,omitempty"`
}

type LBaaSListener struct {
	ListenerID   string `json:"listener_id"`
	PoolID       string `json:"pool_id" required:"true"`
	ProtocolPort int    `json:"protocol_port" required:"true"`
	Weight       int    `json:"weight" required:"true"`
}

type ID struct {
	ID string `json:"id" required:"true"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (string, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return "", err
	}

	raw, err := client.Post(client.ServiceURL("scaling_group"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		ID string `json:"scaling_group_id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.ID, err
}
