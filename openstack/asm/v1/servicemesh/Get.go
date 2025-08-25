package servicemesh

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to obtain details about a service mesh.
func Get(client *golangsdk.ServiceClient, meshId string) (*ServiceMesh, error) {
	// GET /v1/{project_id}/meshes/{mesh_id}
	raw, err := client.Get(client.ServiceURL("meshes", meshId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ServiceMesh
	err = extract.Into(raw.Body, &res)
	return &res, err
}
