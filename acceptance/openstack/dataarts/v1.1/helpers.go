package v1_1

import (
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dataarts/v1.1/cluster"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const (
	dataStoreType    = "cdm"
	dataStoreVersion = "2.10.0.100"
)

const (
	instanceAZ     = "eu-de-01"
	instanceFlavor = "5ddb1071-c5d7-40e0-a874-8a032e81a697"
	instanceType   = "cdm"
)

const clusterTestName = "testAllCases"

func GetTestCluster(t *testing.T, client *golangsdk.ServiceClient) *cluster.ClusterQuery {
	t.Log("check if test cluster is created")

	clusters, err := cluster.List(client)
	th.AssertNoErr(t, err)

	for _, c := range clusters {
		if strings.HasPrefix(c.Name, clusterTestName) {
			return c
		}
	}

	t.Log("create a test cluster")
	c, err := createCluster(t, client)
	th.AssertNoErr(t, err)

	return c
}

func createCluster(t *testing.T, client *golangsdk.ServiceClient) (*cluster.ClusterQuery, error) {

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("SUBNET_ID")
	secGroupId := clients.EnvOS.GetEnv("SECURITY_GROUP_ID")

	dataStore := cluster.Datastore{
		Type:    dataStoreType,
		Version: dataStoreVersion,
	}
	instance := cluster.Instance{
		AZ:        instanceAZ,
		FlavorRef: instanceFlavor,
		Type:      instanceType,
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
			Name:              tools.RandomString(clusterTestName, 5),
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

	t.Log("check cluster status, should be normal")
	th.AssertNoErr(t, waitForState(client, 1200, createResp.Id, "200"))

	t.Log("get cluster details")
	getCluster, err := cluster.Get(client, createResp.Id)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, getCluster)

	return getCluster, err
}

func DeleteCluster(t *testing.T, client *golangsdk.ServiceClient, clusterId string) {
	t.Logf("Attempting to delete DataArts instance: %s", clusterId)

	jobId, err := cluster.Delete(client, clusterId, cluster.DeleteOpts{})
	th.AssertNoErr(t, err)

	t.Logf("DataArts instance deleted: %s, jobId: %s", clusterId, jobId.JobId)
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
