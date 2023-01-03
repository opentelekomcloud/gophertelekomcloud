package organizations

import "github.com/opentelekomcloud/gophertelekomcloud"

func UpdatePermissions(client *golangsdk.ServiceClient, organization string, opts Auth) (r UpdatePermissionsResult) {
	b, err := opts.ToPermissionUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	realBody := []interface{}{b}
	_, r.Err = client.Patch(client.ServiceURL("manage", "namespaces", organization, "access"), realBody, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	return
}
