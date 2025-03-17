package namespace

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get retrieves a specific namespace based on its name
func Get(client *golangsdk.ServiceClient, name string) (*Namespace, error) {
	// GET /apis/cci/v2/namespaces/{name}
	raw, err := client.Get(client.ServiceURL("namespaces", name), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
		JSONBody:    nil,
	})
	if err != nil {
		return nil, err
	}

	var res Namespace
	return &res, extract.Into(raw.Body, &res)
}
