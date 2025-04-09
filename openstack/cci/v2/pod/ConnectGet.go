package pod

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func ConnectGet(client *golangsdk.ServiceClient, namespace string, name string) error {
	// GET /apis/cci/v2/namespaces/{namespace}/pods/{name}/exec
	_, err := client.Get(client.ServiceURL("namespaces", namespace, "pods", name, "exec"), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}
