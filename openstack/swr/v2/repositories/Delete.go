package repositories

import "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, organization, repository string) (r DeleteResult) {
	// DELETE /v2/manage/namespaces/{namespace}/repos/{repository}
	_, r.Err = client.Delete(client.ServiceURL("manage", "namespaces", organization, "repos", repository), nil)
	return
}
