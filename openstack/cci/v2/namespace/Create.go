package namespace

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// CreateOpts contains the options for creating a new namespace
type CreateOpts struct {
	// DryRun when present, indicates that modifications should not be persisted
	DryRun string `json:"-" q:"dryRun,omitempty"`

	// FieldManager is a name associated with the actor or entity making these changes
	FieldManager string `json:"-" q:"fieldManager,omitempty"`

	// Pretty if 'true', then the output is pretty printed
	Pretty string `json:"-" q:"pretty,omitempty"`

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
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Namespace, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("namespaces").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	var result Namespace

	// POST /apis/cci/v2/namespaces
	_, err = client.Post(client.ServiceURL(url.String()), b, &result, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return &result, err
}
