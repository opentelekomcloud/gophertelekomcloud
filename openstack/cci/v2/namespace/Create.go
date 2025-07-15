package namespace

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// CreateOpts contains the options for creating a new namespace
type CreateOpts struct {
	// APIVersion defines the versioned schema of this representation of an object
	APIVersion string `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents
	Kind string `json:"kind,omitempty"`

	// Metadata contains the namespace metadata
	Metadata Metadata `json:"metadata" required:"true"`

	// Spec defines the behavior of the Namespace
	Spec *NamespaceSpec `json:"spec,omitempty"`

	// Status describes the current status of a Namespace
	Status *NamespaceStatus `json:"status,omitempty"`
}

// Create requests the creation of a new namespace
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*CreateOpts, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	var result CreateOpts

	// POST /apis/cci/v2/namespaces
	_, err = client.Post(client.ServiceURL("namespaces"), b, &result, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return &result, err
}
