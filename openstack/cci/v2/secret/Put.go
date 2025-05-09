package secret

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateSecretOpts struct {
	Namespace  string            `json:"-"`
	APIVersion string            `json:"apiVersion,omitempty"`
	Data       map[string]string `json:"data,omitempty"`
	Immutable  bool              `json:"immutable,omitempty"`
	Kind       string            `json:"kind,omitempty"`
	Metadata   *ObjectMeta       `json:"metadata,omitempty"`
	StringData map[string]string `json:"stringData,omitempty"`
	Type       string            `json:"type,omitempty"`
}

func Update(client *golangsdk.ServiceClient, namespace, name string, opts UpdateSecretOpts) (*SecretResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	var result SecretResp

	_, err = client.Put(client.ServiceURL("namespaces", namespace, "secrets", name), b, &result, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})

	return &result, err
}
