package namespace

// Metadata contains the metadata for the namespace
type Metadata struct {
	// Name must be unique within a namespace
	Name string `json:"name,omitempty"`

	// Annotations is an unstructured key value map stored with a resource
	Annotations map[string]string `json:"annotations,omitempty"`

	// Labels contains map of string keys and values
	Labels map[string]string `json:"labels,omitempty"`

	// ClusterName is the name of the cluster which the object belongs to
	ClusterName string `json:"clusterName,omitempty"`

	// CreationTimestamp is a timestamp representing the server time when this object was created
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// DeletionGracePeriodSeconds is the duration in seconds before the object should be deleted
	DeletionGracePeriodSeconds *int64 `json:"deletionGracePeriodSeconds,omitempty"`

	// DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted
	DeletionTimestamp string `json:"deletionTimestamp,omitempty"`

	// Enable identifies whether the resource is available
	Enable *bool `json:"enable,omitempty"`

	// Finalizers must be empty before the object is deleted from the registry
	Finalizers []string `json:"finalizers,omitempty"`

	// GenerateName is an optional prefix for generated names
	GenerateName string `json:"generateName,omitempty"`

	// Generation is a sequence number representing a specific generation of the desired state
	Generation *int64 `json:"generation,omitempty"`

	// ManagedFields contains info managed by controllers
	ManagedFields []ManagedFieldsEntry `json:"managedFields,omitempty"`

	// Namespace defines the space within which each name must be unique
	Namespace string `json:"namespace,omitempty"`

	// OwnerReferences contains the list of objects depended by this object
	OwnerReferences []OwnerReference `json:"ownerReferences,omitempty"`

	// ResourceVersion is an opaque value that represents the internal version of this object
	ResourceVersion string `json:"resourceVersion,omitempty"`

	// SelfLink is a URL representing this object
	SelfLink string `json:"selfLink,omitempty"`

	// UID is the unique in time and space value for this object
	UID string `json:"uid,omitempty"`
}

// ManagedFieldsEntry contains information about fields managed by controllers
type ManagedFieldsEntry struct {
	// APIVersion defines the version of this resource that this field set applies to
	APIVersion string `json:"apiVersion,omitempty"`

	// FieldsType is the discriminator for the different fields format and version
	FieldsType string `json:"fieldsType,omitempty"`

	// FieldsV1 holds the first JSON version format
	FieldsV1 interface{} `json:"fieldsV1,omitempty"`

	// Manager is an identifier of the workflow managing these fields
	Manager string `json:"manager,omitempty"`

	// Operation is the type of operation which lead to this ManagedFieldsEntry being created
	Operation string `json:"operation,omitempty"`

	// Subresource is the name of the subresource used to update that object, or empty if the object was updated through the main resource
	Subresource string `json:"subresource,omitempty"`

	// Time is timestamp of when these fields were set
	Time string `json:"time,omitempty"`
}

// OwnerReference contains the information to let you identify an owning object
type OwnerReference struct {
	// APIVersion of the referent
	APIVersion string `json:"apiVersion" required:"true"`

	// Kind of the referent
	Kind string `json:"kind" required:"true"`

	// Name of the referent
	Name string `json:"name" required:"true"`

	// UID of the referent
	UID string `json:"uid" required:"true"`

	// If true, AND if the owner has the "foregroundDeletion" finalizer, then the owner cannot be deleted
	BlockOwnerDeletion *bool `json:"blockOwnerDeletion,omitempty"`

	// If true, this reference points to the managing controller
	Controller *bool `json:"controller,omitempty"`
}

// NamespaceSpec defines the behavior of the Namespace
type NamespaceSpec struct {
	// Finalizers is an opaque list of values that must be empty to permanently remove object from storage
	Finalizers []string `json:"finalizers,omitempty"`
}

// NamespaceStatus describes the current status of a Namespace
type NamespaceStatus struct {
	// Phase is the current lifecycle phase of the namespace
	Phase string `json:"phase,omitempty"`

	// Conditions represents the latest available observations of a namespace's current state
	Conditions []NamespaceCondition `json:"conditions,omitempty"`
}

// NamespaceCondition contains details for the current condition of this namespace
type NamespaceCondition struct {
	// Type of namespace controller condition
	Type string `json:"type" required:"true"`

	// Status of the condition, one of True, False, Unknown
	Status string `json:"status" required:"true"`

	// Last transition time
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`

	// The reason for the condition's last transition
	Reason string `json:"reason,omitempty"`

	// A human readable message indicating details about the transition
	Message string `json:"message,omitempty"`
}

// Namespace represents the response structure for namespace operations
type Namespace struct {
	// APIVersion defines the versioned schema of this representation of an object
	APIVersion string `json:"apiVersion"`

	// Kind is a string value representing the REST resource this object represents
	Kind string `json:"kind"`

	// Metadata contains the namespace metadata
	Metadata Metadata `json:"metadata"`

	// Spec defines the behavior of the Namespace
	Spec NamespaceSpec `json:"spec"`

	// Status describes the current status of a Namespace
	Status NamespaceStatus `json:"status"`
}
