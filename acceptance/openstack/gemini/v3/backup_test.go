package v3

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/gemini/v3/backup"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGeminiBackupsLifecycle(t *testing.T) {
	instanceId := os.Getenv("OS_INSTANCE_ID")
	if instanceId == "" {
		t.Skip("OS_INSTANCE_ID is required for backup test")
	}

	client, err := clients.NewGeminiDBClient()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to get backup policy for instance: %s", instanceId)
	getResp, err := backup.GetBackupPolicy(client, backup.GetBackupPolicyOpts{
		Type:       "Instance",
		InstanceId: instanceId,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, getResp)

	originalPolicy := getResp.BackupPolicy

	t.Logf("Attempting to set backup policy")

	th.AssertNoErr(t, waitForInstanceAvailable(client, 1200, instanceId))

	setOpts := backup.SetBackupPolicyOpts{
		InstanceId: instanceId,
		BackupPolicy: backup.BackupPolicy{
			KeepDays:  7,
			StartTime: "01:00-02:00",
			Period:    "1,2,3,4,5,6",
		},
	}
	err = backup.SetBackupPolicy(client, setOpts)
	th.AssertNoErr(t, err)
	t.Logf("Backup policy set successfully")

	t.Logf("Verifying updated backup policy")
	updatedResp, err := backup.GetBackupPolicy(client, backup.GetBackupPolicyOpts{
		InstanceId: instanceId,
		Type:       "Instance",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 7, updatedResp.BackupPolicy.KeepDays)
	th.AssertEquals(t, "01:00-02:00", updatedResp.BackupPolicy.StartTime)
	tools.PrintResource(t, updatedResp)

	if originalPolicy != nil {
		t.Logf("Restoring original backup policy")
		restoreOpts := backup.SetBackupPolicyOpts{
			InstanceId: instanceId,
			BackupPolicy: backup.BackupPolicy{
				KeepDays:  originalPolicy.KeepDays,
				StartTime: originalPolicy.StartTime,
				Period:    originalPolicy.Period,
			},
		}
		err = backup.SetBackupPolicy(client, restoreOpts)
		th.AssertNoErr(t, err)
	}
}
