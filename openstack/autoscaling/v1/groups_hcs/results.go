package groups_hcs

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
	InstanceTerminatePolicy   string          `json:"instance_terminate_policy"`
	Notifications             []string        `json:"notifications"`
	DeletePublicip            bool            `json:"delete_publicip"`
	CloudLocationID           string          `json:"cloud_location_id"`
}

type Network struct {
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
