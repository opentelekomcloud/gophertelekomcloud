package organizations

import "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, namespace string) (err error) {
	// DELETE /v2/manage/namespaces/{namespace}
	_, err = client.Delete(client.ServiceURL("manage", "namespaces", namespace), nil)
	return
}
