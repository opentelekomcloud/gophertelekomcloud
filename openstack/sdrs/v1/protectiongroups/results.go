package protectiongroups

type ServerGroupResponseInfo struct {
	// ID specifies the ID of a protection group.
	Id string `json:"id"`
	// Name specifies the name of a protection group.
	Name string `json:"name"`
	// Description specifies the description of a protection group.
	Description string `json:"description"`
	// Status specifies the status of a protection group.
	Status string `json:"status"`
	// Progress specifies the synchronization progress of a protection group. Unit %
	Progress int `json:"progress"`
	// SourceAvailabilityZone specifies the production site AZ configured when a protection group is created.
	SourceAvailabilityZone string `json:"source_availability_zone"`
	// TargetAvailabilityZone specifies the DR site AZ configured when a protection group is created.
	TargetAvailabilityZone string `json:"target_availability_zone"`
	// DomainID specifies the ID of an active-active domain.
	DomainID string `json:"domain_id"`
	// DomainName specifies the name of an active-active domain.
	DomainName string `json:"domain_name"`
	// ProtectedStatus specifies whether protection is enabled or not.
	ProtectedStatus string `json:"protected_status"`
	// ReplicationStatus specifies the data synchronization status.
	ReplicationStatus string `json:"replication_status"`
	// HealthStatus specifies the health status of a protection group.
	HealthStatus string `json:"health_status"`
	// PriorityStation specifies the current production site of a protection group.
	PriorityStation string `json:"priority_station"`
	// ProtectedInstanceNum specifies the number of protected instances in a protection group.
	ProtectedInstanceNum int `json:"protected_instance_num"`
	// ReplicationNum specifies the number of replication pairs in a protection group.
	ReplicationNum int `json:"replication_num"`
	// DisasterRecoveryDrillNum specifies the number of DR drills in a protection group.
	DisasterRecoveryDrillNum int `json:"disaster_recovery_drill_num"`
	// SourceVPCID specifies the ID of the VPC for the production site.
	SourceVPCID string `json:"source_vpc_id"`
	// TargetVPCID specifies the ID of the VPC for the DR site.
	TargetVPCID string `json:"target_vpc_id"`
	// TestVPCID specifies the ID of the VPC used for a DR drill.
	TestVPCID string `json:"test_vpc_id"`
	// DRType specifies the deployment model. The default value is migration, indicating migration within a VPC.
	DRType string `json:"dr_type"`
	// ServerType specifies the type of managed servers.
	ServerType string `json:"server_type"`
	// CreatedAt specifies the time when a protection group was created. Format: yyyy-MM-dd HH:mm:ss.S
	CreatedAt string `json:"created_at"`
	// UpdatedAt specifies the time when a protection group was updated. Format: yyyy-MM-dd HH:mm:ss.S
	UpdatedAt string `json:"updated_at"`
	// ProtectionType specifies the protection mode.
	ProtectionType string `json:"protection_type"`
	// ReplicationModel specifies the protection mode.
	ReplicationModel string `json:"replication_model"`
}
