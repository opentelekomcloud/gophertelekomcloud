package pod

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// Create requests the creation of a new pod
func Create(client *golangsdk.ServiceClient, namespace string, opts Pod) (*Pod, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	var result Pod

	// POST /apis/cci/v2/namespaces/{namespace}/pods
	_, err = client.Post(client.ServiceURL("namespaces", namespace, "pods"), b, &result, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return &result, err
}
