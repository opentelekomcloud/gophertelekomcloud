package groups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

type ListOptsBuilder interface {
	ToGroupListQuery() (string, error)
}

type ListOpts struct {
	Name                string `q:"scaling_group_name"`
	ConfigurationID     string `q:"scaling_configuration_id"`
	Status              string `q:"scaling_group_status"`
	StartNumber         int    `q:"start_number"`
	Limit               int    `q:"limit"`
	EnterpriseProjectID string `q:"enterprise_project_id"`
}

// ToGroupListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToGroupListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

// UpdateOptsBuilder is an interface which can build the map parameter of update function
type UpdateOptsBuilder interface {
	ToGroupUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts is a struct which represents the parameters of update function
type UpdateOpts struct {
	Name                      string          `json:"scaling_group_name,omitempty"`
	DesireInstanceNumber      int             `json:"desire_instance_number"`
	MinInstanceNumber         int             `json:"min_instance_number"`
	MaxInstanceNumber         int             `json:"max_instance_number"`
	CoolDownTime              int             `json:"cool_down_time,omitempty"`
	LBListenerID              string          `json:"lb_listener_id,omitempty"`
	LBaaSListeners            []LBaaSListener `json:"lbaas_listeners,omitempty"`
	AvailableZones            []string        `json:"available_zones,omitempty"`
	Networks                  []ID            `json:"networks,omitempty"`
	SecurityGroup             []ID            `json:"security_groups"`
	HealthPeriodicAuditMethod string          `json:"health_periodic_audit_method,omitempty"`
	HealthPeriodicAuditTime   int             `json:"health_periodic_audit_time,omitempty"`
	HealthPeriodicAuditGrace  int             `json:"health_periodic_audit_grace_period,omitempty"`
	InstanceTerminatePolicy   string          `json:"instance_terminate_policy,omitempty"`
	Notifications             []string        `json:"notifications,omitempty"`
	IsDeletePublicip          *bool           `json:"delete_publicip,omitempty"`
	IsDeleteVolume            *bool           `json:"delete_volume,omitempty"`
	ConfigurationID           string          `json:"scaling_configuration_id,omitempty"`
	EnterpriseProjectID       string          `json:"enterprise_project_id,omitempty"`
	MultiAZPriorityPolicy     string          `json:"multi_az_priority_policy,omitempty"`
}

func (opts UpdateOpts) ToGroupUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

type ActionOptsBuilder interface {
	ToActionMap() (map[string]interface{}, error)
}

type ActionOpts struct {
	Action string `json:"action" required:"true"`
}

func (opts ActionOpts) ToActionMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func doAction(client *golangsdk.ServiceClient, id string, opts ActionOptsBuilder) (r ActionResult) {
	b, err := opts.ToActionMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(enableURL(client, id), &b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
