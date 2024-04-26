package v1_1

import (
	"os"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dataarts/v1.1/cluster"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDataArtsClusterLifecycle(t *testing.T) {
	if os.Getenv("RUN_DATAART_LIFECYCLE") == "" {
		t.Skip("too slow to run in zuul")
	}

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")
	secGroupId := clients.EnvOS.GetEnv("SECURITY_GROUP_ID")

	client, err := clients.NewDataArtsV11Client()
	th.AssertNoErr(t, err)

	dataStore := cluster.Datastore{
		Type:    "cdm",
		Version: "2.9.1.200",
	}
	instance := cluster.Instance{
		AZ:        "eu-de-01",
		FlavorRef: "a79fd5ae-1833-448a-88e8-3ea2b913e1f6",
		Type:      "cdm",
		Nics: []cluster.Nic{
			{
				SecurityGroupId: secGroupId,
				NetId:           subnetID,
			},
		},
	}

	createOpts := cluster.CreateOpts{
		Cluster: cluster.Cluster{
			// setting this parameter to true results in 400 error
			IsScheduleBootOff: pointerto.Bool(false),
			VpcId:             vpcID,
			Name:              tools.RandomString("test-dataarts", 5),
			DataStore:         &dataStore,
			Instances: []cluster.Instance{
				instance,
			},
		},
	}

	t.Log("starting to create a DataArts cluster")
	createResp, err := cluster.Create(client, createOpts, "en")
	th.AssertNoErr(t, err)
	tools.PrintResource(t, createResp)

	t.Log("check cluster status, should be normal")
	th.AssertNoErr(t, waitForState(client, 300, createResp.Id, "200"))
	t.Cleanup(func() { deleteDataArts(t, client, createResp.Id) })

	t.Log("get cluster details")
	getCluster, err := cluster.Get(client, createResp.Id)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, getCluster)

	t.Log("stop cluster")
	_, err = cluster.Stop(client, getCluster.Id, cluster.StopOpts{})
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, waitForState(client, 300, createResp.Id, "900"))

	t.Log("start cluster")
	_, err = cluster.Start(client, getCluster.Id, cluster.StartOpts{})
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, waitForState(client, 300, createResp.Id, "200"))
}

func waitForState(client *golangsdk.ServiceClient, secs int, instanceID string, status string) error {
	jobClient := *client
	jobClient.ResourceBase = jobClient.Endpoint

	return golangsdk.WaitFor(secs, func() (bool, error) {
		resp, err := cluster.Get(client, instanceID)
		if err != nil {
			return false, err
		}

		if resp.Status == status {
			return true, nil
		}

		return false, nil
	})
}

func deleteDataArts(t *testing.T, client *golangsdk.ServiceClient, instanceId string) {
	t.Logf("Attempting to delete DataArts instance: %s", instanceId)

	jobId, err := cluster.Delete(client, instanceId, nil)
	t.Logf(jobId.JobId)
	th.AssertNoErr(t, err)

	t.Logf("DataArts instance deleted: %s", instanceId)
}
