package groups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// CreateResult is a struct returned by CreateGroup request
type CreateResult struct {
	golangsdk.Result
}

// Extract the create group result as a string type.
func (r CreateResult) Extract() (string, error) {
	var s string
	err := r.ExtractIntoStructPtr(s, "scaling_group_id")
	return s, err
}

// DeleteResult contains the body of the deleting group request
type DeleteResult struct {
	golangsdk.ErrResult
}

// GetResult contains the body of getting detailed group request
type GetResult struct {
	golangsdk.Result
}

// Extract method will parse the result body into Group struct
func (r GetResult) Extract() (*Group, error) {
	s := new(Group)
	err := r.ExtractIntoStructPtr(s, "scaling_group")
	return s, err
}

// Group represents the struct of one autoscaling group
type Group struct {
	Name                      string          `json:"scaling_group_name"`
	ID                        string          `json:"scaling_group_id"`
	Status                    string          `json:"scaling_group_status"`
	ConfigurationID           string          `json:"scaling_configuration_id"`
	ConfigurationName         string          `json:"scaling_configuration_name"`
	ActualInstanceNumber      int             `json:"current_instance_number"`
	DesireInstanceNumber      int             `json:"desire_instance_number"`
	MinInstanceNumber         int             `json:"min_instance_number"`
	MaxInstanceNumber         int             `json:"max_instance_number"`
	CoolDownTime              int             `json:"cool_down_time"`
	LBListenerID              string          `json:"lb_listener_id"`
	LBaaSListeners            []LBaaSListener `json:"lbaas_listeners"`
	AvailableZones            []string        `json:"available_zones"`
	Networks                  []Network       `json:"networks"`
	SecurityGroups            []SecurityGroup `json:"security_groups"`
	CreateTime                string          `json:"create_time"`
	VpcID                     string          `json:"vpc_id"`
	Detail                    string          `json:"detail"`
	IsScaling                 bool            `json:"is_scaling"`
	HealthPeriodicAuditMethod string          `json:"health_periodic_audit_method"`
	HealthPeriodicAuditTime   int             `json:"health_periodic_audit_time"`
	HealthPeriodicAuditGrace  int             `json:"health_periodic_audit_grace_period"`
	InstanceTerminatePolicy   string          `json:"instance_terminate_policy"`
	Notifications             []string        `json:"notifications"`
	DeletePublicIP            bool            `json:"delete_publicip"`
	DeleteVolume              bool            `json:"delete_volume"`
	CloudLocationID           string          `json:"cloud_location_id"`
	EnterpriseProjectID       string          `json:"enterprise_project_id"`
	ActivityType              string          `json:"activity_type"`
	MultiAZPriorityPolicy     string          `json:"multi_az_priority_policy"`
}

type Network struct {
	ID            string        `json:"id"`
	IPv6Enable    bool          `json:"ipv6_enable"`
	IPv6Bandwidth IPv6Bandwidth `json:"ipv6_bandwidth"`
}

type IPv6Bandwidth struct {
	ID string `json:"id"`
}

type SecurityGroup struct {
	ID string `json:"id"`
}

type LBaaSListener struct {
	ListenerID   string `json:"listener_id"`
	PoolID       string `json:"pool_id"`
	ProtocolPort int    `json:"protocol_port"`
	Weight       int    `json:"weight"`
}

type GroupPage struct {
	pagination.SinglePageBase
}

// IsEmpty returns true if a ListResult contains no Volumes.
func (r GroupPage) IsEmpty() (bool, error) {
	groups, err := ExtractGroups(r)
	return len(groups) == 0, err
}

// ExtractGroups returns a slice of AS Groups contained in a
// single page of results.
func ExtractGroups(r pagination.Page) ([]Group, error) {
	var s []Group
	err := (r.(GroupPage)).ExtractIntoSlicePtr(&s, "scaling_groups")
	return s, err
}

// UpdateResult is a struct from which can get the result of update method
type UpdateResult struct {
	golangsdk.Result
}

// Extract will deserialize the result to group id with string
func (r UpdateResult) Extract() (string, error) {
	var s string
	err := r.ExtractIntoStructPtr(s, "scaling_group_id")
	return s, err
}

// ActionResult this is the action result which is the result of enable or disable operations
type ActionResult struct {
	golangsdk.ErrResult
}
