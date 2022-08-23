package groups_hcs

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
	Networks                  []Network       `json:"networks" required:"ture"`
	SecurityGroup             []SecurityGroup `json:"security_groups" required:"ture"`
	VpcID                     string          `json:"vpc_id" required:"ture"`
	HealthPeriodicAuditMethod string          `json:"health_periodic_audit_method,omitempty"`
	HealthPeriodicAuditTime   int             `json:"health_periodic_audit_time,omitempty"`
	InstanceTerminatePolicy   string          `json:"instance_terminate_policy,omitempty"`
	Notifications             []string        `json:"notifications,omitempty"`
	IsDeletePublicip          bool            `json:"delete_publicip,omitempty"`
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
		GroupID string `json:"scaling_group_id"`
	}
	err = extract.Into(raw.Body, res)
	return res.GroupID, err
}
