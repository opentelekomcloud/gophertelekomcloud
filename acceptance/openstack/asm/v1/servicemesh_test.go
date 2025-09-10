package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/asm/v1/servicemesh"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestASMServiceMeshLifecycle(t *testing.T) {
	cceClusterId := clients.EnvOS.GetEnv("CLUSTER_ID")
	cceNodeId := clients.EnvOS.GetEnv("NODE_ID")
	if cceClusterId == "" || cceNodeId == "" {
		t.Skip("OS_CLUSTER_ID or OS_NODE_ID is missing but is required for ASM lifecycle test")
	}

	// CREATE V1 CLIENT
	client, err := clients.NewASMV1Client()
	th.AssertNoErr(t, err)

	meshName := "test-acc-asm-mesh"
	createOpts := servicemesh.CreateOpts{
		APIVersion: "v1",
		Kind:       "mesh",
		Metadata: servicemesh.MeshMetadata{
			Name: meshName,
		},
		Spec: servicemesh.MeshSpec{
			Type:    "InCluster",
			Version: "1.18.7-r5",
			ExtendParams: &servicemesh.MeshExtendParams{
				Clusters: []servicemesh.MeshCluster{
					{
						ClusterID: cceClusterId,
						Installation: &servicemesh.InstallationConfig{
							Nodes: &servicemesh.Selector{
								FieldSelector: &servicemesh.FieldSelector{
									Key:      "UID",
									Operator: "In",
									Values:   []string{cceNodeId},
								},
							},
						},
					},
				},
			},
		},
	}

	serviceMesh, err := servicemesh.Create(client, createOpts)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, servicemesh.Delete(client, serviceMesh.Metadata.UID))
	})

	meshList, err := servicemesh.List(client)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(meshList))

	meshGet, err := servicemesh.Get(client, serviceMesh.Metadata.UID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, meshName, meshGet.Metadata.Name)
}
