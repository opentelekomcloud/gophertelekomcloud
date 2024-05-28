package v1_1

import (
	"os"
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	dataartsV11 "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/dataarts/v1.1"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/keypairs"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/servers"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dataarts/v1/script"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ecs/v1/cloudservers"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const keyPairName = "dataarts-test"
const ecsName = "dataarts-ecs-test"

func TestDataArtsJobExecution(t *testing.T) {
	// if os.Getenv("RUN_DATAART_LIFECYCLE") == "" {
	// 	t.Skip("too slow to run in zuul")
	// }

	workspace := os.Getenv("AWS_WORKSPACE")

	clientV1, err := clients.NewDataArtsV1Client()
	th.AssertNoErr(t, err)

	clientV11, err := clients.NewDataArtsV11Client()
	th.AssertNoErr(t, err)

	_ = dataartsV11.GetTestCluster(t, clientV11)
	// defer dataartsV11.DeleteCluster(t, clientV11, cluster.Id)

	kp, _ := getSSHKeys(t)
	// defer deleteSSHKeys(t, clientSSH)
	tools.PrintResource(t, kp)

	ec, _ := getECInstance(t)
	// defer openstack.DeleteCloudServer(t, clientEC, ec.ID)
	tools.PrintResource(t, ec)

	s := getScript(t, clientV1, workspace)
	tools.PrintResource(t, s)

	c := getConnection(t, clientV1, workspace)

	// client, err := clients.NewDataArtsV1Client()
	// th.AssertNoErr(t, err)
	//
	// workspace := ""
	//
	// t.Log("create a job")
	//
	// createOpts := &job.Job{
	// 	Name: jobName,
	// 	Schedule: job.Schedule{
	// 		Type: "EXECUTE_ONCE",
	// 	},
	// 	ProcessType: "BATCH",
	// }
	//
	// err = job.Create(client, *createOpts)
	// th.AssertNoErr(t, err)
	//
	// t.Log("schedule job cleanup")
	// t.Cleanup(func() {
	// 	jobCleanup(t, client)
	// })
}

func getSSHKeys(t *testing.T) (*keypairs.KeyPair, *golangsdk.ServiceClient) {
	t.Helper()

	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	kp, err := keypairs.Get(client, keyPairName).Extract()
	if kp != nil {
		return kp, client
	}
	opts := keypairs.CreateOpts{
		Name: keyPairName,
	}

	kp, err = keypairs.Create(client, opts).Extract()
	th.AssertNoErr(t, err)

	return kp, client
}

func deleteSSHKeys(t *testing.T, client *golangsdk.ServiceClient) {
	th.AssertNoErr(t, keypairs.Delete(client, keyPairName).ExtractErr())
}

func getECInstance(t *testing.T) (*cloudservers.CloudServer, *golangsdk.ServiceClient) {
	t.Helper()

	clientV1, err := clients.NewComputeV1Client()
	th.AssertNoErr(t, err)

	clientV2, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	listOpts := servers.ListOpts{
		Name: ecsName,
	}

	p, err := servers.List(clientV2, listOpts).AllPages()
	ss, err := servers.ExtractServers(p)

	for _, server := range ss {
		if server.Name == ecsName {
			ec, err := cloudservers.Get(clientV1, server.ID).Extract()
			th.AssertNoErr(t, err)
			return ec, clientV1
		}
	}

	createOpts := openstack.GetCloudServerCreateOpts(t)

	createOpts.Name = ecsName
	createOpts.KeyName = keyPairName

	// Create ECSv1 instance
	ecs := openstack.CreateCloudServer(t, clientV1, createOpts)

	return ecs, clientV1
}

func getScript(t *testing.T, client *golangsdk.ServiceClient, workspace string) *script.Script {
	t.Helper()

	s, err := script.Get(client, scriptName, workspace)
	if err == nil {
		return s
	}

	errNew, ok := err.(golangsdk.ErrDefault400)
	th.AssertEquals(t, true, ok)

	if errNew.Actual != 400 {
		th.AssertNoErr(t, err)
	}

	t.Log("create a script")

	createOpts := script.Script{
		Name:           scriptName,
		Type:           "Shell",
		Content:        "echo 123456",
		ConnectionName: "anyConnection",
	}

	err = script.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Logf("waiting for the script to be created, 3 seconds")
	time.Sleep(3 * time.Second)

	s, err = script.Get(client, scriptName, workspace)
	th.AssertNoErr(t, err)
	return s
}

func getConnection(t *testing.T, client *golangsdk.ServiceClient, workspace string) {
	t.Helper()

	s, err := connection..Get(client, scriptName, workspace)
	if err == nil {
		return s
	}

	errNew, ok := err.(golangsdk.ErrDefault400)
	th.AssertEquals(t, true, ok)

	if errNew.Actual != 400 {
		th.AssertNoErr(t, err)
	}
}
