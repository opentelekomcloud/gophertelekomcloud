package pod

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// Update requests the update of a new pod
func Update(client *golangsdk.ServiceClient, namespace, pod string, opts Pod) (*Pod, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	var result Pod

	// PUT /apis/cci/v2/namespaces/{namespace}/pods
	_, err = client.Put(client.ServiceURL("namespaces", namespace, "pods", pod), b, &result, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return &result, err
}
