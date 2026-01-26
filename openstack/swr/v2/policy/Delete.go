package policy

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// This function is used to delete an image retention policy.
func Delete(client *golangsdk.ServiceClient, organization, repository, policy string) (err error) {
	// DELETE /v2/manage/namespaces/{namespace}/repos/{repository}/retentions/{retention_id}
	_, err = client.Delete(client.ServiceURL("manage", "namespaces", organization, "repos", repository, "retentions", policy), nil)
	return err
}
