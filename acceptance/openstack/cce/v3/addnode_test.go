package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/clusters"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/nodes"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAddExistingNode(t *testing.T) {
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")
	clusterID := clients.EnvOS.GetEnv("CLUSTER_ID")
	serverID := clients.EnvOS.GetEnv("SERVER_ID")
	sshKey := clients.EnvOS.GetEnv("KEYPAIR_NAME")
	if vpcID == "" || subnetID == "" || clusterID == "" || sshKey == "" {
		t.Skip("OS_VPC_ID, OS_NETWORK_ID, OS_KEYPAIR_NAME and OS_CLUSTER_ID are required for this test")
	}
	client, err := clients.NewCceV3Client()
	th.AssertNoErr(t, err)

	nodeList, err := nodes.List(client, clusterID, nodes.ListOpts{})
	th.AssertNoErr(t, err)
	num_nodes := len(nodeList)

	addNodeOpts := clusters.AddExistingNodeOpts{
		APIVersion: "v3",
		Kind:       "List",
		NodeList: []clusters.AddNode{
			{
				ServerID: serverID,
				Spec: &clusters.ReinstallNodeSpec{
					OS: "EulerOS 2.9",
					Login: clusters.Login{
						SSHKey: sshKey,
					},
				},
			},
		},
	}

	_, err = clusters.AddExistingNode(client, clusterID, addNodeOpts)
	th.AssertNoErr(t, err)

	nodeList, err = nodes.List(client, clusterID, nodes.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, num_nodes+1, len(nodeList))
}
