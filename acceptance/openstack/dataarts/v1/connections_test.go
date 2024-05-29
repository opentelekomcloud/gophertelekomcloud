package v1_1

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	dataartsV11 "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/dataarts/v1.1"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/keypairs"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/servers"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dataarts/v1.1/cluster"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dataarts/v1/connection"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ecs/v1/cloudservers"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/obs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const connectionName = "testConnection"

const keyPairName = "dataarts-test"
const ecsName = "dataarts-ecs-test"

func TestDataArtsConnectionsLifecycle(t *testing.T) {
	if os.Getenv("RUN_DATAART_LIFECYCLE") == "" {
		t.Skip("too slow to run in zuul")
	}

	workspace := os.Getenv("AWS_WORKSPACE")
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	kms := clients.EnvOS.GetEnv("KMS_ID")

	clientV1, err := clients.NewDataArtsV1Client()
	th.AssertNoErr(t, err)

	ec, c := prepareEnv(t)

	t.Log("create a connection")

	createOpts := connection.Connection{
		Name: connectionName,
		Type: "HOST",
		Config: connection.HOSTConfig{
			IP:          ec.Addresses[vpcID][0].Addr,
			Port:        "22",
			Username:    "linux",
			AgentName:   c.Name,
			KMSKey:      kms,
			KeyLocation: fmt.Sprintf("obs://%s/%s.pem", bucketName, keyPairName),
		},
	}

	err = connection.Create(clientV1, createOpts)
	th.AssertNoErr(t, err)

	t.Log("schedule connection cleanup")
	t.Cleanup(func() {
		t.Logf("attempting to delete connection: %s", connectionName)
		err := connection.Delete(clientV1, connectionName, workspace)
		th.AssertNoErr(t, err)
		t.Logf("connection is deleted: %s", connectionName)
	})

	t.Log("should wait 5 seconds")
	time.Sleep(5 * time.Second)
	t.Log("get connection")

	storedConnection, err := connection.Get(clientV1, connectionName, workspace)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, storedConnection)

	t.Log("modify connection")

	storedConnection.Description = "newDescription"

	err = connection.Update(clientV1, *storedConnection, connection.UpdateOpts{}, workspace)
	th.AssertNoErr(t, err)

	t.Log("get connection")

	storedConnection, err = connection.Get(clientV1, connectionName, workspace)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, storedConnection)
	th.CheckEquals(t, "newDescription", storedConnection.Description)
}

func prepareEnv(t *testing.T) (*cloudservers.CloudServer, *cluster.ClusterQuery) {
	t.Helper()

	clientV11, err := clients.NewDataArtsV11Client()
	th.AssertNoErr(t, err)

	clientOBS, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

	clientSSH, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	clientComputeV1, err := clients.NewComputeV1Client()
	th.AssertNoErr(t, err)

	var wg sync.WaitGroup
	var c *cluster.ClusterQuery
	var kp *keypairs.KeyPair
	var ec *cloudservers.CloudServer

	kpChan := make(chan *keypairs.KeyPair)

	wg.Add(3)

	go func() {
		defer wg.Done()
		c = dataartsV11.GetTestCluster(t, clientV11)
		t.Cleanup(func() {
			t.Log("clean up test cluster")
			dataartsV11.DeleteCluster(t, clientV11, c.Id)
		})
	}()

	go func() {
		defer wg.Done()

		kp = getSSHKeys(t, clientSSH)
		kpChan <- kp

		prepareTestBucket(t, clientOBS)
		uploadSSHKey(t, clientOBS, kp)
		t.Cleanup(func() {
			t.Log("clean up test key pair")
			_ = keypairs.Delete(clientSSH, keyPairName).ExtractErr()
		})
		t.Cleanup(func() {
			t.Log("clean up test bucket")
			defer cleanupBucket(t, clientOBS)
		})
	}()

	go func() {
		defer wg.Done()
		ec = getECInstance(t, clientComputeV1, kpChan)
		t.Cleanup(func() {
			defer openstack.DeleteCloudServer(t, clientComputeV1, ec.ID)
		})
	}()

	wg.Wait()

	return ec, c
}

func getSSHKeys(t *testing.T, client *golangsdk.ServiceClient) *keypairs.KeyPair {
	t.Helper()

	t.Log("trying to get a ssh key pair")
	kp, _ := keypairs.Get(client, keyPairName).Extract()
	if kp != nil {
		return kp
	}

	t.Log("key pair not found, create one")
	opts := keypairs.CreateOpts{
		Name: keyPairName,
	}

	kp, err := keypairs.Create(client, opts).Extract()
	th.AssertNoErr(t, err)

	return kp
}

func uploadSSHKey(t *testing.T, client *obs.ObsClient, kp *keypairs.KeyPair) {
	t.Helper()

	f := fmt.Sprintf("%s.pem", keyPairName)
	t.Logf("upload ssh key %s to obs bucket: %s", f, bucketName)
	uploadFile(t, client, f, strings.NewReader(kp.PrivateKey))
}

func getECInstance(t *testing.T, clientV1 *golangsdk.ServiceClient, kpCh chan *keypairs.KeyPair) *cloudservers.CloudServer {
	t.Helper()
	t.Log("trying to get ec instance")

	<-kpCh

	clientV2, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	listOpts := servers.ListOpts{
		Name: ecsName,
	}

	p, err := servers.List(clientV2, listOpts).AllPages()
	th.AssertNoErr(t, err)

	ss, err := servers.ExtractServers(p)
	th.AssertNoErr(t, err)

	for _, server := range ss {
		if server.Name == ecsName {
			ec, err := cloudservers.Get(clientV1, server.ID).Extract()
			th.AssertNoErr(t, err)
			return ec
		}
	}

	t.Log("ec instance not found, create one")

	createOpts := openstack.GetCloudServerCreateOpts(t)

	createOpts.Name = ecsName
	createOpts.KeyName = keyPairName

	ecs := openstack.CreateCloudServer(t, clientV1, createOpts)

	return ecs
}
