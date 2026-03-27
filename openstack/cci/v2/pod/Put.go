package pod

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UpdateOpts contains the options for updating a pod
type UpdateOpts struct {
	// DryRun when present, indicates that modifications should not be persisted
	DryRun string `json:"-" q:"dryRun,omitempty"`

	// FieldManager is a name associated with the actor or entity making these changes
	FieldManager string `json:"-" q:"fieldManager,omitempty"`

	// Pretty if 'true', then the output is pretty printed
	Pretty string `json:"-" q:"pretty,omitempty"`

	Pod
}

// Update requests the update of a pod
func Update(client *golangsdk.ServiceClient, namespace, podName string, opts UpdateOpts) (*Pod, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("namespaces", namespace, "pods", podName).
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	var result Pod

	// PUT /apis/cci/v2/namespaces/{namespace}/pods/{name}
	_, err = client.Put(client.ServiceURL(url.String()), b, &result, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return &result, err
}
