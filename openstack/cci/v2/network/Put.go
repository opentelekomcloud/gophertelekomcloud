package network

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateNetworkOpts struct {
	Namespace  string         `json:"-"`
	Name       string         `json:"-"`
	APIVersion string         `json:"apiVersion,omitempty"`
	Kind       string         `json:"kind,omitempty"`
	Metadata   *ObjectMeta    `json:"metadata,omitempty"`
	Spec       *NetworkSpec   `json:"spec,omitempty"`
	Status     *NetworkStatus `json:"status,omitempty"`
}

func UpdateNetwork(client *golangsdk.ServiceClient, opts UpdateNetworkOpts) (*NetworkResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	var r NetworkResp
	_, err = client.Put(client.ServiceURL("namespaces", opts.Namespace, "networks", opts.Name), b, &r, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	if err != nil {
		return nil, err
	}

	return &r, nil
}
