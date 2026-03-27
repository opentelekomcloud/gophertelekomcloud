package pod

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// CreateOpts contains the options for creating a pod
type CreateOpts struct {
	// DryRun when present, indicates that modifications should not be persisted
	DryRun string `json:"-" q:"dryRun,omitempty"`

	// FieldManager is a name associated with the actor or entity making these changes
	FieldManager string `json:"-" q:"fieldManager,omitempty"`

	// Pretty if 'true', then the output is pretty printed
	Pretty string `json:"-" q:"pretty,omitempty"`

	Pod
}

// Create requests the creation of a new pod
func Create(client *golangsdk.ServiceClient, namespace string, opts CreateOpts) (*Pod, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("namespaces", namespace, "pods").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	var result Pod

	// POST /apis/cci/v2/namespaces/{namespace}/pods
	_, err = client.Post(client.ServiceURL(url.String()), b, &result, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return &result, err
}
