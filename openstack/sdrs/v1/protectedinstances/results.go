package protectedinstances

type Instance struct {
	// Instance ID
	ID string `json:"id"`
	// Instance Name
	Name string `json:"name"`
	// Instance Description
	Description string `json:"description"`
	// Protection Group ID
	GroupID string `json:"server_group_id"`
	// Instance Status
	Status string `json:"status"`
	// Instance Progress
	Progress int `json:"progress"`
	// Source Server
	SourceServer string `json:"source_server"`
	// Target Server
	TargetServer string `json:"target_server"`
	// Instance CreatedAt time
	CreatedAt string `json:"created_at"`
	// Instance UpdatedAt time
	UpdatedAt string `json:"updated_at"`
	// Production site AZ of the protection group containing the protected instance.
	PriorityStation string `json:"priority_station"`
	// Attachment
	Attachment []Attachment `json:"attachment"`
	// Tags list
	Tags []Tags `json:"tags"`
	// Metadata
	Metadata map[string]string `json:"metadata"`
}

type Tags struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type Attachment struct {
	// Replication ID
	Replication string `json:"replication"`
	// Device Name
	Device string `json:"device"`
}

// IsEmpty determines whether or not a InstancePage is empty.
func (r InstancePage) IsEmpty() (bool, error) {
	instances, err := ExtractProtectedInstances(r)
	return len(instances) == 0, err
}
