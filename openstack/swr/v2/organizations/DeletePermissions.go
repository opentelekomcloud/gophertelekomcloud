package organizations

import "github.com/opentelekomcloud/gophertelekomcloud"

func DeletePermissions(client *golangsdk.ServiceClient, organization string, userID string) (err error) {
	// DELETE /v2/manage/namespaces/{namespace}/access
	_, err = client.DeleteWithBody(client.ServiceURL("manage", "namespaces", organization, "access"), []interface{}{userID}, nil)
	return
}
