package pod

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func ConnectPost(client *golangsdk.ServiceClient, namespace string, name string) error {
	// POST /apis/cci/v2/namespaces/{namespace}/pods/{name}/exec
	_, err := client.Post(client.ServiceURL("namespaces", namespace, "pods", name, "exec"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}
