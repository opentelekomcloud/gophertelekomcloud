package groups

type Group struct {
	// Specifies the name of the AS group.
	Name string `json:"scaling_group_name"`
	// Specifies the AS group ID.
	ID string `json:"scaling_group_id"`
	// Specifies the status of the AS group.
	Status string `json:"scaling_group_status"`
	// Specifies the AS configuration ID.
	ConfigurationID string `json:"scaling_configuration_id"`
	// Specifies the AS configuration name.
	ConfigurationName string `json:"scaling_configuration_name"`
	// Specifies the number of current instances in the AS group.
	ActualInstanceNumber int `json:"current_instance_number"`
	// Specifies the expected number of instances in the AS group.
	DesireInstanceNumber int `json:"desire_instance_number"`
	// Specifies the minimum number of instances in the AS group.
	MinInstanceNumber int `json:"min_instance_number"`
	// Specifies the maximum number of instances in the AS group.
	MaxInstanceNumber int `json:"max_instance_number"`
	// Specifies the cooldown period (s).
	CoolDownTime int `json:"cool_down_time"`
	// Specifies the ID of a typical ELB listener. ELB listener IDs are separated using a comma (,).
	LBListenerID string `json:"lb_listener_id"`
	// Specifies enhanced load balancers.
	LBaaSListeners []LBaaSListener `json:"lbaas_listeners"`
	// Specifies the AZ information.
	AvailableZones []string `json:"available_zones"`
	// Specifies the network information.
	Networks []ID `json:"networks"`
	// Specifies the security group information.
	SecurityGroups []ID `json:"security_groups"`
	// Specifies the time when an AS group was created. The time format complies with UTC.
	CreateTime string `json:"create_time"`
	// Specifies the ID of the VPC to which the AS group belongs.
	VpcID string `json:"vpc_id"`
	// Specifies details about the AS group. If a scaling action fails, this parameter is used to record errors.
	Detail string `json:"detail"`
	// Specifies the scaling flag of the AS group.
	IsScaling bool `json:"is_scaling"`
	// Specifies the health check method.
	HealthPeriodicAuditMethod string `json:"health_periodic_audit_method"`
	// Specifies the health check interval.
	HealthPeriodicAuditTime int `json:"health_periodic_audit_time"`
	// Specifies the grace period for health check.
	HealthPeriodicAuditGrace int `json:"health_periodic_audit_grace_period"`
	// Specifies the instance removal policy.
	InstanceTerminatePolicy string `json:"instance_terminate_policy"`
	// Specifies the notification mode.
	// EMAIL refers to notification by email.
	Notifications []string `json:"notifications"`
	// Specifies whether to delete the EIP bound to the ECS when deleting the ECS.
	DeletePublicIP bool `json:"delete_publicip"`
	// Specifies whether to delete the data disks attached to the ECS when deleting the ECS.
	DeleteVolume bool `json:"delete_volume"`
	// This parameter is reserved.
	CloudLocationID string `json:"cloud_location_id"`
	// Specifies the enterprise project ID.
	EnterpriseProjectID string `json:"enterprise_project_id"`
	// Specifies the type of the AS action.
	ActivityType string `json:"activity_type"`
	// Specifies the priority policy used to select target AZs when adjusting the number of instances in an AS group.
	MultiAZPriorityPolicy string `json:"multi_az_priority_policy"`
	// Specifies the description of the AS group. The value can contain 1 to 256 characters.
	Description string `json:"description"`
}
