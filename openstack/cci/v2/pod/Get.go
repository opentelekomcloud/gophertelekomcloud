package pod

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get retrieves a specific pod based on its name
func Get(client *golangsdk.ServiceClient, nameSpace, pod string) (*Pod, error) {
	// GET /apis/cci/v2/namespaces/{namespace}/pods/{name}
	raw, err := client.Get(client.ServiceURL("namespaces", nameSpace, "pods", pod), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
		JSONBody:    nil,
	})
	if err != nil {
		return nil, err
	}

	var res Pod
	return &res, extract.Into(raw.Body, &res)
}
