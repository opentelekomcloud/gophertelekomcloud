package configmap

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// UpdateOpts contains all the values needed to update a ConfigMap
type UpdateOpts struct {
	// APIVersion defines the versioned schema of this representation of an object
	APIVersion string `json:"apiVersion,omitempty"`

	// BinaryData contains the binary data
	BinaryData map[string]string `json:"binaryData,omitempty"`

	// Data contains the configuration data
	Data map[string]string `json:"data,omitempty"`

	// Immutable if set to true, ensures that data stored in the ConfigMap cannot be updated
	Immutable *bool `json:"immutable,omitempty"`

	// Kind is a string value representing the REST resource this object represents
	Kind string `json:"kind,omitempty"`

	// Metadata contains the object metadata
	Metadata *ObjectMeta `json:"metadata,omitempty"`
}

// Update modifies the specified ConfigMap
func Update(client *golangsdk.ServiceClient, namespace string, name string, opts UpdateOpts) (*ConfigMap, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /apis/cci/v2/namespaces/{namespace}/configmaps/{name}
	raw, err := client.Put(client.ServiceURL("namespaces", namespace, "configmaps", name), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})
	if err != nil {
		return nil, err
	}

	var res ConfigMap
	return &res, extract.Into(raw.Body, &res)
}
