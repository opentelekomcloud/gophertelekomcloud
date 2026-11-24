package endpoints

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// UpdateOpts contains the parameters of the Update function.
type UpdateOpts struct {
	// Endpoint name.
	// It can contain 1 to 64 characters. Only letters, digits, underscores (_), hyphens (-), and periods (.) are allowed.
	Name string `json:"name" required:"true"`
}

// Update is used to modify an endpoint.
// Send PUT /v2.1/endpoints/{endpoint_id}
func Update(client *golangsdk.ServiceClient, endpointID string, opts UpdateOpts) (*Endpoint, error) {
	if endpointID == "" {
		return nil, golangsdk.ErrMissingInput{Argument: "endpointID"}
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(v21URL(client, "endpoints", endpointID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res Endpoint
	err = extract.IntoStructPtr(raw.Body, &res, "endpoint")
	return &res, err
}
