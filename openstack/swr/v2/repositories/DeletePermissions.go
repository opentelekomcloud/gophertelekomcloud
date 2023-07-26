package repositories

import "github.com/opentelekomcloud/gophertelekomcloud"

func DeletePermissions(client *golangsdk.ServiceClient, organization, repository string, userID string) (err error) {
	// DELETE /v2/manage/namespaces/{namespace}/repos/{repository}/access
	_, err = client.DeleteWithBody(client.ServiceURL("manage", "namespaces", organization, "repos", repository, "access"), []any{userID}, nil)
	return
}
