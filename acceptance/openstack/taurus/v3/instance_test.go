package v3

import (
	"os"
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/taurus/v3/instance"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/taurus/v3/job"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func createTaurusInstance(t *testing.T, client *golangsdk.ServiceClient, vpcID, subnetID, secGroupID string) *instance.CreateResponse {
	t.Logf("Attempting to create taurus db")

	taurusName := tools.RandomString("tf-taurus", 3)

	opts := instance.CreateOpts{
		Name:            taurusName,
		Region:          "eu-de",
		Mode:            "Cluster",
		Flavor:          "gaussdb.mysql.xlarge.x86.8",
		VpcId:           vpcID,
		SubnetId:        subnetID,
		SecurityGroupId: secGroupID,
		ConfigurationId: "43570e0de32e40c5a15f831aa5ce4176pr07",
		Password:        "Root1234#",
		AZMode:          "multi",
		MasterAZ:        "eu-de-01",
		SlaveCount:      1,
		DataStore: instance.DataStoreOpt{
			Type:    "gaussdb-mysql",
			Version: "8.0",
		},
		BackupStrategy: &instance.BackupStrategyOpt{
			StartTime: "08:00-09:00",
			KeepDays:  "1",
		},
		ChargeInfo: &instance.ChargeInfo{
			ChargingMode: "postPaid",
		},
	}

	createResp, err := instance.Create(client, opts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createResp.Instance.Name, taurusName)
	th.AssertEquals(t, createResp.Instance.MasterAZ, opts.MasterAZ)
	th.AssertEquals(t, createResp.Instance.AZMode, opts.AZMode)

	th.AssertNoErr(t, job.WaitForJobSuccess(client, 600, createResp.JobId))

	return createResp
}

func TestTaurusInstanceLifecycle(t *testing.T) {
	t.Skip("too long to run in ci")
	vpcID := os.Getenv("OS_VPC_ID")
	subnetID := os.Getenv("OS_NETWORK_ID")
	secGroupID := os.Getenv("OS_SECURITY_GROUP_ID")

	client, err := clients.NewTaurusDBV3Client()
	th.AssertNoErr(t, err)

	createResp := createTaurusInstance(t, client, vpcID, subnetID, secGroupID)

	t.Cleanup(func() {
		t.Logf("Attempting to delete taurus db")
		_, err = instance.Delete(client, createResp.Instance.Id)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to list taurus db instances")

	listResp, err := instance.List(client, instance.ListOpts{Id: createResp.Instance.Id})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(listResp), 1)

	t.Logf("Attempting to query taurus db instance")

	getResp, err := instance.Get(client, createResp.Instance.Id)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, listResp[0].Name, getResp.Name)
	th.AssertEquals(t, listResp[0].Status, getResp.Status)
	th.AssertEquals(t, listResp[0].Type, getResp.Type)
	th.AssertEquals(t, listResp[0].Port, getResp.Port)

	updatedName := getResp.Name + "-updated"

	t.Logf("Attempting to update taurus db instance name")

	updateNameJob, err := instance.UpdateName(client, createResp.Instance.Id, updatedName)
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, job.WaitForJobSuccess(client, 600, *updateNameJob))

	t.Logf("Attempting to reset taurus db instance password")

	err = instance.UpdatePass(client, createResp.Instance.Id, "Root1234#New")
	th.AssertNoErr(t, err)

	t.Logf("Attempting to change taurus db instance flavor")
	resizeFlavorJob, err := instance.Resize(client, createResp.Instance.Id, "gaussdb.mysql.2xlarge.x86.8")
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, job.WaitForJobSuccess(client, 1200, *resizeFlavorJob))

	t.Logf("Attempting to change taurus db instance port")
	portChangeJob, err := instance.UpdatePort(client, createResp.Instance.Id, 3307)
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, job.WaitForJobSuccess(client, 600, *portChangeJob))

	t.Logf("Attempting to change taurus db collection period")
	monitoringJob, err := instance.UpdateSecondLevelMonitoring(client, instance.UpdateSecondLevelMonitoringOpts{
		InstanceId:    createResp.Instance.Id,
		Period:        5,
		MonitorSwitch: true,
	})
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, job.WaitForJobSuccess(client, 600, *monitoringJob))

	t.Logf("Attempting to get taurus db collection period")

	getMonitoringResp, err := instance.GetSecondLevelMonitoring(client, createResp.Instance.Id)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getMonitoringResp.Period, 5)
	th.AssertEquals(t, getMonitoringResp.MonitorSwitch, true)
}

func TestTaurusReadReplicaLifecycle(t *testing.T) {
	t.Skip("too long to run in ci")
	vpcID := os.Getenv("OS_VPC_ID")
	subnetID := os.Getenv("OS_NETWORK_ID")
	secGroupID := os.Getenv("OS_SECURITY_GROUP_ID")

	client, err := clients.NewTaurusDBV3Client()
	th.AssertNoErr(t, err)

	createResp := createTaurusInstance(t, client, vpcID, subnetID, secGroupID)

	t.Cleanup(func() {
		t.Logf("Attempting to delete taurus db")
		_, err = instance.Delete(client, createResp.Instance.Id)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to create taurus db read replica")

	createReplicaJobs, err := instance.CreateReplica(client, createResp.Instance.Id,
		instance.CreateReplicaOpts{
			Priorities: []int{1, 2},
		},
	)
	th.AssertNoErr(t, err)

	replicaJobs := strings.Split(*createReplicaJobs, ",")

	th.AssertNoErr(t, job.WaitForJobSuccess(client, 600, replicaJobs[0]))
	th.AssertNoErr(t, job.WaitForJobSuccess(client, 600, replicaJobs[1]))

	getResp, err := instance.Get(client, createResp.Instance.Id)
	th.AssertNoErr(t, err)

	nodeId := getResp.Nodes[0].Id

	t.Cleanup(func() {
		t.Logf("Attempting to delete taurus db read replica")
		deleteNodeJob, err := instance.DeleteReplica(client, createResp.Instance.Id, nodeId)
		th.AssertNoErr(t, err)
		th.AssertNoErr(t, job.WaitForJobSuccess(client, 600, *deleteNodeJob))
	})
}

func TestTaurusListDatastores(t *testing.T) {
	client, err := clients.NewTaurusDBV3Client()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to list taurus db data stores")
	listResp, err := instance.ListDatastores(client, "gaussdb-mysql")
	th.AssertNoErr(t, err)

	tools.PrintResource(t, listResp)
}

func TestTaurusListFlavors(t *testing.T) {
	client, err := clients.NewTaurusDBV3Client()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to list taurus db flavors")
	flavorOpts := instance.ListFlavorsOpts{
		AvailabilityZoneMode: "multi",
		DatabaseName:         "gaussdb-mysql",
	}
	listResp, err := instance.ListFlavors(client, flavorOpts)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, listResp[0])
}
