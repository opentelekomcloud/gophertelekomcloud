package organizations

import "github.com/opentelekomcloud/gophertelekomcloud"

func GetPermissions(client *golangsdk.ServiceClient, organization string) (r GetPermissionsResult) {
	_, r.Err = client.Get(client.ServiceURL("manage", "namespaces", organization, "access"), &r.Body, nil)
	return
}
