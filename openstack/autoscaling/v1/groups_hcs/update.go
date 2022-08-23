package groups_hcs

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	Name                      string          `json:"scaling_group_name,omitempty"`
	DesireInstanceNumber      int             `json:"desire_instance_number"`
	MinInstanceNumber         int             `json:"min_instance_number"`
	MaxInstanceNumber         int             `json:"max_instance_number"`
	CoolDownTime              int             `json:"cool_down_time,omitempty"`
	LBListenerID              string          `json:"lb_listener_id,omitempty"`
	LBaaSListeners            []LBaaSListener `json:"lbaas_listeners,omitempty"`
	AvailableZones            []string        `json:"available_zones,omitempty"`
	Networks                  []Network       `json:"networks,omitempty"`
	SecurityGroup             []SecurityGroup `json:"security_groups,omitempty"`
	HealthPeriodicAuditMethod string          `json:"health_periodic_audit_method,omitempty"`
	HealthPeriodicAuditTime   int             `json:"health_periodic_audit_time,omitempty"`
	InstanceTerminatePolicy   string          `json:"instance_terminate_policy,omitempty"`
	Notifications             []string        `json:"notifications,omitempty"`
	IsDeletePublicip          bool            `json:"delete_publicip,omitempty"`
	ConfigurationID           string          `json:"scaling_configuration_id,omitempty"`
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (string, error) {
	body, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return "", err
	}

	raw, err := client.Put(client.ServiceURL("scaling_group", id), body, nil, &golangsdk.RequestOpts{
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
