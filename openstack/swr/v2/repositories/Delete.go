package repositories

import "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, organization, repository string) (err error) {
	// DELETE /v2/manage/namespaces/{namespace}/repos/{repository}
	_, err = client.Delete(client.ServiceURL("manage", "namespaces", organization, "repos", repository), nil)
	return
}
