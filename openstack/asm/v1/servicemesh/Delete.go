package servicemesh

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// This function is used to delete a service mesh.
func Delete(client *golangsdk.ServiceClient, meshId string) error {
	// DELETE /v1/{project_id}/meshes/{mesh_id}
	_, err := client.Delete(client.ServiceURL("meshes", meshId), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
