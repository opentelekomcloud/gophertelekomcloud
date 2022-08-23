package groups_hcs

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// CreateGroupBuilder is an interface from which can build the request of creating group
type CreateOptsBuilder interface {
	ToGroupCreateMap() (map[string]interface{}, error)
}

// CreateGroupOps is a struct contains the parameters of creating group
type CreateOpts struct {
	Name                      string              `json:"scaling_group_name" required:"true"`
	ConfigurationID           string              `json:"scaling_configuration_id,omitempty"`
	DesireInstanceNumber      int                 `json:"desire_instance_number,omitempty"`
	MinInstanceNumber         int                 `json:"min_instance_number,omitempty"`
	MaxInstanceNumber         int                 `json:"max_instance_number,omitempty"`
	CoolDownTime              int                 `json:"cool_down_time,omitempty"`
	LBListenerID              string              `json:"lb_listener_id,omitempty"`
	LBaaSListeners            []LBaaSListenerOpts `json:"lbaas_listeners,omitempty"`
	AvailableZones            []string            `json:"available_zones,omitempty"`
	Networks                  []NetworkOpts       `json:"networks" required:"ture"`
	SecurityGroup             []SecurityGroupOpts `json:"security_groups" required:"ture"`
	VpcID                     string              `json:"vpc_id" required:"ture"`
	HealthPeriodicAuditMethod string              `json:"health_periodic_audit_method,omitempty"`
	HealthPeriodicAuditTime   int                 `json:"health_periodic_audit_time,omitempty"`
	InstanceTerminatePolicy   string              `json:"instance_terminate_policy,omitempty"`
	Notifications             []string            `json:"notifications,omitempty"`
	IsDeletePublicip          bool                `json:"delete_publicip,omitempty"`
}

type NetworkOpts struct {
	ID string `json:"id,omitempty"`
}

type SecurityGroupOpts struct {
	ID string `json:"id,omitempty"`
}

type LBaaSListenerOpts struct {
	ListenerID   string `json:"listener_id" required:"true"`
	ProtocolPort int    `json:"protocol_port" required:"true"`
	Weight       int    `json:"weight,omitempty"`
}

func (opts CreateOpts) ToGroupCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

type ListOptsBuilder interface {
	ToGroupListQuery() (string, error)
}

type ListOpts struct {
	Name            string `q:"scaling_group_name"`
	ConfigurationID string `q:"scaling_configuration_id"`
	Status          string `q:"scaling_group_status"`
}

// ToGroupListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToGroupListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

// UpdateOptsBuilder is an interface which can build the map paramter of update function
type UpdateOptsBuilder interface {
	ToGroupUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts is a struct which represents the parameters of update function
type UpdateOpts struct {
	Name                      string              `json:"scaling_group_name,omitempty"`
	DesireInstanceNumber      int                 `json:"desire_instance_number"`
	MinInstanceNumber         int                 `json:"min_instance_number"`
	MaxInstanceNumber         int                 `json:"max_instance_number"`
	CoolDownTime              int                 `json:"cool_down_time,omitempty"`
	LBListenerID              string              `json:"lb_listener_id,omitempty"`
	LBaaSListeners            []LBaaSListenerOpts `json:"lbaas_listeners,omitempty"`
	AvailableZones            []string            `json:"available_zones,omitempty"`
	Networks                  []NetworkOpts       `json:"networks,omitempty"`
	SecurityGroup             []SecurityGroupOpts `json:"security_groups,omitempty"`
	HealthPeriodicAuditMethod string              `json:"health_periodic_audit_method,omitempty"`
	HealthPeriodicAuditTime   int                 `json:"health_periodic_audit_time,omitempty"`
	InstanceTerminatePolicy   string              `json:"instance_terminate_policy,omitempty"`
	Notifications             []string            `json:"notifications,omitempty"`
	IsDeletePublicip          bool                `json:"delete_publicip,omitempty"`
	ConfigurationID           string              `json:"scaling_configuration_id,omitempty"`
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
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(client.ServiceURL("scaling_group", id, "action"), &b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
