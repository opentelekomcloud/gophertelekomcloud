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

	createResp, err := cluster.Create(client, createOpts)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, createResp)

	th.AssertNoErr(t, waitForStateAvailable(client, 1200, createResp.Id))
	t.Cleanup(func() { deleteDataArts(t, client, createResp.Id) })

	getCluster, err := cluster.Get(client, createResp.Id)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, getCluster)
}

func waitForStateAvailable(client *golangsdk.ServiceClient, secs int, instanceID string) error {
	jobClient := *client
	jobClient.ResourceBase = jobClient.Endpoint

	return golangsdk.WaitFor(secs, func() (bool, error) {
		resp, err := cluster.Get(client, instanceID)
		if err != nil {
			return false, err
		}

		if resp.Status == "200" {
			return true, nil
		}

		return false, nil
	})
}

func deleteDataArts(t *testing.T, client *golangsdk.ServiceClient, instanceId string) {
	t.Logf("Attempting to delete DataArts instance: %s", instanceId)

	err := cluster.Delete(client, instanceId)
	th.AssertNoErr(t, err)

	t.Logf("DataArts instance deleted: %s", instanceId)
}
