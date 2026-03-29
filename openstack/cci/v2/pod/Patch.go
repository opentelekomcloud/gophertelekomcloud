package pod

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// PatchOpts contains the options for partially updating a pod
type PatchOpts struct {
	// Namespace of the pod
	Namespace string `json:"-" required:"true"`

	// Name of the pod
	Name string `json:"-" required:"true"`

	// DryRun when present, indicates that modifications should not be persisted
	DryRun string `q:"dryRun,omitempty"`

	// FieldManager is a name associated with the actor or entity making these changes
	FieldManager string `q:"fieldManager,omitempty"`

	// Force indicates that the server should force Apply requests
	Force *bool `q:"force,omitempty"`

	// Pretty if 'true', then the output is pretty printed
	Pretty string `q:"pretty,omitempty"`
}

// Patch partially updates the specified pod.
// The body parameter should be the patch content matching the contentType.
// Supported content types:
//   - "application/json-patch+json" (JSON Patch, RFC 6902)
//   - "application/merge-patch+json" (Merge Patch, RFC 7386)
//   - "application/strategic-merge-patch+json" (Strategic Merge Patch)
func Patch(client *golangsdk.ServiceClient, opts PatchOpts, body interface{}, contentType string) (*Pod, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("namespaces", opts.Namespace, "pods", opts.Name).
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Patch(client.ServiceURL(url.String()), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
		MoreHeaders: map[string]string{
			"Content-Type": contentType,
		},
	})
	if err != nil {
		return nil, err
	}

	var res Pod
	return &res, extract.Into(raw.Body, &res)
}
