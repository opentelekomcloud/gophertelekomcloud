package v1_1

import (
	"os"
	"testing"

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
	subnetID := clients.EnvOS.GetEnv("SUBNET_ID")
	secGroupId := clients.EnvOS.GetEnv("SECURITY_GROUP_ID")

	client, err := clients.NewDataArtsV11Client()
	th.AssertNoErr(t, err)

	dataStore := cluster.Datastore{
		Type:    "cdm",
		Version: "2.10.0.100",
	}
	instance := cluster.Instance{
		AZ:        "eu-de-01",
		FlavorRef: "5ddb1071-c5d7-40e0-a874-8a032e81a697",
		Type:      "cdm",
		Nics: []cluster.Nic{
			{
				SecurityGroupId: secGroupId,
				NetId:           subnetID,
			},
		},
	}

	createOpts := cluster.CreateOpts{
		XLang: "en",
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
	createResp, err := cluster.Create(client, createOpts)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, createResp)
	t.Log("DataArts cluster was created")

	t.Log("schedule clusters cleanup")
	t.Cleanup(func() { DeleteCluster(t, client, createResp.Id) })

	t.Log("check cluster status, should be normal")
	th.AssertNoErr(t, waitForState(client, 1200, createResp.Id, "200"))

	t.Log("get cluster details")
	getCluster, err := cluster.Get(client, createResp.Id)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, getCluster)
}
