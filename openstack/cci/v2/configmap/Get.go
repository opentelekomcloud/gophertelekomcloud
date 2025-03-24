package configmap

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get retrieves a specific configmap based on its namespace and name
func Get(client *golangsdk.ServiceClient, namespace string, name string) (*ConfigMap, error) {
	// GET /apis/cci/v2/namespaces/{namespace}/configmaps/{name}
	raw, err := client.Get(client.ServiceURL("namespaces", namespace, "configmaps", name), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
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
