package configmap

// ObjectMeta standard object metadata
type ObjectMeta struct {
	// Annotations is an unstructured key value map stored with a resource
	Annotations map[string]string `json:"annotations,omitempty"`

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

	// Finalizers is a list of identifiers for resource cleanup
	Finalizers []string `json:"finalizers,omitempty"`

	// GenerateName is an optional prefix for generating a unique name
	GenerateName string `json:"generateName,omitempty"`

	// Generation is a sequence number representing a specific generation of the desired state
	Generation *int64 `json:"generation,omitempty"`

	// Labels map of string keys and values for organization
	Labels map[string]string `json:"labels,omitempty"`

	// ManagedFields maps workflow-id and version to the set of fields
	ManagedFields []ManagedFieldsEntry `json:"managedFields,omitempty"`

	// Name must be unique within a namespace
	Name string `json:"name"`

	// Namespace defines the space within which each name must be unique
	Namespace string `json:"namespace,omitempty"`

	// OwnerReferences list of objects depended by this object
	OwnerReferences []OwnerReference `json:"ownerReferences,omitempty"`

	// ResourceVersion is an opaque value that represents the internal version of this object
	ResourceVersion string `json:"resourceVersion,omitempty"`

	// SelfLink is a URL representing this object
	SelfLink string `json:"selfLink,omitempty"`

	// UID is the unique in time and space value for this object
	UID string `json:"uid,omitempty"`
}

// ManagedFieldsEntry contains workflow-managed fields
type ManagedFieldsEntry struct {
	// APIVersion defines the version of this resource
	APIVersion string `json:"apiVersion,omitempty"`

	// FieldsType is the discriminator for the different fields format and version
	FieldsType string `json:"fieldsType,omitempty"`

	// FieldsV1 holds the first JSON version format
	FieldsV1 interface{} `json:"fieldsV1,omitempty"`

	// Manager is an identifier of the workflow managing these fields
	Manager string `json:"manager,omitempty"`

	// Operation is the type of operation which lead to this ManagedFieldsEntry
	Operation string `json:"operation,omitempty"`

	// Subresource is the name of the subresource used to update that object, or empty if the object was updated through the main resource
	Subresource string `json:"subresource,omitempty"`

	// Time is timestamp of when these fields were set
	Time string `json:"time,omitempty"`
}

// OwnerReference contains enough information to let you identify an owning object
type OwnerReference struct {
	// APIVersion of the referent
	APIVersion string `json:"apiVersion"`

	// BlockOwnerDeletion will block garbage collection of the owner
	BlockOwnerDeletion *bool `json:"blockOwnerDeletion,omitempty"`

	// Controller identifies whether this OwnerReference points to the managing controller
	Controller *bool `json:"controller,omitempty"`

	// Kind of the referent
	Kind string `json:"kind"`

	// Name of the referent
	Name string `json:"name"`

	// UID of the referent
	UID string `json:"uid"`
}

// ConfigMapResponse represents the response when creating a ConfigMap
type ConfigMap struct {
	// APIVersion defines the versioned schema of this representation of an object
	APIVersion string `json:"apiVersion"`

	// BinaryData contains the binary data
	BinaryData map[string]string `json:"binaryData"`

	// Data contains the configuration data
	Data map[string]string `json:"data"`

	// Immutable if set to true, ensures that data stored in the ConfigMap cannot be updated
	Immutable *bool `json:"immutable"`

	// Kind is a string value representing the REST resource this object represents
	Kind string `json:"kind"`

	// Metadata contains the object metadata
	Metadata ObjectMeta `json:"metadata"`
}
