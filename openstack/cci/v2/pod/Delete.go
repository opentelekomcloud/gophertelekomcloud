package pod

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type DeleteOpts struct {
	// Name of the namespace to delete
	NameSpace string `json:"-" required:"true"`

	// Name of the pod to delete
	PodName string `json:"-" required:"true"`

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

// Delete removes a pod with specified options
func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) (*Pod, error) {
	b, err := build.RequestBody(opts.Body, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.DeleteWithBody(client.ServiceURL("namespaces", opts.NameSpace, "pods", opts.PodName), b, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})
	if err != nil {
		return nil, err
	}

	var res Pod
	err = extract.Into(raw.Body, &res)
	return &res, err
}
