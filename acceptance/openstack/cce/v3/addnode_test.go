package v3

import (
	"os"
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/nodes"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestExistingNodeLifecycle(t *testing.T) {
	vpcID := os.Getenv("OS_VPC_ID")
	subnetID := os.Getenv("OS_NETWORK_ID")
	clusterID := os.Getenv("OS_CLUSTER_ID")
	serverID := os.Getenv("OS_SERVER_ID")
	sshKey := os.Getenv("OS_KEYPAIR_NAME")
	if vpcID == "" || subnetID == "" || clusterID == "" || sshKey == "" {
		t.Skip("OS_VPC_ID, OS_NETWORK_ID, OS_KEYPAIR_NAME and OS_CLUSTER_ID are required for this test")
	}
	client, err := clients.NewCceV3Client()
	th.AssertNoErr(t, err)

	nodeList, err := nodes.List(client, clusterID, nodes.ListOpts{})
	th.AssertNoErr(t, err)
	numNodes := len(nodeList)

	addNodeOpts := nodes.AcceptOpts{
		ClusterID:  clusterID,
		ApiVersion: "v3",
		Kind:       "List",
		NodeList: []nodes.AddNode{
			{
				ServerID: serverID,
				Spec: nodes.ReinstallNodeSpec{
					OS: "EulerOS 2.9",
					Login: nodes.LoginSpec{
						SshKey: sshKey,
					},
				},
			},
		},
	}

	jobId, err := nodes.Accept(client, addNodeOpts)
	th.AssertNoErr(t, err)

	var nodeId string

	err = waitForJobCompletion(client, jobId, 600, &nodeId)
	th.AssertNoErr(t, err)

	nodeList, err = nodes.List(client, clusterID, nodes.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, numNodes+1, len(nodeList))

	// Remove the node
	removeOpts := nodes.RemoveNodesOpts{
		ClusterID:  clusterID,
		ApiVersion: "v3",
		Kind:       "RemoveNodesTask",
		Spec: nodes.RemoveNodesSpec{
			Login: nodes.LoginSpec{
				SshKey: sshKey,
			},
			Nodes: []nodes.NodeItem{
				{
					UID: nodeId,
				},
			},
		},
	}

	_, err = nodes.Remove(client, removeOpts)
	th.AssertNoErr(t, err)

	err = golangsdk.WaitFor(1800, func() (bool, error) {
		_, err := nodes.Get(client, clusterID, nodeId)
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}
		return false, nil
	})
	th.AssertNoErr(t, err)

	nodeList, err = nodes.List(client, clusterID, nodes.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, numNodes, len(nodeList))
}

func waitForJobCompletion(client *golangsdk.ServiceClient, jobId string, secs int, nodeId *string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		job, err := nodes.GetJobDetails(client, jobId)
		if err != nil {
			return false, err
		}

		for _, s := range job.Spec.SubJobs {
			if s.Status.Phase == "Success" {
				nodeId = &job.Metadata.ID
				return true, nil
			}
		}

		time.Sleep(10 * time.Second)

		return false, nil
	})
}
