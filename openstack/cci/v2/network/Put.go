package network

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	Namespace       string         `json:"-"`
	Name            string         `json:"-"`
	DryRun          string         `json:"-" q:"dryRun,omitempty"`
	FieldManager    string         `json:"-" q:"fieldManager,omitempty"`
	FieldValidation string         `json:"-" q:"fieldValidation,omitempty"`
	Pretty          string         `json:"-" q:"pretty,omitempty"`
	APIVersion      string         `json:"apiVersion,omitempty"`
	Kind            string         `json:"kind,omitempty"`
	Metadata        *ObjectMeta    `json:"metadata,omitempty"`
	Spec            *NetworkSpec   `json:"spec,omitempty"`
	Status          *NetworkStatus `json:"status,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*Network, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("namespaces", opts.Namespace, "networks", opts.Name).
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	var r Network
	_, err = client.Put(client.ServiceURL(url.String()), b, &r, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	if err != nil {
		return nil, err
	}

	return &r, nil
}
