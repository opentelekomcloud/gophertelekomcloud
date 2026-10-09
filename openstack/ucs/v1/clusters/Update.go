package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	Kind       string          `json:"kind" required:"true"`
	APIVersion string          `json:"apiVersion" required:"true"`
	Metadata   *UpdateMetadata `json:"metadata,omitempty"`
	Spec       *UpdateSpec     `json:"spec,omitempty"`
}

type UpdateMetadata struct {
	Annotations map[string]string `json:"annotations,omitempty"`
}

type UpdateSpec struct {
	Country string `json:"country,omitempty"`
	City    string `json:"city,omitempty"`
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Put(client.ServiceURL("clusters", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
