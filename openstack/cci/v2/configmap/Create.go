package configmap

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// CreateOpts contains the parameters for creating a ConfigMap
type CreateOpts struct {
	// APIVersion defines the versioned schema of this representation of an object
	APIVersion string `json:"apiVersion,omitempty"`

	// BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'
	BinaryData map[string]string `json:"binaryData,omitempty"`

	// Data contains the configuration data
	Data map[string]string `json:"data,omitempty"`

	// Immutable, if set to true, ensures that data stored in the ConfigMap cannot be updated
	Immutable *bool `json:"immutable,omitempty"`

	// Kind is a string value representing the REST resource this object represents
	Kind string `json:"kind,omitempty"`

	// Metadata contains the object metadata
	Metadata ObjectMeta `json:"metadata,omitempty"`
}

// Create requests the creation of a new ConfigMap
func Create(client *golangsdk.ServiceClient, namespace string, opts CreateOpts) (*ConfigMap, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	var result ConfigMap

	_, err = client.Post(client.ServiceURL("namespaces", namespace, "configmaps"), b, &result, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})

	return &result, err
}
