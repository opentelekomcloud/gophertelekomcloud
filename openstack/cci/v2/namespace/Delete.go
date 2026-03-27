package namespace

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// DeleteOpts contains all the values needed to delete a namespace
type DeleteOpts struct {
	// Name of the namespace to delete
	Name string `json:"-" required:"true"`

	// Body contains the delete options
	Body DeleteBody `json:"-" required:"true"`
}

// DeleteBody represents the body of delete request
type DeleteBody struct {
	// APIVersion defines the versioned schema
	APIVersion string `json:"apiVersion,omitempty"`

	// Kind represents the REST resource
	Kind string `json:"kind,omitempty"`

	// GracePeriodSeconds is the duration in seconds before deletion
	GracePeriodSeconds *int64 `json:"gracePeriodSeconds,omitempty"`

	// PropagationPolicy determines how garbage collection will be performed
	PropagationPolicy string `json:"propagationPolicy,omitempty"`

	// DryRun when present, indicates modifications should not be persisted
	DryRun []string `json:"dryRun,omitempty"`

	// OrphanDependents determines if dependent objects should be orphaned
	OrphanDependents *bool `json:"orphanDependents,omitempty"`

	// Preconditions must be fulfilled before deletion
	Preconditions *Preconditions `json:"preconditions,omitempty"`
}

// Preconditions represents conditions that must be fulfilled before deletion
type Preconditions struct {
	// ResourceVersion specifies the target ResourceVersion
	ResourceVersion string `json:"resourceVersion,omitempty"`

	// UID specifies the target UID
	UID string `json:"uid,omitempty"`
}

// StatusDetails represents additional information about the status
type StatusDetails struct {
	// Name of the resource
	Name string `json:"name,omitempty"`

	// Group of the resource
	Group string `json:"group,omitempty"`

	// Kind of the resource
	Kind string `json:"kind,omitempty"`

	// Causes of the status
	Causes []StatusCause `json:"causes,omitempty"`

	// RetryAfterSeconds specifies the time to wait before retrying
	RetryAfterSeconds int `json:"retryAfterSeconds,omitempty"`

	// UID of the resource
	UID string `json:"uid,omitempty"`
}

// StatusCause represents the cause of the status
type StatusCause struct {
	// Field that caused the error
	Field string `json:"field,omitempty"`

	// Message describes the cause
	Message string `json:"message,omitempty"`

	// Reason provides machine-readable description
	Reason string `json:"reason,omitempty"`
}

// Delete removes a namespace with specified options
func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) (*DeleteResult, error) {
	url := client.ServiceURL("namespaces", opts.Name)

	b, err := build.RequestBody(opts.Body, "")
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
	// APIVersion of the response
	APIVersion string `json:"apiVersion,omitempty"`

	// Kind of the response
	Kind string `json:"kind,omitempty"`

	// Status of the operation: Success or Failure
	Status *NamespaceStatus `json:"status"`

	// Message provides human-readable description
	Message string `json:"message,omitempty"`

	// Code represents the HTTP status code
	Code int `json:"code,omitempty"`

	// Reason provides machine-readable description for Failure status
	Reason string `json:"reason,omitempty"`

	// Details provides additional information
	Details *StatusDetails `json:"details,omitempty"`

	// Metadata contains list metadata
	Metadata ListMeta `json:"metadata,omitempty"`
}
