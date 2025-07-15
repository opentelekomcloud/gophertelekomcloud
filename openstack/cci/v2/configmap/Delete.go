package configmap

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// DeleteOpts contains all the values needed to delete a ConfigMap
type DeleteOpts struct {
	// APIVersion defines the versioned schema of this representation of an object
	APIVersion string `json:"apiVersion,omitempty"`

	// DryRun when present, indicates that modifications should not be persisted
	DryRun []string `json:"dryRun,omitempty"`

	// GracePeriodSeconds the duration in seconds before the object should be deleted
	GracePeriodSeconds *int64 `json:"gracePeriodSeconds,omitempty"`

	// Kind is a string value representing the REST resource this object represents
	Kind string `json:"kind,omitempty"`

	// OrphanDependents should the dependent objects be orphaned
	OrphanDependents *bool `json:"orphanDependents,omitempty"`

	// Preconditions must be fulfilled before a deletion is carried out
	Preconditions *Preconditions `json:"preconditions,omitempty"`

	// PropagationPolicy whether and how garbage collection will be performed
	PropagationPolicy string `json:"propagationPolicy,omitempty"`
}

// Preconditions must be fulfilled before an operation (update, delete, etc.) is carried out
type Preconditions struct {
	// ResourceVersion specifies the target ResourceVersion
	ResourceVersion string `json:"resourceVersion,omitempty"`

	// UID specifies the target UID
	UID string `json:"uid,omitempty"`
}

// Delete removes a ConfigMap
func Delete(client *golangsdk.ServiceClient, namespace string, name string, opts DeleteOpts) (*DeleteResult, error) {
	url := client.ServiceURL("namespaces", namespace, "configmaps", name)

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.DeleteWithBody(url, b, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})
	if err != nil {
		return nil, err
	}

	var res DeleteResult
	err = extract.Into(raw.Body, &res)
	return &res, err
}

// DeleteResult represents the API response after deletion
type DeleteResult struct {
	// APIVersion defines the versioned schema of this representation of an object
	APIVersion string `json:"apiVersion,omitempty"`

	// BinaryData contains the binary data
	BinaryData map[string]string `json:"binaryData,omitempty"`

	// Data contains the configuration data
	Data map[string]string `json:"data,omitempty"`

	// Immutable if set to true, ensures that data stored in the ConfigMap cannot be updated
	Immutable *bool `json:"immutable,omitempty"`

	// Kind is a string value representing the REST resource this object represents
	Kind string `json:"kind,omitempty"`

	// Metadata contains the object metadata
	Metadata ObjectMeta `json:"metadata,omitempty"`
}
