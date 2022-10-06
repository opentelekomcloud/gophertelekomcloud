package groups

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
	SecurityGroups            []ID            `json:"security_groups"`
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
	ID            string `json:"id"`
	IPv6Enable    bool   `json:"ipv6_enable"`
	IPv6Bandwidth ID     `json:"ipv6_bandwidth"`
}
