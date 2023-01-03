package organizations

import "github.com/opentelekomcloud/gophertelekomcloud"

func DeletePermissions(client *golangsdk.ServiceClient, organization string, userID string) (r DeletePermissionsResult) {
	_, r.Err = client.Request("DELETE", client.ServiceURL("manage", "namespaces", organization, "access"), &golangsdk.RequestOpts{
		JSONBody: []interface{}{userID},
	})
	return
}
