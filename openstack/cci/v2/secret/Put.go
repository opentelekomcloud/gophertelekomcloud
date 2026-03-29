package secret

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	DryRun       string            `json:"-" q:"dryRun,omitempty"`
	FieldManager string            `json:"-" q:"fieldManager,omitempty"`
	Pretty       string            `json:"-" q:"pretty,omitempty"`
	APIVersion   string            `json:"apiVersion,omitempty"`
	Data         map[string]string `json:"data,omitempty"`
	Immutable    *bool             `json:"immutable,omitempty"`
	Kind         string            `json:"kind,omitempty"`
	Metadata     *ObjectMeta       `json:"metadata,omitempty"`
	StringData   map[string]string `json:"stringData,omitempty"`
	Type         string            `json:"type,omitempty"`
}

func Update(client *golangsdk.ServiceClient, namespace, name string, opts UpdateOpts) (*Secret, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("namespaces", namespace, "secrets", name).
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	var result Secret

	_, err = client.Put(client.ServiceURL(url.String()), b, &result, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})

	return &result, err
}
