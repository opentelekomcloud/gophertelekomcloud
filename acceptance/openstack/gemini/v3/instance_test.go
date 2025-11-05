package v3

import (
	"os"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/gemini/v3/instance"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/gemini/v3/job"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func createGeminiInstance(t *testing.T, client *golangsdk.ServiceClient, vpcID, subnetID, secGroupID string) *instance.CreateResp {
	t.Logf("Attempting to create gemini db")

	geminiName := tools.RandomString("tf-gemini", 3)

	opts := instance.CreateOpts{
		Name: geminiName,
		DataStore: instance.DataStoreOpt{
			Type:          "cassandra",
			Version:       "3.11",
			StorageEngine: "rocksDB",
		},
		Region:           "eu-de",
		AvailabilityZone: "eu-de-01,eu-de-02,eu-de-03",
		VpcId:            vpcID,
		SubnetId:         subnetID,
		SecurityGroupId:  secGroupID,
		Password:         "Root1234#",
		Mode:             "Cluster",
		Flavor: []instance.FlavorOpt{
			{
				Num:      "3",
				Size:     "500",
				Storage:  "ULTRAHIGH",
				SpecCode: "geminidb.cassandra.xlarge.8",
			},
		},
	}

	createResp, err := instance.Create(client, opts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createResp.Name, geminiName)
	th.AssertEquals(t, createResp.AvailabilityZone, opts.AvailabilityZone)

	th.AssertNoErr(t, job.WaitForJobCompletion(client, 1200, createResp.JobId))

	th.AssertNoErr(t, waitForInstanceAvailable(client, 1200, createResp.Id))

	return createResp
}

func TestGeminiInstanceLifecycle(t *testing.T) {
	t.Skip("too long to run in ci")
	vpcID := os.Getenv("OS_VPC_ID")
	subnetID := os.Getenv("OS_NETWORK_ID")
	secGroupID := os.Getenv("OS_SECURITY_GROUP_ID")

	client, err := clients.NewGeminiDBClient()
	th.AssertNoErr(t, err)

	createResp := createGeminiInstance(t, client, vpcID, subnetID, secGroupID)

	t.Cleanup(func() {
		t.Logf("Attempting to delete gemini db")
		_, err = instance.Delete(client, createResp.Id)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to list gemini db instances")

	listResp, err := instance.ListGeminiDB(client, instance.ListGeminiDBOpts{Id: createResp.Id})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(listResp), 1)

	var foundInstance *instance.ListResult
	for i := range listResp {
		if listResp[i].Id == createResp.Id {
			foundInstance = &listResp[i]
			break
		}
	}

	th.AssertEquals(t, foundInstance.Name, createResp.Name)
	th.AssertEquals(t, foundInstance.VpcId, createResp.VpcId)
	th.AssertEquals(t, foundInstance.SubnetId, createResp.SubnetId)
	th.AssertEquals(t, foundInstance.SecurityGroupId, createResp.SecurityGroupId)

	updatedName := createResp.Name + "-updated"

	t.Logf("Attempting to update gemini db instance name")

	err = instance.RenameInstance(client, instance.RenameInstanceOpts{
		InstanceID: createResp.Id,
		Name:       updatedName},
	)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to reset gemini db instance password")

	err = instance.ResetPassword(client, instance.ResetPasswordOpts{
		InstanceId: createResp.Id,
		Password:   "Root1234#New"},
	)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to change gemini db instance flavor")
	resizeFlavorJob, err := instance.ResizeInstance(client, instance.ResizeInstanceOpts{
		InstanceID:     createResp.Id,
		TargetSpecCode: "geminidb.cassandra.2xlarge.8"},
	)
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, job.WaitForJobCompletion(client, 2000, resizeFlavorJob))
}

func TestGeminiInstanceAdvancedOperations(t *testing.T) {
	t.Skip("too long to run in ci")
	vpcID := os.Getenv("OS_VPC_ID")
	subnetID := os.Getenv("OS_NETWORK_ID")
	secGroupID := os.Getenv("OS_SECURITY_GROUP_ID")
	altSecGroupID := os.Getenv("OS_ALT_SECURITY_GROUP_ID")

	client, err := clients.NewGeminiDBClient()
	th.AssertNoErr(t, err)

	createResp := createGeminiInstance(t, client, vpcID, subnetID, secGroupID)

	t.Cleanup(func() {
		t.Logf("Attempting to delete gemini db")
		_, err = instance.Delete(client, createResp.Id)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to change security group")
	changeSecGroupResp, err := instance.ChangeSecGroup(client, instance.ChangeSecGroupOpts{
		InstanceId:      createResp.Id,
		SecurityGroupId: altSecGroupID,
	})
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, job.WaitForJobCompletion(client, 600, changeSecGroupResp.JobId))

	t.Logf("Attempting to extend volume")
	extendVolumeResp, err := instance.ExtendVolume(client, instance.ExtendVolumeOpts{
		InstanceId: createResp.Id,
		Size:       600,
	})
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, job.WaitForJobCompletion(client, 600, extendVolumeResp.JobId))

	t.Logf("Attempting to configure disk auto expansion")
	err = instance.ConfigureAutoExpansion(client, instance.DiskAutoExpansionOpts{
		InstanceIds:  []string{createResp.Id},
		SwitchOption: "on",
		Policy: &instance.DiskAutoExpansionPolicy{
			Threshold: 90,
			Step:      10,
			Size:      1000,
		},
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to get auto expansion settings")
	autoExpResp, err := instance.GetAutoExpansion(client, createResp.Id)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, autoExpResp.Policy.Threshold, 90)
	th.AssertEquals(t, autoExpResp.Policy.Step, 10)

	th.AssertNoErr(t, waitForInstanceAvailable(client, 1200, createResp.Id))
}

func TestGeminiInstanceNodeOperations(t *testing.T) {
	t.Skip("too long to run in ci")
	vpcID := os.Getenv("OS_VPC_ID")
	subnetID := os.Getenv("OS_NETWORK_ID")
	secGroupID := os.Getenv("OS_SECURITY_GROUP_ID")

	client, err := clients.NewGeminiDBClient()
	th.AssertNoErr(t, err)

	createResp := createGeminiInstance(t, client, vpcID, subnetID, secGroupID)

	t.Cleanup(func() {
		t.Logf("Attempting to delete gemini db")
		_, err = instance.Delete(client, createResp.Id)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to enlarge node count")
	enlargeNodeResp, err := instance.EnlargeNode(client, instance.EnlargeNodeOpts{
		InstanceId: createResp.Id,
		Num:        1,
		SubnetId:   subnetID,
	})
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, job.WaitForJobCompletion(client, 1200, enlargeNodeResp.JobId))

	t.Logf("Attempting to reduce node count")
	numToReduce := 1
	reduceNodeResp, err := instance.ReduceNode(client, instance.ReduceNodeOpts{
		InstanceID: createResp.Id,
		Num:        &numToReduce,
	})
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, job.WaitForJobCompletion(client, 1200, reduceNodeResp.JobId))
}

func TestGeminiListDatastores(t *testing.T) {
	client, err := clients.NewGeminiDBClient()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to list Gemini db")
	listResp, err := instance.ListGeminiDB(client, instance.ListGeminiDBOpts{})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, listResp)
}

func waitForInstanceAvailable(client *golangsdk.ServiceClient, secs int, instanceID string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		gemInstances, err := instance.ListGeminiDB(client, instance.ListGeminiDBOpts{
			Id: instanceID,
		})
		if err != nil {
			return false, err
		}
		if gemInstances[0].Status == "normal" && len(gemInstances[0].Actions) == 0 {
			return true, nil
		}
		return false, nil
	})
}
