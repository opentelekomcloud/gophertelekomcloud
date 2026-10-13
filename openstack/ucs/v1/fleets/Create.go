package fleets

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	Metadata CreateMetadata `json:"metadata" required:"true"`
	Spec     *CreateSpec    `json:"spec,omitempty"`
}

type CreateMetadata struct {
	Name string `json:"name" required:"true"`
}

type CreateSpec struct {
	ClusterIDs  []string `json:"clusterIds,omitempty"`
	Description string   `json:"description,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	raw, err := client.Post(client.ServiceURL("clustergroups"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		UID string `json:"uid"`
	}
	return res.UID, extract.Into(raw.Body, &res)
}
