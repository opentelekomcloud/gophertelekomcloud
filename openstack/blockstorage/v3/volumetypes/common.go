package volumetypes

// VolumeType contains information associated with an OpenStack volume type.
type VolumeType struct {
	// ID is the unique identifier for the volume type.
	ID string `json:"id"`
	// Name is the human-readable display name for the volume type.
	Name string `json:"name"`
	// Description is the human-readable description for the volume type.
	Description string `json:"description"`
	// ExtraSpecs contains arbitrary key-value pairs defined by the user.
	ExtraSpecs map[string]string `json:"extra_specs"`
	// IsPublic indicates whether the volume type is publicly visible.
	IsPublic bool `json:"is_public"`
	// QosSpecID is the associated QoS specification ID.
	QosSpecID string `json:"qos_specs_id"`
	// PublicAccess is the extended public visibility attribute.
	PublicAccess bool `json:"os-volume-type-access:is_public"`
}
