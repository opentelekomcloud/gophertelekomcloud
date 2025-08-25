package v3

import (
	"os"
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/taurus/v3/backup"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/taurus/v3/instance"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/taurus/v3/job"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestTaurusBackupLifecycle(t *testing.T) {
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

	th.AssertNoErr(t, waitForInstanceAvailable(client, 500, createResp.Instance.Id))

	t.Logf("Attempting to get taurus db backup policy")
	getPolicy, err := backup.GetPolicy(client, createResp.Instance.Id)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, getPolicy.KeepDays, 1)
	th.AssertEquals(t, getPolicy.StartTime, "08:00-09:00")

	t.Logf("Attempting to update taurus db backup policy")
	updatePolicy, err := backup.UpdatePolicy(client, backup.UpdatePolicyOpts{
		InstanceId: createResp.Instance.Id,
		StartTime:  "10:00-11:00",
		KeepDays:   2,
		Period:     "1,2,3,4,5",
	})
	th.AssertNoErr(t, err)

	th.AssertEquals(t, updatePolicy.Status, "COMPLETED")
	th.AssertEquals(t, updatePolicy.InstanceName, createResp.Instance.Name)

	t.Logf("Attempting to create taurus db manual backup")

	createBackupResp, err := backup.Create(client, backup.CreateOpts{
		InstanceId:  createResp.Instance.Id,
		Name:        "Test-taurus-backup",
		Description: "TF backup",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createBackupResp.Backup.Name, "Test-taurus-backup")
	th.AssertEquals(t, createBackupResp.Backup.Description, "TF backup")

	th.AssertNoErr(t, job.WaitForJobSuccess(client, 600, createBackupResp.JobId))

	t.Cleanup(func() {
		t.Logf("Attempting to delete taurus db backup")
		_, err = backup.Delete(client, createBackupResp.Backup.Id)
		th.AssertNoErr(t, err)
		th.AssertNoErr(t, WaitForBackupDelete(client, 500, createBackupResp.Backup.Id))
	})

	t.Logf("Attempting to list taurus db backup policy")

	listResp, err := backup.List(client, backup.ListOpts{
		InstanceId: createResp.Instance.Id,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(listResp) >= 1, true)
	th.AssertEquals(t, listResp[0].Name, "Test-taurus-backup")
	th.AssertEquals(t, listResp[0].Description, "TF backup")
}

func WaitForBackupDelete(client *golangsdk.ServiceClient, secs int, backupId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		backups, err := backup.List(client, backup.ListOpts{
			BackupId: backupId,
		})
		if err != nil {
			return false, err
		}

		time.Sleep(15 * time.Second)

		if len(backups) == 0 {
			return true, nil
		}

		return false, nil
	})
}

func waitForInstanceAvailable(client *golangsdk.ServiceClient, secs int, instanceID string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		dcsInstances, err := instance.Get(client, instanceID)
		if err != nil {
			return false, err
		}
		if dcsInstances.Status == "ACTIVE" {
			return true, nil
		}
		return false, nil
	})
}
